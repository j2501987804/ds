package lists

import (
	"ds/containers"
	"ds/utils"
)

type List interface {
	Get(index int) (interface{},bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Sort(comparator utils.Comparator)
	Insert(index int,values ...interface{})
	Set(index int,value interface{})

	containers.Container
}
