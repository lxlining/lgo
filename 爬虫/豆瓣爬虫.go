package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 21:04
 **/

//爬取指定页面
func httpgets(url string) (result string , err error) {
	res,err1:=http.Get(url)
	if err1 != nil {
		err=err1
		return
	}

	defer res.Body.Close()
	//循环爬取一页
	buf:=make([]byte,4096)
	for{
		n,err2:=res.Body.Read(buf)
		if n==0{
			break
		}
		if err2 != nil && err2!=io.EOF{
			err=err2
			return
		}
		fmt.Println(string(buf[:n]))
		result+=string(buf[:n])
	}
	return
}
func save2file(i int,filename,filescore,peoplenum [][]string)  {
	f,err:=os.Create("爬虫/页面/"+strconv.Itoa(i)+"页"+".txt")
	if err != nil {
		fmt.Println("create file err",err)
		return
	}
	defer f.Close()
	n:=len(filename)  //得到条目数
	//打印抬头
	f.WriteString("电影名称"+"\t\t\t"+"电影评分"+"\t\t\t"+"评分人数"+"\n")
	for j := 0; j < n; j++ {
		f.WriteString(filename[j][1]+"\t\t\t"+filescore[j][1]+"\t\t\t"+peoplenum[j][1]+"\n")
	}}
func SpiderPage(i int,page chan int)  {
	url:="https://movie.douban.com/top250?start="+strconv.Itoa((i-1)*25)+"&filter="
	res,err:=httpgets(url)
	if err != nil {
		fmt.Println("get err",err)
		return
	}

	//解析编译正则表达式  电影名称
	ret:=regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	//提取信息
	filename:=ret.FindAllStringSubmatch(res,-1)

	//解析编译正则表达式  评分
	ret=regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	//提取信息
	filescore:=ret.FindAllStringSubmatch(res,-1)

	//解析编译正则表达式 评分人数
	ret=regexp.MustCompile(`<span>(.*?)人评价</span>`)
	//提取信息
	peoplenum:=ret.FindAllStringSubmatch(res,-1)

	//将信息封装到文件中
	save2file(i,filename,filescore,peoplenum)
	//同步主go程
	page<-i
}
func toWork(start,end int)  {
	page:=make(chan int)
	fmt.Printf("正在爬取%d到%d页",start,end)
	for i := start; i <=end; i++ {
		go SpiderPage(i,page)
	}
	for i := start; i <=end; i++ {
		fmt.Println("第",<-page ,"页，读取完毕")
	}
}

func main() {
	//指定起始页 终止页
	var start,end int
	fmt.Println("请输入爬去的起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入爬去的终止页：")
	fmt.Scan(&end)

	toWork(start,end)
}
