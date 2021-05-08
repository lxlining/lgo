package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/3 22:41
 **/
func main() {
	rand2.Seed(time.Now().UnixNano())
	//random:=rand2.Intn(900)+100
	random := make([]int, 3)
	random[0] = rand2.Intn(9) + 1
	random[1] = rand2.Intn(10)
	random[2] = rand2.Intn(10)
	var num int
	var flag int = 0
	for {
		for {
			fmt.Println("请输入一个三位数：")
			fmt.Scan(&num)
			fmt.Println(random)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入错误，重新输入")
		}
		unum := make([]int, 3)
		unum[0] = num / 100
		unum[1] = num / 10 % 10
		unum[2] = num % 10

		for i := 0; i < 3; i++ {
			if unum[i] > random[i] {
				fmt.Printf("您输入的第%d位数太大了\n", i+1)
			} else if unum[i] < random[i] {
				fmt.Printf("您输入的第%d位数太小了\n", i+1)
			} else {
				fmt.Printf("您输入的第%d位数相同\n", i+1)
				flag++
			}
		}
		if flag == 3 {
			fmt.Println("成功")
			break
		} else {
			flag = 0
		}
	}
}
