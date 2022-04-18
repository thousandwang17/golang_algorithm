package LinkList

type linkList struct {
	head *Node
	tail *Node
}

type Node struct {
	Val  interface{}
	pre  *Node
	next *Node
}

// FIFO data struct
// use two-way linked list when  need  Append func
// otherwise , one-way linked list is satisfy FIFO
func New() *linkList {
	return new(linkList)
}

func NewNode(val interface{}) *Node {
	n := new(Node)
	n.Val = val
	return n
}

func (l linkList) Append(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	n.pre = l.tail
	l.tail = n
}

func (l linkList) Prepend(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	n.next = l.head
	l.head.pre = n
	l.head = n
}

func (l linkList) Pop() (*Node, bool) {
	if l.tail == nil {
		return nil, false
	}

	n := l.tail
	pre := l.tail.pre
	l.tail = pre
	l.tail.next = nil
	return n, true
}

func (l linkList) Shift() (*Node, bool) {

	if l.head == nil {
		return nil, false
	}

	n := l.head
	next := l.head.next

	l.head = next
	l.head.pre = nil
	return n, true
}
