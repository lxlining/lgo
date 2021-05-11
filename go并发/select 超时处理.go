package main

import (
	"fmt"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 20:37
 **/
func main() {
	ch:=make(chan int)
	quit:=make(chan bool)
	go func() {
		for  {
			select {
			case num:=<-ch:
				fmt.Println("读",num)
			case <-time.After(3*time.Second):
				quit<-true
				return  //goto label
			}
		}
		//label:
	}()
	for i := 0; i < 10; i++ {
		ch<-i
		time.Sleep(1500*time.Millisecond)
	}
	<-quit
	fmt.Println("over")
}
