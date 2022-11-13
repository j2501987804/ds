package arraylist

import (
	"ds/lists"
	"ds/utils"
	"fmt"
	"strings"
)

var _ lists.List = (*List)(nil)

type List struct {
	size     int
	elements []interface{}
}

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.checkIndex(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.checkIndex(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
}

func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for i := 0; i < len(values); i++ {
		list.elements[list.size] = values[i]
		list.size++
	}
}

func (list *List) Contains(values ...interface{}) bool {
	for _, v := range values {
		var found bool
		for i := 0; i < list.size; i++ {
			if v == list.elements[i] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.checkIndex(index) {
		if index == list.size {
			list.Add(values...)
			return
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

func (list *List) Set(index int, value interface{}) {
	if !list.checkIndex(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

func (list *List) Empty() bool {
	return list.size == 0
}
func (list *List) Size() int {
	return list.size
}
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *List) Values() []interface{} {
	return list.elements
}

func (list *List) String() string {
	values := make([]string, 0, list.size)
	for i := 0; i < list.size; i++ {
		values = append(values, fmt.Sprintf("%v", list.elements[i]))
	}
	return "Arraylist \n" + strings.Join(values, ", ")
}

// 缩容
// 如果1/4的容量大于长度就进行缩小
func (list *List) shrink() {
	currentCap := cap(list.elements)
	if list.size <= int(float32(currentCap)*0.25) {
		list.resize(list.size)
	}
}

// 2倍扩容
func (list *List) growBy(n int) {
	currentCap := len(list.elements)
	if list.size+n >= currentCap {
		newCap := 2 * (currentCap + n)
		list.resize(newCap)
	}
}

func (list *List) checkIndex(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Swap swaps the two values at the specified positions.
func (list *List) Swap(i, j int) {
	if list.checkIndex(i) && list.checkIndex(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}
