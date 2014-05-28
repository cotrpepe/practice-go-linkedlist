package linked_list

import (
	"errors"
)

type node struct {
	value int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (list *LinkedList) Size() int {
	cnt := 0
	for node := list.head; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (list *LinkedList) Empty() bool {
	return list.head == nil
}

func (list *LinkedList) Add(value int) {
	node := new(node)
	node.value = value
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
}

func (list *LinkedList) Delete(value int) error {
	prev := new(node)
	for node := list.head; node != nil; node = node.next {
		if node.value == value {
			if node == list.head {
				list.head = node.next
			}
			prev.next = node.next
			return nil
		}
		prev = node
	}
	return errors.New("can't find deleting node")
}

