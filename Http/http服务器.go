package main

import (
	"fmt"
	"net/http"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 17:21
 **/
func Handler1(w http.ResponseWriter,r *http.Request)  {
	//w 写给客户端（浏览器）的数据
	//r 从客户端 独到的数据
	w.Write([]byte("hello brow"))
	fmt.Println("addr--",r.RemoteAddr)
	fmt.Println("body--",r.Body)
}
func main() {
	http.HandleFunc("/tt",Handler1)
	http.ListenAndServe("127.0.0.1:8000",nil)
}
