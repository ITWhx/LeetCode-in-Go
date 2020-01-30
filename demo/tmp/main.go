package main

import (
	"fmt"
	"time"
)

/**
 * Created by 吴昊轩 on 2019/12/29.
 */
var c = make(chan int)

func main() {

	for i := 0; i < 100; i++ {
		go print(i)
		c <- 0
	}
	time.Sleep(time.Second * 5)
}

func print(i int) {
	<-c
	fmt.Println(i)
}
