package main

import (
	"fmt"
	"net/http"
	"os"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 17:36
 **/
func openFile(fname string,w http.ResponseWriter)  {
	path:="./"+fname
	f,err:=os.Open(path)
	if err != nil {
		fmt.Println("open err",err)
		w.Write([]byte("文件不存在！"))
		return
	}
	defer f.Close()
	buf:=make([]byte,4096)
	for{
		n,_:=f.Read(buf)
		if n==0 {
			return
		}
		w.Write(buf[:n])
	}
}
func Handler2(w http.ResponseWriter,r *http.Request)  {
	//w 写给客户端（浏览器）的数据
	//r 从客户端 独到的数据
	fmt.Println("客户端请求：",r.URL)
	openFile(r.URL.String(),w)
}
func main() {
	http.HandleFunc("/",Handler2)
	http.ListenAndServe("127.0.0.1:8000",nil)
}
