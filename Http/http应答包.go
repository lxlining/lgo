package main

import "net/http"

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 16:51
 **/
//web回调函数
func Handler(w http.ResponseWriter,r *http.Request)  {
	//w 写给客户端（浏览器）的数据
	//r 从客户端 独到的数据
	w.Write([]byte("hello brow"))
}
func main() {
	//注册回调函数
	http.HandleFunc("/tt",Handler)
	//绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000",nil)

}
