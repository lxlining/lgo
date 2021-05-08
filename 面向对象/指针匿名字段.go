package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 21:04
 **/
type person2 struct {
	name string
	age int
	sex string
}
type stu2 struct {
	*person2//匿名字段实现继承
	name string
	id int
	score int
}
func main() {
	var ss stu2=stu2{&person2{"xx",101,"女"},"aaaa",10001,100}
	var s stu2
	s.person2=new(person2)
	var per=person2{"aaa",101,"男"}
	s.person2=&per
	s.name="琪琪"
	s.sex="女"
	s.age=12
	s.person2.name="qiqi"
	s.id=111
	s.score=100
	fmt.Println(s.name)
	fmt.Println(s.id)
	fmt.Println(s.score)
	fmt.Println(per)
	fmt.Println(ss)
}
