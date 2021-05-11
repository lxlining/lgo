package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 22:55
 **/

func recvFile(conn net.Conn,file string)  {
	f,err:=os.Create(file)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer f.Close()
	buf:=make([]byte,4096)
	for{
		n,_:=conn.Read(buf)
		if n==0 {
			fmt.Println("接收文件完成")
			return
		}
		f.Write(buf[:n])
	}
}
func main() {
	listener,err:=net.Listen("tcp","127.0.0.1:9001")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer listener.Close()
	conn,err:=listener.Accept()
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer conn.Close()
	buf:=make([]byte,4096)

	n,err:=conn.Read(buf)
	filename:=string(buf[:n])
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	conn.Write([]byte("ok"))
	recvFile(conn,filename)
}
