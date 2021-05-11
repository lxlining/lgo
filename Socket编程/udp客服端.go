package main

import (
	"fmt"
	"net"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 17:00
 **/
func main() {
	dail,err:=net.Dial("udp","127.0.0.1:9000")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer dail.Close()
	//发送数据
	dail.Write([]byte("Are you ok?"))
	//接收数据
	buf:=make([]byte ,4096)
	n,err:=dail.Read(buf)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	dail.Write(buf[:n])
	fmt.Println("服务器回传：",string(buf[:n]))


}
