package ds

import (
	"errors"
)

type LinkedNode struct {
	value interface{}
	next  *LinkedNode
	prev  *LinkedNode
}

func (l *LinkedNode) IsEmpty() bool {
	if l.next == nil && l.prev == nil {
		return true
	}

	if l.next == l && l.prev == l {
		return true
	}

	return false
}

func (l *LinkedNode) AddNode(newNode *LinkedNode) {
	if l.IsEmpty() {
		l.next = newNode
		l.prev = newNode
		newNode.prev = l
		newNode.next = l
	} else {
		currFirstNode := l.next
		l.next = newNode
		newNode.next = currFirstNode
		newNode.prev = l
		currFirstNode.prev = newNode
	}
}

// Should only be called from the head of the list (the head node is essentially the sentinel)
func (l *LinkedNode) Traverse(nodeVisitorFn func(data interface{})) {
	for curr := l.next; curr != l; curr = curr.next {
		nodeVisitorFn(curr.value)
	}
}

func DeleteNode(l *LinkedNode) error {
	if l.IsEmpty() {
		return errors.New("DeleteNode called on an empty list")
	}

	next := l.next
	prev := l.prev

	next.prev = prev
	prev.next = next

	return nil
}

func NewLinkedNode(data interface{}) *LinkedNode {
	return &LinkedNode{value: data}
}

func NewEmptyLinkedNode() *LinkedNode {
	return &LinkedNode{}
}
