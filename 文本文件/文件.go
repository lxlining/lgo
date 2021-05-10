package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * @Description 
 * @Author lixl
 * @Date 2021/5/9 11:48
 **/
//创建文件
func main01() {
	fp,err:=os.Create("文本文件/a.txt")
	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}
	defer fp.Close()
	fmt.Println("文件创建成功！")
	fp.WriteString("hahahhahahahhaha")
}
func main()  {
	//fp,err:=os.Open("文本文件/b.txt")  //只读打开
	fp,err:=os.OpenFile("文本文件/b.txt",os.O_RDWR,6)
	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}
	defer fp.Close()
	/*b:= []byte("哈哈哈哈")
	n,_:=fp.Seek(0,os.SEEK_END)
	fmt.Println(n)
	fp.WriteAt(b,n)//会覆盖原内容*/
	//字符串与字符切片可以相互转换
	/*str:="aaabbbccc"
	b=[]byte(str)*/
	b:=make([]byte,100)
	fp.Read(b)
	fmt.Println(string(b))
	for i := 0; i <len(b) ; i++ {
		if b[i]!=0{
			fmt.Printf("%c",b[i])
		}
	}
	//按行读取
	r:=bufio.NewReader(fp)
	c,_:= r.ReadBytes('\n')
	fmt.Println(111)
	fmt.Print(string(c))
	fmt.Println(111)

}