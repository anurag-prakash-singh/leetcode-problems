package ds

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	list := NewEmptyLinkedNode()

	if !list.IsEmpty() {
		t.Fatalf("Expected list to be empty")
	}
}

func TestAddAndTraverse(t *testing.T) {
	head := NewEmptyLinkedNode()

	head.AddNode(NewLinkedNode(1))
	head.AddNode(NewLinkedNode(2))

	visitorFn := func(d interface{}) {
		num, ok := d.(int)

		if !ok {
			t.Fatalf("Unexpected type")
		}

		t.Logf("Data: %d\n", num)
	}

	head.Traverse(visitorFn)
}

func TestDeleteAllNodes(t *testing.T) {
	head := NewEmptyLinkedNode()

	node1 := NewLinkedNode(1)
	node2 := NewLinkedNode(2)

	head.AddNode(node1)
	head.AddNode(node2)

	DeleteNode(node2)

	// visitorFn := func(d interface{}) {
	// 	num, ok := d.(int)

	// 	if !ok {
	// 		t.Fatalf("Unexpected type")
	// 	}

	// 	t.Logf("Data: %d\n", num)
	// }

	DeleteNode(node1)

	if !head.IsEmpty() {
		t.Fatalf("Expected list to be empty but it's not")
	}
}
