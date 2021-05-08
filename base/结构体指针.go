package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/7 21:45
 **/
type stu struct {
	id int
	name string
	sex string
}
func main() {
	var s stu=stu{1001,"aaa","男"}
	var p *stu
	p=&s
	fmt.Println(p.name)
	//结构体切片
	var ss []stu=make([]stu,1)
	pp:=&ss
	*pp=append(*pp,stu{1001,"aaa","男"})
	fmt.Println(ss)
}
