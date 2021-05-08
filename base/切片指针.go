package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/7 21:08
 **/


func main11111() {
	var slice[]int=[]int{1,2,3,4,5}
	var p *[]int
	p=&slice
	*p=append(*p,1,2,3)
	fmt.Println(slice)

	var q *[]int
	q=new([]int)
	*q=append(*q,1,2,3,4,5,6)
	fmt.Println(*q)
	for i,v:=range *q{
		fmt.Println(i,v)
	}
}

//指针数组
func main0()  {
	var arr [3]*int
	a:=10
	b:=20
	c:=30
	arr[0]=&a
	arr[1]=&b
	arr[2]=&c
	for i := 0; i <len(arr) ; i++ {
		fmt.Println(*arr[i])
	}
}
//切片指针
func main()  {
	var arr []*int
	a:=10
	b:=20
	c:=30
	arr=append(arr,&a,&b,&c)
	fmt.Println(arr)
	for i,v:=range arr{
		fmt.Println(i,*v)
	}
}