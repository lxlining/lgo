package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 21:15
 **/
type ta struct {
	name string
	id int
}
type tb struct {
	ta
	sex string
	age int
}
type tc struct {
	tb
	score int
}
type td struct {
	tb
}
func main() {
	var s tc
	s.name="aa"
	s.id=101
	s.age=23
	s.score=100
	s.sex="nn"
	fmt.Println(s)
}
