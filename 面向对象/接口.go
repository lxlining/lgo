package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 23:00
 **/
type per struct {
	name string
	age int
}
type Stu struct {
	per
	score int
}
type teacher struct {
	per
	sub string
}

func (s *Stu)SayHello()  {
	fmt.Printf("大家好，我是%s，我%d岁，我的成绩是%d分\n",s.name,s.age,s.score)
}
func (t *teacher)SayHello()  {
	fmt.Printf("大家好，我是%s，我%d岁，我的学科是%s分\n",t.name,t.age,t.sub)
}
//接口定义
type Hunman interface {
	SayHello()
}

//多态实现
func SayHello(h Hunman)  {
	h.SayHello()
}
func main() {
	var stu Stu=Stu{per{"aa",22},100}
	var tea teacher=teacher{per{"bb",33},"shuxue"}
	stu.SayHello()
	tea.SayHello()
	var h Hunman
	//h=&stu
	h=&Stu{per{"xx",22},100}
	//h.SayHello()
	SayHello(h)
}
//接口继承  同方法
