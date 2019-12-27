package main

import (
	"fmt"
	"net"
	"time"
)

/**
 * Created by 吴昊轩 on 2019/12/2.
 */

func subsets(nums []int) [][]int {
	res := [][]int{}
	num := 1 << uint32(len(nums))
	for i := 0; i < num; i++ {
		tmp := []int{}
		idx := 0
		iCopy := i
		for iCopy != 0 {
			if iCopy&1 == 1 {
				tmp = append(tmp, nums[idx])
			}
			idx++
			iCopy >>= 1
		}
		res = append(res, tmp)
	}
	return res
}
func main() {
	addr, e := net.ResolveTCPAddr("tcp4", "localhost:8888")
	if e != nil {
		panic(e)
	}
	listener, e := net.ListenTCP("tcp", addr)
	if e != nil {
		panic(e)
	}

	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("server handle conn error:%v", e)
		}

	}
}

func handle(conn net.Conn) {

}

type TokenBucket struct {
	Cap         int64
	Token       int64
	Rate        int64
	RefreshTime int64
}

func (t *TokenBucket) RateLimit() bool {
	now := time.Now().Unix()

}
