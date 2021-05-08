package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/7 22:03
 **/
func main() {
	a:=10
	p:=&a
	pp:=&p  //二级指针存储一级指针的地址   ppp:=&pp 三级指针存储二级指针的地址
	fmt.Println(a,*p,**pp)
}
