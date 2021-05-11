package main

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/10 20:43
 **/
func main() {
	//死锁1
	/*ch:=make(chan int)
	ch<-123
	num:=<-ch
	fmt.Println(num)*/
	//死锁2
	/*ch:=make(chan int)
	num:=<-ch	//此时，没人读
	fmt.Println(num)
	go func() {
		ch<-123
	}()*/
	//死锁3
	ch:=make(chan int)
	ch1:=make(chan int)

	go func() {
		for  {
			select {
			case num:=<-ch:
				ch1<-num

			}
		}
	}()
	for  {
		select {
		case num:=<-ch1:
			ch<-num

		}
	}
}
