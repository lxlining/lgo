package main

import (
	"fmt"
)

func main()  {
	//main1()
	tt()
}

func tt()  {
	a,b,c := 10,20,101
	a,b = b,a//次方法类型必须一致
	fmt.Println(a,b,c)
}

func main1()  {
	//fmt.Println(time.Now())
	fmt.Println("hello go!")
	fmt.Println(hee())
	//声明--定义
	var a int
	var aa int = 10
	fmt.Println(a)
	fmt.Println(aa)

	//fmt.Println(time.Now())

	a1 := 12
	fmt.Println(a1)
}

func hee() int {
	return 111
}
