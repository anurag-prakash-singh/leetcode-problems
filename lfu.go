package main

import (
	"errors"
	"fmt"
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

type LFUCache struct {
	capacity int
	size     int
	freqList *LinkedNode
	keyMap   map[int]valueDesc // int -> valueDesc
}

type valueDesc struct {
	valueNode    *LinkedNode // Node containing the actual value
	freqListNode *LinkedNode // Node from the freq list containing the value node
}

type freqData struct {
	freq        int
	valuesNodes *LinkedNode
}

type keyValueData struct {
	key   int
	value int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity: capacity, freqList: NewEmptyLinkedNode(), keyMap: make(map[int]valueDesc)}
}

func (this *LFUCache) Get(key int) int {
	vdesc, ok := this.keyMap[key]

	if !ok {
		return -1
	}

	valueFreqNode := vdesc.freqListNode
	valueFreqData := valueFreqNode.value.(freqData)
	nextFreqNode := valueFreqNode.next

	nextFreqData := freqData{}
	if nextFreqNode.value != nil {
		nextFreqData = nextFreqNode.value.(freqData)
	}

	var higherFreqNode *LinkedNode

	if nextFreqNode == this.freqList || nextFreqData.freq != valueFreqData.freq+1 {
		// Either:
		// - this particular freq node is the last one, or
		// - the next higher freq node is for higher frequency than we need
		// For both cases, create a new freq node.
		newFreqData := freqData{freq: valueFreqData.freq + 1, valuesNodes: NewEmptyLinkedNode()}
		newFreqNode := NewLinkedNode(newFreqData)
		valueFreqNode.AddNode(newFreqNode)
		higherFreqNode = newFreqNode
	} else if nextFreqData.freq == valueFreqData.freq+1 {
		higherFreqNode = nextFreqNode
	} else {
		panic("freqNode conditions not met")
	}

	// Delete the value node from its current list and add it to the new freq node's value node list
	DeleteNode(vdesc.valueNode)
	higherFreqData := higherFreqNode.value.(freqData)
	higherFreqData.valuesNodes.AddNode(vdesc.valueNode)

	// We also need to update the freq node since it has changed
	vdesc.freqListNode = higherFreqNode
	this.keyMap[key] = vdesc

	// If the old freq node's value node list is empty, we delete the freq node itself
	if valueFreqData.valuesNodes.IsEmpty() {
		DeleteNode(valueFreqNode)
	}

	return vdesc.valueNode.value.(keyValueData).value
}

func (this *LFUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		vdesc, ok := this.keyMap[key]

		if !ok {
			panic(fmt.Sprintf("We should have found the %d in keyMap", key))
		}

		kvData := vdesc.valueNode.value.(keyValueData)
		kvData.value = value
		vdesc.valueNode.value = kvData

		return
	}

	// if ok {
	// 	// Update this value. No changes needed to freqList
	// 	kvData := vdesc.valueNode.value.(keyValueData)
	// 	kvData.value = value
	// 	vdesc.valueNode.value = kvData

	// 	return
	// }

	if this.capacity == 0 {
		return
	}

	// TODO: remove node if we're at capacity
	if this.size == this.capacity {
		// Delete one node to make space
		firstFreqNode := this.freqList.next
		firstFreqData := firstFreqNode.value.(freqData)
		firstFreqValuesNodes := firstFreqData.valuesNodes

		// Delete the last (oldest) node in the valuesNodes pointed at by this freqNode
		valueNodeToDelete := firstFreqValuesNodes.prev
		DeleteNode(valueNodeToDelete)

		if firstFreqValuesNodes.IsEmpty() {
			DeleteNode(firstFreqNode)
		}

		delete(this.keyMap, valueNodeToDelete.value.(keyValueData).key)
		this.size--
	}

	newVDesc := valueDesc{}
	newValueNode := NewLinkedNode(keyValueData{key: key, value: value})
	newVDesc.valueNode = newValueNode

	// key not found in map:
	// check if there's a freq node that has never been accessed. If
	// there's one, the new value will go into that freq node's list.
	// Otherwise, we'll need to create a new freq node and add this new
	// value node to that freq node's list.
	if this.freqList.IsEmpty() {
		newValuesNodes := NewEmptyLinkedNode()
		newValuesNodes.AddNode(newValueNode)
		newFreqData := freqData{freq: 0, valuesNodes: newValuesNodes}
		newFreqNode := NewLinkedNode(newFreqData)
		this.freqList.AddNode(newFreqNode)
		newVDesc.freqListNode = newFreqNode
	} else {
		// does the first node have freq 0? If so, we just add to that freqNode list
		firstFreqNode := this.freqList.next
		// We want to panic if the types are wrong
		firstFreqData := firstFreqNode.value.(freqData)

		if firstFreqData.freq == 0 {
			firstFreqValuesNodes := firstFreqData.valuesNodes
			firstFreqValuesNodes.AddNode(newValueNode)
			newVDesc.freqListNode = firstFreqNode
		} else {
			// We need to add new freq node with freq 0
			newValuesNodes := NewEmptyLinkedNode()
			newValuesNodes.AddNode(newValueNode)
			newFreqData := freqData{freq: 0, valuesNodes: newValuesNodes}
			newFreqNode := NewLinkedNode(newFreqData)
			this.freqList.AddNode(newFreqNode)
			newVDesc.freqListNode = newFreqNode
		}
	}

	this.keyMap[key] = newVDesc

	this.size++
}

