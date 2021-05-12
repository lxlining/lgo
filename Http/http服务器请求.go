package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 16:32
 **/
func errFunc(err error,info string)  {
	if err != nil {
		fmt.Println(info,err)
		//return  //返回当前函数调用
		//runtime.Goexit()  //结束当前go程
		os.Exit(1)  //无论什么时候 将当前进程结束
	}
}
func main() {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	errFunc(err,"net.listen err")
	defer listener.Close()

	conn,err:=listener.Accept()
	errFunc(err,"Accpet err")
	defer conn.Close()

	buf:=make([]byte ,4096)
	n,err:=conn.Read(buf)
	if n==0{
		return
	}
	errFunc(err,"conn.read err")

	fmt.Println(string(buf[:n]))

}
