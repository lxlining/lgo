package main

import (
	"fmt"
	"strings"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 13:01
 **/

func main() {
	str:="bhdg,fjshfsffd"
	//在一个字符串中查找另一个字符串
	value:=strings.Contains(str,"dg")
	if value {
		fmt.Println("find")
	}else{
		fmt.Println("not found")
	}
	fmt.Println(value)
	//join 字符串拼接
	s:=[]string{"gdh","jhdgjs","hsdg","idfug"}
	str=strings.Join(s,"-")
	fmt.Println(str)
	//index 字符串查找，返回下标
	//repeat 重复 次数
	//replace 字符串替换
	x:=strings.Replace(str,"i","|",-1)
	fmt.Println(x)
	//split  字符串分割
	//trim  字符串两端去掉 内容
	//Fileds 去掉字符串的空格 返回切片

}
//strconv.format 其他转字符串
//strconv.parse 字符串转其他

