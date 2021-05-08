package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/4 23:03
 **/
type students struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main1110() {
	m := make(map[int]students)
	m[1] = students{101, "lll", "男", 20, "北京"}
	m[2] = students{102, "ppp", "男", 24, "cq"}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func main() {
	m := make(map[int][]students)
	m[101] = append(m[101], students{101, "lll", "男", 20, "北京"}, students{102, "ppp", "男", 24, "cq"})
	m[102] = append(m[102], students{103, "kkk", "男", 24, "cq"}, students{104, "hhh", "男", 24, "cq"})
	for k, v := range m {
		for i, data := range v {
			fmt.Println("key:", k, "index:", i, "data:", data)
		}
	}
}
