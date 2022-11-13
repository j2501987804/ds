package main

import (
	"ds/lists/arraylist"
	"ds/utils"
)

func main() {
	list := arraylist.New()
	list.Add("a")                         // ["a"]
	list.Add("c", "b")                    // ["a","c","b"]
	list.Sort(utils.StringComparator)     // ["a","b","c"]
	println(list.Get(0))                  // "a",true
	println(list.Get(100))                // nil,false
	println(list.Contains("a", "b", "c"))// true
	println(list.Contains("a", "b", "c", "d")) // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	println(list.Empty())                      // true
	println(list.Size())                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}
