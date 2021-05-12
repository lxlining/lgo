package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 22:01
 **/

//https://m.pengfu.com/xiaohua_1.html
//<h1 class="f18"><a href="https://m.pengfu.com/content/1857275/" title="讨价还价">讨价还价</a></h1>
//<h1 class="f18"><a href="https://m.pengfu.com/content/1857275/" title="讨价还价">讨价还价</a></h1>
//<a class="dp-b" href="https://m.pengfu.com/u/3428263">紫由</a>
//<div class="con-txt">
//<div class="con-txt">			</div>

//获取一个网页所有内容
func Httpgets(url string) (result string , err error) {
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
		//fmt.Println(string(buf[:n]))
		result+=string(buf[:n])
	}
	return
}

func SpiderJPage(url string)(title,content string,err error)  {
	res, err1 := Httpgets(url)
	if err1 != nil {
		err=err1
		return
	}

	//解析编译正则表达式  title
	ret := regexp.MustCompile(`<h1 class="f18">(.*?)`)
	//提取信息
	alls := ret.FindAllStringSubmatch(res, 1)
	for _,jurl:=range alls{
		//fmt.Println("jurl--",jurl[1])
		title,content,err=SpiderJPage(jurl[1])
		if err != nil {
			fmt.Println("spiderJPage err",err)
			continue
		}
	}
	return
}
func SpiderPage1(i int/*,page chan int*/) {
	url := "https://m.pengfu.com/xiaohua_"+strconv.Itoa((i)*10)+".html"
	res, err1 := Httpgets(url)
	if err1 != nil {
		//err=err1
		return
	}

	//解析编译正则表达式  电影名称
	ret := regexp.MustCompile(`<h1 class="f18"><a href="(?s:(.*?))"`)
	//提取信息
	alls := ret.FindAllStringSubmatch(res, -1)
	for _,jurl:=range alls{
		//fmt.Println("jurl--",jurl[1])
		title,content,err:=SpiderJPage(jurl[1])
		if err != nil {
			fmt.Println("spiderJPage err",err)
			continue
		}
		fmt.Println(title,content)
	}
	/*//解析编译正则表达式  评分
	ret = regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	//提取信息
	filescore := ret.FindAllStringSubmatch(res, -1)

	//解析编译正则表达式 评分人数
	ret = regexp.MustCompile(`<span>(.*?)人评价</span>`)
	//提取信息
	peoplenum := ret.FindAllStringSubmatch(res, -1)

	//将信息封装到文件中
	save2file(i, filename, filescore, peoplenum)*/
	//同步主go程
	/*page <- i*/
}
func ToWork(start,end int)  {
	//page:=make(chan int)
	fmt.Printf("正在爬取%d到%d页",start,end)
	for i := start; i <=end; i++ {
		SpiderPage1(i/*,page*/)
	}
	/*for i := start; i <=end; i++ {
		fmt.Println("第",<-page ,"页，读取完毕")
	}*/
}
func main() {
	//指定起始页 终止页
	var start,end int
	fmt.Println("请输入爬去的起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入爬去的终止页：")
	fmt.Scan(&end)

	ToWork(start,end)
}

/*

package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
)

// 获取一个网页所有的内容， result 返回
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

// 抓取一个网页，带有 10 个段子 —— 10 URL
func SpiderPage(idx int)  {
	// 拼接URL
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(idx) + ".html"

	// 封装函数获取段子的URL
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	// 解析、编译正则
	ret := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)

	// 提取需要信息 —— 每个段子的 URL
	alls := ret.FindAllStringSubmatch(result, -1)

	for _, jokeURL := range alls {
		fmt.Println("jokeURL:", jokeURL[1])
	}
}

func toWork(start, end int)  {
	fmt.Printf("正在爬取 %d 到 %d 页...\n", start, end)

	for i:=start; i<=end; i++ {
		SpiderPage(i)
	}
}

func main()  {
	// 指定爬取起始、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)

	toWork(start, end)
}

 */