package main

import "fmt"

/**
 * @Description 符合类型
 * @Author lixl
 * @Date 2021/4/12 18:04
 **/
func f() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//值传递 不会改变外部变量
func test(arr [10]int) {

}
func main() {
	f()
	//自动类型推到
	arr1 := [...]int{}
	fmt.Println(arr1)
}
