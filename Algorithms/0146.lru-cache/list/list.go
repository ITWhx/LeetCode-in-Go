/**
*
*@author 吴昊轩
*@create 2019-09-3013:52
 */
package list

type Element struct {
	pre, next *Element
	Value     interface{}
}
type List struct {
	root *Element
	len  int
}
