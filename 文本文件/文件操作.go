package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 12:46
 **/

func main() {
	//文件拷贝
	fp,err:=os.Open("文本文件/a.txt")
	fp1,err1:=os.Create("文本文件/d.txt")
	if err !=nil&&err1 !=nil{
		fmt.Println("拷贝失败")
		return
	}
	defer fp.Close()
	defer fp1.Close()

	b:=make([]byte, 1024)
	for{
		n,err:=fp.Read(b)
		if err == io.EOF{
			break
		}
		fp1.Write(b[:n])
	}
	fmt.Println("完成")
}
