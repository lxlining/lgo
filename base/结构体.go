package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/4 21:53
 **/

type student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main() {
	var s student
	s.id = 1001
	s.name = "暂时"
	s.sex = "男"
	s.age = 20
	s.addr = "北京"
	fmt.Println(s)

	var ss student = student{101, "浏览", "女", 20, "上海"}
	fmt.Println(ss.name)
}
