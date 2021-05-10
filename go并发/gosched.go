package main

import (
	"fmt"
	"runtime"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 22:07
 **/

func main() {
	go func() {
		//for i := 0; i < 10; i++ {
			fmt.Println("-gosched----")
			//time.Sleep(time.Second)
		//}
	}()
	for  {
		runtime.Gosched()
		fmt.Println("--------------main--------")
		//time.Sleep(100*time.Millisecond)
	}
}
