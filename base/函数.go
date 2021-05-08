package main

import "fmt"

/**
 * @Description 函数
 * @Author lixl
 * @Date 2021/4/11 15:30
 **/
func main11() {
	/*aa(1)
	aa(1,2,3,4,5,6)
	t2(1,2,3,4,5,6,7,8,9)
	t1(1)*/

	a, b, c := tt1(5, 10)
	fmt.Println(a, b, c)
}

//不定参数  参数名...类型
func aa(a ...int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i, data := range a {
		fmt.Println(i, data)
	}
}

func t1(b ...int) {
	for i := 0; i < len(b); i++ {
		fmt.Println(i, b[i])
	}
}
func t2(a ...int) {
	t1(a[0:]...) //传递不定参
}

//多返回值
func tt1(a int, b int) (i int, j int, k int) {
	i = a
	j = b
	k = a + b
	return
}

//type 定义函数类型   为已存在的类型起别名

//程序中程序出现相同变量名，函数自己定义了局部变量，优先， 没有找外部--全局变量
func main111() {
	var a int = 0
	fmt.Println(a)
	cc()
}
func cc() {
	var a int = 1
	a = 5
	fmt.Println(a)
}

//匿名函数
func main1111() {
	a := 1
	b := 2
	func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
	f := func(a int, b int) {
		fmt.Println(a + b)
	}
	f(a, b)
}

//闭包
func f1() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
func main22() {
	for i := 0; i < 5; i++ {
		//f:=f1()
		fmt.Println(f1())
	}
}

//递归
func ff(a int) {
	if a == 0 {
		return
	}
	a--
	ff(a)
	fmt.Println(a)
}

var sum int = 1

func fff(n int) {
	if n == 1 {
		return
	}
	fff(n - 1)
	sum *= n
}
func main() {
	//ff(5)
	fff(5)
	fmt.Println(sum)
}