func printFreqList(freqList *LinkedNode) {
	freqNodeVisitor := func(d interface{}) {
		fd := d.(freqData)

		fmt.Printf("freq: %d\n", fd.freq)

		valueNodeVisitor := func(d interface{}) {
			kvData := d.(keyValueData)

			fmt.Printf("\t%d -> %d\n", kvData.key, kvData.value)
		}

		fd.valuesNodes.Traverse(valueNodeVisitor)
	}

	freqList.Traverse(freqNodeVisitor)
}

func printValues(head *LinkedNode) {
	visitorFn := func(d interface{}) {
		num, ok := d.(int)

		if !ok {
			fmt.Printf("Unexpected type\n")
		}

		fmt.Printf("Data: %d\n", num)
	}

	head.Traverse(visitorFn)
}

func test1() {
	lfuCache := Constructor(2)

	lfuCache.Put(1, 1)
	lfuCache.Put(2, 2)
	lfuCache.Put(3, 3)

	// cache.put(1, 1);
	// cache.put(2, 2);
	// cache.get(1);       // returns 1
	// cache.put(3, 3);    // evicts key 2
	// cache.get(2);       // returns -1 (not found)
	// cache.get(3);       // returns 3.
	// cache.put(4, 4);    // evicts key 1.
	// cache.get(1);       // returns -1 (not found)
	// cache.get(3);       // returns 3
	// cache.get(4);

	fmt.Printf("lfuCache.Get(1) = %d\n", lfuCache.Get(1))
	fmt.Printf("lfuCache.Get(2) = %d\n", lfuCache.Get(2))
	fmt.Printf("lfuCache.Get(2) = %d\n", lfuCache.Get(2))
	fmt.Printf("lfuCache.Get(3) = %d\n", lfuCache.Get(3))
	fmt.Printf("lfuCache.Get(3) = %d\n", lfuCache.Get(3))

	// firstFreqNode := lfuCache.freqList.next
	// firstFreqData := firstFreqNode.value.(freqData)
	// printValues(firstFreqData.valuesNodes)

	printFreqList(lfuCache.freqList)
}

func test2() {
	lfuCache := Constructor(2)

	lfuCache.Put(1, 1)
	lfuCache.Put(2, 2)
	fmt.Printf("lfuCache.Get(1) = %d (expect 1)\n", lfuCache.Get(1))
	lfuCache.Put(3, 3)
	fmt.Printf("lfuCache.Get(2) = %d (expect -1)\n", lfuCache.Get(2))
	fmt.Printf("lfuCache.Get(3) = %d (expect 3)\n", lfuCache.Get(3))
	lfuCache.Put(4, 4)
	fmt.Printf("lfuCache.Get(1) = %d (expect -1)\n", lfuCache.Get(1))
	fmt.Printf("lfuCache.Get(3) = %d (expect 3)\n", lfuCache.Get(3))
	fmt.Printf("lfuCache.Get(4) = %d (expect 4)\n", lfuCache.Get(4))

	// cache.put(1, 1);
	// cache.put(2, 2);
	// cache.get(1);       // returns 1
	// cache.put(3, 3);    // evicts key 2
	// cache.get(2);       // returns -1 (not found)
	// cache.get(3);       // returns 3.
	// cache.put(4, 4);    // evicts key 1.
	// cache.get(1);       // returns -1 (not found)
	// cache.get(3);       // returns 3
	// cache.get(4);

	printFreqList(lfuCache.freqList)
}

func test3() {
	// ["LFUCache","put","put","put","put","get"]
	//[[2],[3,1],[2,1],[2,2],[4,4],[2]]
	lfuCache := Constructor(2)

	lfuCache.Put(3, 1)
	lfuCache.Put(2, 1)
	lfuCache.Put(2, 2)
	lfuCache.Put(4, 4)
	fmt.Printf("lfuCache.Get(2) = %d (expect 2)\n", lfuCache.Get(2))

	printFreqList(lfuCache.freqList)
}

func test4() {
	// ["LFUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	lfuCache := Constructor(2)

	lfuCache.Put(2, 1)
	lfuCache.Put(1, 1)
	lfuCache.Put(2, 3)
	lfuCache.Put(4, 1)
	// printFreqList(lfuCache.freqList)
	fmt.Printf("lfuCache.Get(1) = %d (expect -1)\n", lfuCache.Get(1))
	fmt.Printf("lfuCache.Get(2) = %d (expect 3)\n", lfuCache.Get(2))

	printFreqList(lfuCache.freqList)
}

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
