package main

import (
	"fmt"
	"net"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 17:05
 **/
func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("dial err",err)
		return
	}
	defer conn.Close()

	httpRequest:="Get /tt HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	conn.Write([]byte(httpRequest))

	buf:=make([]byte,4096)
	n,_:=conn.Read(buf)
	if n==0 {
		return
	}
	fmt.Println(string(buf))
}
