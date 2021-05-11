package main

import (
	"fmt"
	"net"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/5/11 22:00
 **/
func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("udp服务器创建完成")
	udpConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("udp通信完成")
	defer udpConn.Close()
	//读取客户端
	buf := make([]byte, 4096)
	for {
		// 返回 读取到的字节数 客户端地址 error
		n, cltAddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		//处理数据
		fmt.Println("服务器读到：", string(buf[:n]), "来自==", cltAddr)

		//回写
		daytime := time.Now().String()
		_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}
