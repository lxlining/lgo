package main

import (
	"fmt"
	"regexp"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 20:31
 **/
func main() {
	//(?s:(*?))//单行模式
	str:="<div>abc</div> mcb ahc <div>cnbacc</div> ncbcbcda ac 3.14 2 5.1 0.1 0. .0 1.5412"
	ret:=regexp.MustCompile(`<div>(?s:(.*?))</div>`)// 反引号 `d.d`
	als:=ret.FindAllStringSubmatch(str,-1)
	//fmt.Println(strings.Trim(string(als)),"<div>")
	for _,one:=range als{
		//fmt.Println("one[0]=",one[0]) //one[0]= <div>abc</div>
		fmt.Println("one[1]=",one[1]) //one[1]= abc
	}
}
