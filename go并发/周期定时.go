package main

import (
	"fmt"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 19:10
 **/
func main() {
	quit:=make(chan bool)
	fmt.Println(time.Now())
	my:=time.NewTicker(time.Second)
	//fmt.Println(time.Now())
	i:=0
	go func() {
		for{
			now:=<-my.C
			i++
			fmt.Println(now)
			if i==10{
				quit<-true
				break  //return runtime.GoExit  同
			}
		}
	}()

	<-quit
}
