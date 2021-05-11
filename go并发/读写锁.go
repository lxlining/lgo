package main

import (
	"fmt"
	rand2 "math/rand"
	"sync"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 21:04
 **/
//读  ＋R Rlock()
//不要将互斥锁、读写锁与channel 混用  造成隐形死锁

var rwMutex sync.RWMutex
var value int
func read( idx int)  {
	for true {
		rwMutex.RLock()
		num:=value
		fmt.Printf("读%d==%d\n",idx,num)
		rwMutex.RUnlock()
	}

}
func write(idx int)  {
	for true {
		num:=rand2.Intn(1000)
		rwMutex.Lock()
		value= num
		fmt.Printf("写%d==%d\n",idx,num)
		time.Sleep(300*time.Millisecond)
		rwMutex.Unlock()
	}
}
func main() {
	rand2.Seed(time.Now().UnixNano())
	//quit:=make(chan bool)
	//ch:=make(chan int)
	for i := 0; i < 5; i++ {
		go read(i+1)
	}
	for i := 0; i < 5; i++ {
		go write(i+1)
	}
	for true {
		;
	}
}
