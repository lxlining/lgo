package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 20:53
 **/
//全局定义
var mutex sync.Mutex

//打印机
func printer1(s string) {
	mutex.Lock()	//访问共享数据前加锁  一把锁
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
	mutex.Unlock()
}
func person11() {
	printer1("hello")

}
func person21() {
	printer1("world")
}
func main() {
	go person11()
	go person21()
	for {
		;
	}
}

