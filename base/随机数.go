package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

/**
 * @Description
 * @Author lixl
 * @Date 2021/4/20 20:58
 **/

func main() {
	rand2.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand2.Intn(1000))
	}

}
