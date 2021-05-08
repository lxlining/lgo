package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/4/20 21:45
 **/
//动态数组
func main001() {
	//var arr []int
	s := make([]int, 5)
	s[0] = 123
	s[1] = 234
	s = append(s, 111, 222, 333)
	fmt.Println(s)
	//遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Println(cap(s))
}

//切片截取
func main002() {
	s := []int{1, 2, 3, 4, 5}
	slice := s[2:5] //包含起始位置  不包含结束位置
	fmt.Println(slice)
}

func main003() {
	//追加 append
	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	//拷贝 copy
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := make([]int, 6)
	copy(s2, s1)
	fmt.Println(s2)
	s3 := make([]int, 3)
	copy(s3, s1[1:])
	fmt.Println(s3)
}

func test1(s []int) {
	fmt.Printf("%p\n", s)
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	test1(s)
	fmt.Printf("%p\n", s)
}
