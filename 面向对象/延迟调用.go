package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 11:27
 **/
//defer放栈区 先进后出
//函数中有返回值，不能使用defer
//recover 只在 defer 中使用 必须在错误出现之前
func tt(i int)  {
	var arr [10]int
	defer func() {
		//fmt.Println(recover())//拦截错误信息
		err:=recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i]=100
	fmt.Println(arr)
}
func main() {
	defer fmt.Println("hdjfujsdsgfchmsdhku")
	fmt.Println("fhxgfdgzdgmsdhku")
	defer fmt.Println("zzztzsdsgfchmsdhku")
	fmt.Println("retetsdsgfchmsdhku")
	tt(10)
}
