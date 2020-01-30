package main

import "fmt"

/**
 * Created by 吴昊轩 on 2019/12/29.
 */

//如果结构体的所有成员变量都是可比较的，那么结构体就可比较
//如果结构体中存在不可比较的成员变量，那么结构体就不能比较
//结构体之间进行转换需要他们具备完全相同的成员(字段名、字段类型、字段个数)
type T1 struct {
	Num  int
	Name string
	ptr  *string
	ptr2 *int
}

type T2 struct {
	Num  int
	Name string
}

type T3 struct {
	Num  int
	Name string
	S    []int //slice,map是不能比较的成员变量

}

func main() {
	t1 := T1{2, "wgx", new(string), new(int)}
	t2 := T1{2, "wgx", new(string), new(int)}
	fmt.Printf("T1 %t\n", t1 == t2)

	t3 := T2{2, "wgx"}
	t4 := T2{2, "wgx"}
	fmt.Printf("T2 %t\n", t3 == t4)

	t5 := T3{2, "whx", make([]int, 2)}
	t6 := T3{2, "whx", make([]int, 2)}
	fmt.Printf("T3 %t\n", t5 == t6)

}
