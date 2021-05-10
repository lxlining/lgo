package main

import (
	"fmt"
	"runtime"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 22:28
 **/
func main() {
	fmt.Println(runtime.GOROOT())
	n:=runtime.GOMAXPROCS(4)
	fmt.Println("n===",n)
for{
	go fmt.Print(111111)
	fmt.Print(22222)
}
	n=runtime.GOMAXPROCS(4)
	fmt.Println("n===",n)
}
