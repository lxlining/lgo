package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 17:12
 **/
func main() {
	ch :=make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch<-i
		}
		close(ch)
	}()
	//time.Sleep(2*time.Second)
	for {
		if num,ok:=<-ch;ok==true{
			fmt.Println("dudaoshuju---",num)
		}else{
			break
		}
	}
}
