package main

import "fmt"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/8 23:55
 **/
type aa interface {

}
func main() {
	//空接口可以接受任意类型数据
	var i interface{}
	i=10
	fmt.Println(i)
	i="shshs"
	fmt.Println(i)

	//空接口切片
	var s []interface{}
	//ss:=make([]interface{},1)
	s=append(s,1,2,3,"gdhsh",'0',"'1'")
	fmt.Println(s)
}
//断言
//data,ok:=v.(int)
//data,ok:=v.(string)
//data,ok:=v.([2]int)
//data,ok:=v.([]int)