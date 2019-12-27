/**
*
*@author 吴昊轩
*@create 2019-10-2319:02
 */
package problem0460

import "container/list"

type LFUCache1 struct {
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

func Constructor1(capacity int) LFUCache1 {
	return LFUCache1{
		Cap:    capacity,
		Keym:   make(map[int]*list.Element, capacity),
		Cntm:   make(map[int]*list.List),
		MinCnt: 0,
	}
}

func (this *LFUCache1) Get(key int) int {
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
			l := new(list.List)
			this.Keym[key] = l.PushFront(e.Value)
			this.Cntm[e.Value.(*node).Cnt] = l
		}
		return e.Value.(*node).Value
	}
	return -1
}

func (this *LFUCache1) Put(key int, value int) {
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
