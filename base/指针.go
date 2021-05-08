package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/7 20:05
 **/

func main110() {
	var a int = 10
	var p *int
	p = &a
	fmt.Println(p)
	var q *int
	*q=123
	fmt.Println(*q)
}

func swap(a *int, b *int)  {
	temp:=*a
	*a=*b
	*b=temp
}
func main100()  {
	//new操作
	var p *int
	fmt.Println(p)
	p=new(int)//堆区创建空间 值为变量类型的默认值
	fmt.Println(p)//打印p的值--地址
	fmt.Println(*p)//打印p对应空间的值
	a:=10
	b:=20
	swap(&a,&b)
	fmt.Println(a,b)
}

//数组指针
func main()  {
	var arr [5]int=[5]int{1,2,3,4,5}
	//p:=&arr
	var p *[5]int
	p=&arr
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println((*p)[0])
	fmt.Println(p[0])


}
