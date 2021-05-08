package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/4 22:45
 **/
type Student struct {
	id   int
	name string
}

func main() {
	var arr [5]Student
	arr[0].id = 10
	arr[0].name = "aaa"
	arr[1].id = 11
	arr[1].name = "111"
	arr[2].id = 2
	arr[2].name = "222"
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].id > arr[j+1].id {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
