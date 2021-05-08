package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 21:27
 **/
func add(a int , b int) int  {
	return a+b
}

type Int int

func (a Int)sub(b Int)Int  {
	return a-b
}
func main() {
	i:=add(10,20)
	fmt.Println(i)
	var a Int=10
	I:=a.sub(10)
	fmt.Println(I)
}
//结构体 s *Stu  指针  地址传递 改变变量
//可以直接使用结构体指针对应的方法
//多个 p []person
//方法重写
//调用子类 子类.方法名
//调用父类 子类.父类.方法名

