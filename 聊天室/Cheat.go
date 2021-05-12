package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 23:46
 **/
//创建用户结构体
type Client struct {
	C chan string
	Name string
	Addr string
}

//创建全局map 存储用户信息
var onlineMap map[string]Client

//创建全局channel 存储用户消息
var message =make(chan string)

func Manger()  {
	//初始化onlineMap
	onlineMap=make(map[string]Client)
	//监听channel消息
	for{
		msg:=<-message
		//发送给用户
		for _,clnt:=range onlineMap{
			clnt.C<-msg
		}
	}
}
func WriteToClient(cli Client,conn net.Conn)  {
	//监听用户自带channel是否有消息
	for msg:=range cli.C{
		conn.Write([]byte(msg+"\n"))
	}
}
func MakeMsg(cli Client,msg string)(buf string)  {
	buf="["+cli.Addr+"]"+cli.Name+"--"+msg
	return
}
func HandlerConn(conn net.Conn)  {
	defer conn.Close()
	//创建用户
	netAddr:=conn.RemoteAddr().String()
	cli:=Client{make(chan string),netAddr,netAddr}
	//将新连接用户添加到map中
	onlineMap[netAddr]=cli
	//创建channel 判断用户是否活跃
	hasData:=make(chan bool)

	//创建专门传用户消息的go程
	go WriteToClient(cli,conn)
	//发送用户上线消息 -- 全局channel
	//message<-"["+netAddr+"]"+cli.Name+"--login"
	message<-MakeMsg(cli,"login")
	//创建channel 判断用户退出
	isQuit:=make(chan bool)
	//创建go程 处理用户消息
	go func() {
		buf:=make([]byte,4096)
		for{
			n,err:=conn.Read(buf)
			if n==0{
				isQuit<-true
				fmt.Println("检测到用户退出",cli.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:",err)
				return
			}
			//将用户消息保存到msg
			msg:=string(buf[:n-1])
			//提取用户在线列表
			if msg=="who"&&len(msg)==3{
				conn.Write([]byte("online user list:\n"))
				//遍历map，获取用户
				for _,user:=range onlineMap{
					userinfo:=user.Addr+":"+user.Name+"\n"
					conn.Write([]byte(userinfo))
				}
			}else if len(msg)>=8&&msg[:6]=="rename"{
				newname:=strings.Split(msg,"|")[1]
				cli.Name=newname  //修改结构体成员
				onlineMap[netAddr]=cli  //更新用户列表
				conn.Write([]byte("rename success!\n"))
			}else {
				//将读到的信息，写入channel
				message<-MakeMsg(cli,msg)
			}
			hasData<-true
		}
	}()
	//保证，不退出
	for{
		select {
			case <-isQuit:
				close(cli.C)//监听退出
				delete(onlineMap,cli.Addr)  //将用户从在线列表移除
				message<-MakeMsg(cli,"loginout") //写出用户退出信息到channel
				return
			case <-hasData:

			case <-time.After(time.Second*20):
			delete(onlineMap,cli.Addr)  //将用户从在线列表移除
			message<-MakeMsg(cli,"timeout----loginout") //写出用户退出信息到channel
			return

		}
	}
}
func main() {
	//创建监听Socket
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen err:",err)
		return
	}
	defer listener.Close()
	//创建管理者go程
	go Manger()
	//循环监听客服端请求
	for{
		conn,err:=listener.Accept()
		if err != nil {
			fmt.Println("Connect err:",err)
			return
		}
		//启动go程
		go HandlerConn(conn)

	}
}
