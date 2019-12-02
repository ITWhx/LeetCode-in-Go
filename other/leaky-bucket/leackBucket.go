/**
*
*@author 吴昊轩
*@create 2019-11-2916:13
 */
package leaky_bucket

import (
	"sync/atomic"
	"time"
)

type LeakyBucket struct {
	cap         int64 // 桶的容量
	rate        int64 // 漏桶出水速度
	refreshTime int64
	water       int64 // 当前水量(当前累积请求数)
}

func (l *LeakyBucket) RateLimit() bool {
	now := time.Now().Unix()
	// 先执行漏水，计算剩余水量
	l.water = max(0, atomic.AddInt64(&l.water, -1*(now-l.refreshTime)*l.rate))
	l.refreshTime = now
	if l.water < l.cap {
		// 水桶还没满，继续加 1
		atomic.AddInt64(&l.water, 1)
		return true
	} else {
		// 水满，拒绝流入
		return false
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type TokenBucket struct {
	Capacity    int64 // 桶的容量
	Rate        int64 // 令牌放入速度
	Tokens      int64 // 当前令牌数量
	RefreshTime int64
}

func (t *TokenBucket) rateLimit() bool {
	now := time.Now().Unix()

	// 先添加令牌
	t.Tokens = min(t.Capacity, atomic.AddInt64(&t.Tokens, (now-t.RefreshTime)*t.Rate))
	t.RefreshTime = now

	if t.Tokens < 1 {
		// 令牌数小于 1 拒绝请求
		return false
	} else {
		// 还有令牌，领取令牌
		atomic.AddInt64(&t.Tokens, -1)
		return true
	}
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
