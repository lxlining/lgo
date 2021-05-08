package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/3 23:02
 **/
func main0011() {
	var m map[int]string
	m = make(map[int]string, 1)
	m[11] = "张三"
	m[8] = "aaa"
	m[1] = "李四"
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	//删除delete
	delete(m, 1)

	fmt.Println(m)
}

func demo(m map[int]string) {
	m[1010] = "abc"
}
func main() {
	m := make(map[int]string)
	m[101] = "aaa"
	m[102] = "bbb"
	m[103] = "sss"
	demo(m)
	fmt.Println(m)
}
