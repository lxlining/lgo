package main

import (
	"fmt"
	"runtime"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 20:16
 **/
func main() {
	ch:=make(chan int)
	quit:=make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch<-i
		}
		close(ch)
		quit<-true
		runtime.Goexit()
	}()
	for{
		select {
		case num:=<-ch:
			fmt.Println("读",num)
		case <-quit:
			return

		}
		fmt.Println("=======")
	}
}
