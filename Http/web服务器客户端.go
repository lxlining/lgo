package main

import (
	"fmt"
	"net/http"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/12 17:49
 **/
func main() {
	resp,err:=http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("get err",err)
		return
	}
	defer resp.Body.Close()
	//读应答包
	fmt.Println(resp.Header,"\n--",resp.Status,"\n--",resp.Proto)
	buf:=make([]byte,4096)
	var res string
	for{
		n,_:=resp.Body.Read(buf)
		if n==0 {
			fmt.Println("read finish!")
			break
		}
		res+=string(buf[:n])
	}
	fmt.Println(res)
}
