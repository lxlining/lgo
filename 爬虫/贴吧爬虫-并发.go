package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 20:03
 **/

//
func HttpGet1(url string) (result string,err error) {
	res,err1:=http.Get(url)
	if err1 != nil {
		err=err1
		return
	}
	defer res.Body.Close()

	buf:=make([]byte,4096)
	for  {
		n,err2:=res.Body.Read(buf)
		if n==0 {
			fmt.Println("读取完毕")
			break
		}
		if err2 !=nil && err2!= io.EOF{
			err=err2
			return
		}
		result+=string(buf[:n])
	}
	return
}
//爬取单个页面
func SpiderOne(i int,page chan int)  {
	url:="https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
	result,err:=HttpGet1(url)
	if err != nil {
		fmt.Println("get err",err)
		return
	}
	//保存为一个文件
	f,err:=os.Create("爬虫/页面/第"+strconv.Itoa(i)+"页"+".html")
	//fmt.Println("result: ",result)
	if err != nil {
		fmt.Println("create err",err)
		return
	}
	f.WriteString(result)
	f.Close()  //保存一个 关一个1
	page<-i  //与主go程同步
}
//爬取
func working1(start,end int)  {
	fmt.Printf("正在爬取第%d页到%d页\n",start,end)
	page:=make(chan int)

	//循环爬取每一页
	for i := start; i <= end; i++ {
		go SpiderOne(i,page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 个网页爬取完毕\n",<-page)

	}
}
func main() {
	//指定起始页 终止页
	var start,end int
	fmt.Println("请输入爬去的起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入爬去的终止页：")
	fmt.Scan(&end)

	working1(start,end)
}
