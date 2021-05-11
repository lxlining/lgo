package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 16:44
 **/

func main() {
	ch:=make(chan string)
	fmt.Println("lan=",len(ch),"cap",cap(ch))
	//死锁
	//ch<-"子go打印完毕"
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i=",i)
		}
		//通知主go打印完毕
		ch<-"子go打印完毕"
	}()
	str:=<-ch
	fmt.Println(str)
}
