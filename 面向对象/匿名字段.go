package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 20:49
 **/

type person struct {
	name string
	age int
	sex string
}
type stu struct {
	person//匿名字段实现继承
	id int
	socre int
}
func main() {
	var s stu
	s.name="aaa"
	//s.person.name="aaa"  同上
	s.id=1004
	s.sex="男"
	s.age=24
	s.socre=100
	fmt.Println(s)
	var ss stu=stu{person{"bbb",100,"安"},101,99}
	fmt.Println(ss)
}
