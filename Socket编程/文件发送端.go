package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 22:42
 **/

func sendFile(conn net.Conn,filepath string)  {
	fp,err:=os.Open(filepath)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer fp.Close()
	buf:=make([]byte,4096)
	for{
		n,err:=fp.Read(buf)
		if err != nil {
			if err ==io.EOF{
				fmt.Println("文件传输完毕")
			}else{
				fmt.Println("err:",err)
			}
			return
		}
		_,err=conn.Write(buf[:n])
		if err != nil {
			fmt.Println("err:",err)
			return
		}

	}
}
func main() {
	list:=os.Args
	if len(list) !=2{
		fmt.Println("文件格式 go run xxx.go 文件名")
		return
	}
	filepath:=list[1]
	fileinfo,err:=os.Stat(filepath)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	filename:=fileinfo.Name()

	conn,err:=net.Dial("tcp","127.0.0.1:9001")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(filename))
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	if "ok"==string(buf[:n]){
		sendFile(conn,filepath)
		return
	}

}
