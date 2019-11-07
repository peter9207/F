package shapes

type LinkedList struct {
	Length int64
	head   *Node
	tail   *Node
}

type Node struct {
	val  float64
	next *Node
}

func NewLinkedList() (l *LinkedList) {
	return &LinkedList{}
}

func (l *LinkedList) AddFirst(val float64) {

	l.Length = l.Length + 1
	node := Node{val: val, next: l.head}
	l.head = &node

}

func (l *LinkedList) AddLast(val float64) {
	l.Length = l.Length + 1

	node := Node{val: val, next: nil}
	l.tail.next = &node
	l.tail = &node
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) RemoveFirst() (result float64) {

	l.Length = l.Length - 1
	result = l.head.val
	l.head = l.head.next
	return
}
