package main

import "fmt"

/**
 * @Description 函数
 * @Author lixl
 * @Date 2021/4/11 15:30
 **/
func main() {
	aa(1)
	aa(1,2,3,4,5,6)
}
//不定参数  参数名...类型
func aa(a ...int)  {
	for i:=0;i< len(a);i++{
		fmt.Println(a[i])
	}
	/*for i,data:=range a{

	}*/
}