package main

import (
	"fmt"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 18:56
 **/

//定时器 time.NewTimer
//定时器  停止、重置  stop()  reset()
func main() {
	fmt.Println(time.Now())
	mytime:=time.NewTimer(2*time.Second)
	now:=<-mytime.C
	fmt.Println(now)
	now=<-time.After(time.Second*3)
	fmt.Println(now)

}
