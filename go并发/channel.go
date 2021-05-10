package main

import (
	"fmt"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 22:38
 **/
//channel 是一种数据类型 对应管道 FIFO

//全局定义
var channel = make(chan int)

//打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}
func person1() {
	printer("hello")
	channel <- 1
}
func person2() {
	<- channel
	printer("world")
}
func main() {
	go person1()
	go person2()
	for {
		;
	}
}
