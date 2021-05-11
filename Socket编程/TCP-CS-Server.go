package main

import (
	"fmt"
	"net"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 16:46
 **/

func main() {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	defer listener.Close()
	fmt.Println("等待客服端连接....")
	conn,err:=listener.Accept()
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端建立连接")
	//读取客户端数据
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	//处理
	conn.Write(buf[:n])
	fmt.Println("服务器读到：",string(buf[:n]))

}
