package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 20:59
 **/

type person1 struct {
	name string
	age int
	sex string
}
type stu1 struct {
	person1//匿名字段实现继承
	id int
	name string
	score int
}
func main() {
	var s stu1
	//采用就近原则 使用子类信息
	s.name="xxx"
	s.person1.name="xxxxxxx"
	s.id=10001
	s.age=201
	s.sex="s"
	fmt.Println(s)
}
