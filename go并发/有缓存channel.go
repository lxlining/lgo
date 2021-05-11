package main

import (
	"fmt"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 16:59
 **/

func main() {
	ch:=make(chan int,3)
	fmt.Println("lan=",len(ch),"cap",cap(ch))
	go func() {
		for i := 0; i < 10; i++ {
			ch<-i
			fmt.Println("子go程i=",i)

		}

	}()
	//time.Sleep(2*time.Second)
	for i := 0; i < 8; i++ {
		num:=<- ch
		fmt.Println("主go程读--",num)
	}
}
