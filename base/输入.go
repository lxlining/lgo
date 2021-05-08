package main

import (
	"fmt"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/4/9 12:11
 **/
func main() {
	/*var a int
	fmt.Scan(&a)
	fmt.Printf("%p\n",&a)//%p占位符  输出内存地址
	fmt.Println(a)*/
	/*a:=time.Now()
	fmt.Println(fb(5))
	b:=time.Now().Sub(a)
	fmt.Println(b)*/
	//yy(2016)
	//if b>0 {run()}else{}
	ranged()

}
func ranged() {
	n := 690
	if n > 650 {
		if n > 680 {
			fmt.Println("1111111111111111111")
		} else if n > 690 {
			fmt.Println("222222222222222")
		} else {
			fmt.Println("333333333333333")
		}
	} else {
		fmt.Println("55555555555555")
	}
}
func fb(i int) int {
	if i < 2 {
		return 1
	}
	{
		return fb(i-1) + fb(i-2)
	}
}
func run() {
	a := time.Now()
	//a:=time.Now()
	//const i int=1000000
	//var s string
	for i := 0; i <= 1000000000000; i++ {
		i = i + 1000
	}
	//fmt.Println(s)
	b := time.Now().Sub(a)
	//c := b - a
	fmt.Println(b)
}
func c() {
	//bool
	//float32 float64   默认保留7位末尾位+1  15位小数
	//ASCII 0-48 A-65 a-97
	//string  var a string  \0字符串的结束标志
	//字符串相加连接 c:=a+b
	//一个汉字3个字符
	//byte ‘a’

}
func yy(i int) {
	x := (i%400 == 0) || (i%4 == 0 && i%100 != 0)
	if x {
		fmt.Println("yes")
	}
	fmt.Println("no")
}
