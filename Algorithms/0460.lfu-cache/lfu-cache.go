package problem0460

import "container/list"

//// LFUCache 实现了 Least Frequently Used (LFU) cache
//type LFUCache struct {
//	// 用于检查 key 的存在性
//	m   map[int]*entry
//	pq  PQ
//	cap int
//}
//
//// Constructor 构建了 LFUCache
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		m:   make(map[int]*entry, capacity),
//		pq:  make(PQ, 0, capacity),
//		cap: capacity,
//	}
//}
//
//// Get 获取 key 的值
//func (c *LFUCache) Get(key int) int {
//	ep, ok := c.m[key]
//	if ok {
//		c.pq.update(ep)
//		return ep.value
//	}
//	return -1
//}
//
//// Put 把 key， value 成对地放入 LFUCache
//// 如果 LFUCache 已满的话，会删除掉 LFUCache 中使用最少的 key
//func (c *LFUCache) Put(key int, value int) {
//	if c.cap <= 0 {
//		return
//	}
//	ep, ok := c.m[key]
//	// key 已存在，就更新 value
//	if ok {
//		c.m[key].value = value
//		c.pq.update(ep)
//		return
//	}
//
//	ep = &entry{key: key, value: value}
//	// pq 已满的话，需要先删除，再插入
//	if len(c.pq) == c.cap {
//		temp := heap.Pop(&c.pq).(*entry)
//		delete(c.m, temp.key)
//	}
//	c.m[key] = ep
//	heap.Push(&c.pq, ep)
//}
//
///**
// * Your LFUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
//
//// entry 是 priorityQueue 中的元素
//type entry struct {
//	key, value int
//	// 以下是辅助参数，由 heap.Interface 实现的函数自动处理
//	frequence, index int
//	date             time.Time
//}
//
//// PQ implements heap.Interface and holds entries.
//type PQ []*entry
//
//func (pq PQ) Len() int { return len(pq) }
//
//func (pq PQ) Less(i, j int) bool {
//	if pq[i].frequence == pq[j].frequence {
//		return pq[i].date.Before(pq[j].date)
//	}
//
//	return pq[i].frequence < pq[j].frequence
//}
//
//func (pq PQ) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//// Push 往 pq 中放 entry
//func (pq *PQ) Push(x interface{}) {
//	n := len(*pq)
//	entry := x.(*entry)
//	entry.index = n
//	entry.date = time.Now()
//	*pq = append(*pq, entry)
//}
//
//// Pop 从 pq 中取出最优先的 entry
//func (pq *PQ) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	entry := old[n-1]
//	entry.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return entry
//}
//
//// update modifies the priority of an entry in the queue.
//func (pq *PQ) update(entry *entry) {
//	entry.frequence++
//	entry.date = time.Now()
//	heap.Fix(pq, entry.index)
//}
type LFUCache struct {
	Cap    int
	Keym   map[int]*list.Element //通过Key查找
	Cntm   map[int]*list.List    //通过次数查找
	MinCnt int                   //最小的访问次数，只有在put需要删除项目，且缓存容量达上限时才用的上，新的项目cnt必然是1，所以只要记录最小的即可
}

type node struct {
	Key   int
	Value int
	Cnt   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Cap:    capacity,
		Keym:   make(map[int]*list.Element, capacity),
		Cntm:   make(map[int]*list.List),
		MinCnt: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	e, ok := this.Keym[key]
	if ok { //存在Key，则更新Key对应的访问次数，并更新Cntm和MinCnt
		this.Cntm[e.Value.(*node).Cnt].Remove(e)
		if this.Cntm[e.Value.(*node).Cnt].Len() <= 0 { //对应访问次数的列表被清空时，应删除Cntm的元素
			delete(this.Cntm, e.Value.(*node).Cnt)
			if this.MinCnt == e.Value.(*node).Cnt {
				this.MinCnt++ //如果被删除的元素正好是最小次数，则最小次数++
			}
		}
		e.Value.(*node).Cnt++
		e2, ok := this.Cntm[e.Value.(*node).Cnt]
		if ok {
			this.Keym[key] = e2.PushFront(e.Value)
		} else {
			l := list.New()
			this.Keym[key] = l.PushFront(e.Value)
			this.Cntm[e.Value.(*node).Cnt] = l
		}
		return e.Value.(*node).Value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.Cap <= 0 {
		return
	}
	e, ok := this.Keym[key]
	if ok { //存在Key，则更新Key对应的Value和访问次数，并更新Cntm和MinCnt
		this.Cntm[e.Value.(*node).Cnt].Remove(e)
		if this.Cntm[e.Value.(*node).Cnt].Len() <= 0 { //对应访问次数的列表被清空时，应删除Cntm的元素
			delete(this.Cntm, e.Value.(*node).Cnt)
			if this.MinCnt == e.Value.(*node).Cnt {
				this.MinCnt++ //如果被删除的元素正好是最小次数，则最小次数++
			}
		}
		e.Value.(*node).Cnt++
		e.Value.(*node).Value = value
		e2, ok := this.Cntm[e.Value.(*node).Cnt]
		if ok {
			this.Keym[key] = e2.PushFront(e.Value)
		} else {
			l := list.New()
			this.Keym[key] = l.PushFront(e.Value)
			this.Cntm[e.Value.(*node).Cnt] = l
		}
	} else { //不存在Key，则新增Key，并更新Cntm和MinCnt
		for this.Cap <= len(this.Keym) { //容量不足，得删除掉最不经常使用的
			e := this.Cntm[this.MinCnt].Back()
			this.Cntm[this.MinCnt].Remove(e)
			if this.Cntm[this.MinCnt].Len() <= 0 { //&& this.MinCnt != 1 {
				delete(this.Cntm, this.MinCnt) //this.MinCnt为1时不需要删除，因为下面马上就要添加了
			}
			delete(this.Keym, e.Value.(*node).Key)
		}
		this.MinCnt = 1 //毋庸置疑
		n := node{key, value, 1}
		e2, ok := this.Cntm[1]
		if ok {
			this.Keym[key] = e2.PushFront(&n)
		} else {
			l := list.New()
			this.Keym[key] = l.PushFront(&n)
			this.Cntm[1] = l
		}
	}
}
