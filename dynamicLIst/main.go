package main

import (
	"fmt"
)

type List[T comparable] struct {
	Items []T
}

func main() {

	list1 := List[int]{Items: []int{1, 2, 3, 4}}
	list1.Add(33)
	list1.InsertAt(2, 10)
	list1.removeItem(100)
	fmt.Println(list1)

}

func (list *List[T]) Add(item T) {
	list.Items = append(list.Items, item)
}

func (list *List[T]) InsertAt(index int, item T) {
	if index == len(list.Items) {
		list.Items = append(list.Items, item)
	} else {
		list.Items = append(list.Items[:index+1], list.Items[index:]...)
		list.Items[index] = item

	}
}

func (list *List[T]) find(item T) int {
	for idx, element := range list.Items {
		if element == item {
			return idx
		}
	}
	return -1
}

func (list *List[T]) removeItem(item T) {
	itemIndex := list.find(item)
	if itemIndex != -1 {

		list.Items = append(list.Items[:itemIndex], list.Items[itemIndex+1:]...)
	} else {
		fmt.Println("not item found")
	}
}
