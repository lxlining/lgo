package main

import (
	"errors"
	"fmt"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/9 11:15
 **/


func test(a int,b int) (value int,err error) {
	if b==0{
		err=errors.New("0不能作为除数")
		return
	}else{
		value=a/b
		return
	}
}
func main() {
	value,err:=test(10,0)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(value)
	}
}
//panic异常 程序终止
