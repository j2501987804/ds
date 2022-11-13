// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package containers provides core interfaces and functions for data structures.
//
// Container is the base interface for all data structures to implement.
//
// Iterators provide stateful iterators.
//
// Enumerable provides Ruby inspired (each, select, map, find, any?, etc.) container functions.
//
// Serialization provides serializers (marshalers) and deserializers (unmarshalers).
package containers

import "ds/utils"

// Container is base interface that all data structures implement.
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}

// 取出排序容器内的元素
// 容器内原来的排序不影响
func GetSortedValues(container Container, comparator utils.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	utils.Sort(values, comparator)
	return values
}
