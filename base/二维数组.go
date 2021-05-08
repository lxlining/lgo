package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/4/20 21:38
 **/
func main() {
	var arr [3][4]int
	arr[1][1] = 1
	arr[2][2] = 2
	//len(arr)二维数组的行数   len(arr[0]) 二维数组的列数
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
