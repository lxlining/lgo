package main

import (
	"fmt"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 21:59
 **/
//主go程结束 子go程退出
func main() {
	go func (){
		for i := 0; i <5 ; i++ {
			fmt.Println("------------goroutine---------")
			time.Sleep(time.Second)

		}
	}()
	for i := 0; i <5 ; i++ {
		fmt.Println("------------hello world---------")
		time.Sleep(100*time.Millisecond)
		if i==1{
			break
		}
	}
}
