package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 16:52
 **/
func main() {
	ch:=make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程i=",i)
			ch<-i
		}

	}()
	for i := 0; i < 5; i++ {
		num:=<- ch
		fmt.Println("主go程读--",num)
	}
}
