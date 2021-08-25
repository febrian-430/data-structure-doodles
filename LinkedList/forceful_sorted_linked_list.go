package linkedlist

import (
	"errors"
	"fmt"

	"github.com/febrian-430/data-structure-doodles/persistor"
)

type LinkedListNode struct {
	Val int
	//add iterable interface
	//so a tree can be in a linked list and vversa
	next *LinkedListNode
}

func (node *LinkedListNode) setNext(val int) error {
	old_next := node.next
	node.next = &LinkedListNode{
		Val:  val,
		next: old_next,
	}
	return nil
}

func (node *LinkedListNode) Next() *LinkedListNode {
	return node.next
}

var (
	EXCLUSIVE_FACTOR = 5
)

type ForcefullySortedLinkedList struct {
	head             *LinkedListNode
	tail             *LinkedListNode
	current          *LinkedListNode
	exclusive_factor int

	logger persistor.Persistor
}

func NewLinkedList() ForcefullySortedLinkedList {
	return ForcefullySortedLinkedList{
		head:             nil,
		tail:             nil,
		exclusive_factor: EXCLUSIVE_FACTOR,
		logger:           persistor.NewInMemory(),
	}
}

func (list *ForcefullySortedLinkedList) Head() *LinkedListNode {
	if list.Empty() {
		return nil
	}

	return list.head
}

func (list *ForcefullySortedLinkedList) Next() *LinkedListNode {
	if list.Empty() {
		return nil
	}
	if list.current == nil {
		list.current = list.head
		return list.head
	}
	list.current = list.current.next
	return list.current
}

func (list *ForcefullySortedLinkedList) Last() (int, error) {
	node, err := list.last()
	if err != nil {
		return -1, err
	}
	return node.Val, nil
}

func (list ForcefullySortedLinkedList) Empty() bool {
	return (list.head == nil)
}

func (list *ForcefullySortedLinkedList) Push(vals ...int) error {
	for _, val := range vals {
		err := list.push(val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (list *ForcefullySortedLinkedList) setHead(val int) error {
	if list.head != nil {
		return errors.New("Head already exists")
	}
	node := &LinkedListNode{
		Val: val,
	}
	list.head = node
	list.tail = node
	return nil
}

func (list *ForcefullySortedLinkedList) isExclusiveValue(val int) bool {
	return (val % list.exclusive_factor) == 0
}

func (list *ForcefullySortedLinkedList) push(val int) error {
	if list.Empty() {
		list.logger.Save(fmt.Sprintf("%v saved as head", val))
		return list.setHead(val)
	} else if list.isExclusiveValue(val) {
		list.logger.Save(fmt.Sprintf("%v is an exclusive value!", val))

		return list.pushExclusive(val)
	}
	err := list.pushNormal(val)
	if err != nil {
		list.logger.Save(fmt.Sprintf("%v, not today :-(", val))
	} else {
		list.logger.Save(fmt.Sprintf("%v, is tail now :-)", val))
	}
	return err
}

func (list *ForcefullySortedLinkedList) pushNormal(val int) error {
	last, err := list.last()
	if err != nil {
		list.setHead(val)
		return nil
	}
	if last.Val < val {
		list.setNextOf(last, val)
		return nil
	}
	return errors.New("YOU DONT BELONG HERE")
}

func (list *ForcefullySortedLinkedList) pushExclusive(val int) error {
	cur := list.head
	for cur != nil {
		if cur.Val < val {
			err := list.setNextOf(cur, val)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (list *ForcefullySortedLinkedList) setNextOf(node *LinkedListNode, val int) error {
	var err error
	err = node.setNext(val)
	if err != nil {
		return err
	}
	return list.recalculateTail(node)
}

func (list *ForcefullySortedLinkedList) recalculateTail(node *LinkedListNode) error {
	if node == list.tail && node.next != nil {
		list.tail = node.next
	}
	return nil
}

func (list *ForcefullySortedLinkedList) last() (*LinkedListNode, error) {
	if list.tail == nil {
		return nil, errors.New("Tail is nil")
	}
	return list.tail, nil
}
