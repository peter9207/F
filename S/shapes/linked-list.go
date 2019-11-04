package shapes

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	val  int64
	next *Node
}

func NewLinkedList() (l *LinkedList) {
	return LinkedList{}
}

func (l *LinkedList) AddFirst(val int64) {

	node := Node{val: val, next: l.head}
	l.head = node

}

func (l *LinkedList) AddLast(val int64) {

	node := Node{val: val, next: nil}
	l.tail.next = &node
	l.tail = &node

}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) RemoveFirst() (result int64) {

	result = l.head.val
	l.head = l.head.next
	return
}