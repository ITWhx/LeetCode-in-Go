/**
*
*@author 吴昊轩
*@create 2019-09-3013:52
 */
package list

type Element struct {
	pre, next *Element
	list      *List
	Value     interface{}
}
type List struct {
	root Element
	len  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.pre = &l.root
	l.len = 0
	return l
}
func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.pre
}
func (l *List) insert(e, at *Element) *Element {
	e.list = l
	n := at.next
	at.next = e
	e.pre = at
	e.next = n
	n.pre = e
	l.len++
	return e
}

func (l *List) insertValue(value interface{}, at *Element) *Element {
	return l.insert(&Element{Value: value}, at)
}

func (l *List) remove(e *Element) *Element {
	e.pre.next = e.next
	e.next.pre = e.pre
	e.pre = nil
	e.next = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.next.pre = e.pre
	e.pre.next = e.next

	n := at.next
	at.next = e
	e.pre = at
	e.next = n
	n.pre = e
	return e
}

func (l *List) Remove(e *Element) *Element {
	if e.list == l {
		l.remove(e)
	}
	return e
}

func (l *List) PushFront(value interface{}) *Element {
	l.lazyInit()
	return l.insertValue(value, &l.root)
}
func (l *List) PushBack(value interface{}) *Element {
	l.lazyInit()
	return l.insertValue(value, l.root.pre)
}

func (l *List) InsertBefore(e, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insert(e, mark.pre)
}

func (l *List) InsertAfter(e, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insert(e, mark)
}

func (l *List) MoveToBefore(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}
