package main

import (
	"fmt"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 22:34
 **/
func main() {
	//获取命令行参数
	list:=os.Args
	if len(list) !=2{
		fmt.Println("err")
		return
	}
	fmt.Println(list)
	path:=list[1]
	fileinfo,err:=os.Stat(path)
	if err != nil {
		fmt.Println("err",err)
		return
	}
	fmt.Println("文件名：",fileinfo.Name())
	fmt.Println("文件大小：",fileinfo.Size())

}
