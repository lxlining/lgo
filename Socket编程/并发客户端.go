package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 17:33
 **/
func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer conn.Close()
	//获取键盘输入
	go func() {
		for  {
			str:=make([]byte ,4096)
			n,err:=os.Stdin.Read(str)
			if err != nil {
				fmt.Println("err:",err)
				continue
			}
			//写服务器
			conn.Write(str[:n])
		}
	}()
	//服务器回写
	buf:=make([]byte,4096)
	for  {
		n,err:=conn.Read(buf)
		if err != nil {
			fmt.Println("err:",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
