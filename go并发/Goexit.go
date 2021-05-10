package main

import (
	"fmt"
	"runtime"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 22:15
 **/
func test()  {
	fmt.Println("ccccccccccc")
	runtime.Goexit()
	fmt.Println("ddddddddddddd")
}
func main() {

	go func() {
		defer fmt.Println("aaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbbbb")
	}()
	for{
		;
	}
}
