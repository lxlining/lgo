package main

import (
	"fmt"
)

//枚举
const (
	a=iota+1 //后续不会随着 +1而更改枚举数 1
	b=iota  //1
	e=iota  //2
	d=iota
)

func main() {
	/*aaa(1)
	aaa(2)
	aaa(0)*/
	//ss()
	//fmt.Println(a,b,c)
	//rand()
	//cf()
	qian()
}

func qian()  {
	for i:=1;i<20;i++{//5
		for j:=1;j<33;j++{//3
			/*for k:=0;k<=100;k+=3{//1/3
				if (5*i+j*3+k/3)==100 && (i+j+k)==100 && k%3==0{
					fmt.Printf("%d %d %d \n",i,j,k)
				}
			}*/
			if(5*i+3*j+(100-i-j)/3==100) && (100-i-j)%3==0{
				fmt.Printf("%d %d %d \n",i,j,100-i-j)
			}
		}
	}
}

/**
九九乘法表
 */
func cf()  {
	for i:=1;i<10;i++{
		for j:=1;j<i;j++{
			fmt.Printf("%dx%d=%d   ",j,i,i*j)
		}
		fmt.Println()
	}
}
func rand()  {
	var i int
	for i=0; i<10; i++{
		fmt.Println(i+1)
	}

}

func ss()  {
	var n int
	fmt.Println("input n")
	fmt.Scan(&n)
	switch n/2 {
		case 0:
			fmt.Println(111)
			fallthrough //让switch自动跳转到下一个条件
	case 1:
		fmt.Println(222)
	case 2:
		fmt.Println(333)
	default:
		fmt.Println(0001)
	}
}
/**
 * @Description if执行效率低， switch比较高
 * @Param
 * @return
 **/
func aaa(n int)  {
	switch n {
	case 1:
		fmt.Println(111)
	case 2:
		fmt.Println(222)
	case 3:
		fmt.Println(333)
	default:
		fmt.Println(444)



	}
}