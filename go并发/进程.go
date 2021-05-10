package main

import (
	"fmt"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 21:47
 **/
func sing(){
	for i := 0; i <300 ; i++ {
		fmt.Println("---唱歌中---")
		//time.Sleep(100*time.Millisecond)
	}
}
func dance()  {
	for i := 0; i <300 ; i++ {
		fmt.Println("---跳舞中---")
		//time.Sleep(100*time.Millisecond)
	}
}
func main() {
	/*sing()
	dance()*/
	go sing()
	go dance()
	for{
		;
	}
}
