package main

import (
	"fmt"
	"net"
	"strings"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 17:11
 **/
func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	//获取客户端ip
	addr:=conn.RemoteAddr()
	fmt.Println("连接ip:",addr)
	//循环读取
	buf:=make([]byte,4096)
	for  {
		n,err:=conn.Read(buf)
		// /n or /r/n
		if "exit\n"==string(buf[:n])||"exit\r\n"==string(buf[:n]){
			fmt.Println("客服端离开。。。。。。。")
			return
		}
		if n==0 {
			fmt.Println("客服端离开。。。。。。。")
			return
		}
		if err != nil {
			fmt.Println("err:",err)
			return
		}
		//处理
		//
		fmt.Println("服务器读到：",string(buf[:n]))
		//回写 客户端
		conn.Write([]byte (strings.ToUpper(string(buf[:n]))))
	}
}
func main() {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	defer listener.Close()
	//监听客户端
	fmt.Println("等待客服端连接....")
	for{
		conn,err:=listener.Accept()
		if err != nil {
			fmt.Println("err:",err)
			return
		}
		defer conn.Close()
		//数据通信
		go HandlerConnect(conn)
	}
}
