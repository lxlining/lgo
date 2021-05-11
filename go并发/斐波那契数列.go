package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 20:25
 **/

func fibnacci(ch chan int,quit chan bool)  {
	for  {
		select {
		case num:=<-ch:
			fmt.Println(num)
		case <-quit:
			return
		}
	}
}
func main() {
	ch:=make(chan int)
	quit:=make(chan bool)
	go fibnacci(ch,quit)
	x,y:=1,1
	for i := 0; i < 20; i++ {
		ch<-x
		x,y=y,x+y
	}
	quit<-true
}
