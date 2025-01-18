package main

import (
	"fmt"
	"sync"
)

type IntNode struct {
	Value int
	Next  *IntNode
}

func (n IntNode) GetNext() *IntNode {
	return n.Next
}

func NewIntNode(value int) *IntNode {
	return &IntNode{value, nil}
}

func NewIntList() *IntList {
	return &IntList{size: 0, Head: nil, mu: sync.Mutex{}}
}

type IntList struct {
	size int
	Head *IntNode
	mu   sync.Mutex
}

func (l *IntList) Size() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.size
}

func (l *IntList) Get(index int) (*IntNode, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 || index >= l.Size() {
		return nil, fmt.Errorf("неверный индекс списка")
	}
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node, nil
}

func (l *IntList) Add(el int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := NewIntNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		node := l.Head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = newNode
	}
	l.size++
}

func (l *IntList) Insert(el int, index int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 || index > l.size {
		return fmt.Errorf("неверный индекс списка")
	}
	newNode := NewIntNode(el)
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
	} else {
		node, err := l.Get(index - 1)
		if err != nil {
			return err
		}
		newNode.Next = node.Next
		node.Next = newNode
	}
	l.size++
	return nil
}

func (l *IntList) Remove(index int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 || index >= l.size {
		return fmt.Errorf("неверный индекс списка")
	}

	if index == 0 {
		l.Head = l.Head.Next
	} else {
		node := l.Head
		for i := 0; i < index-1; i++ {
			node = node.Next
		}
		node.Next = node.Next.Next
	}
	l.size--
	return nil
}

func (l *IntList) Print() {
	l.mu.Lock()
	defer l.mu.Unlock()

	node := l.Head
	if node != nil {
		for node != nil {
			fmt.Printf("%d\t", node.Value)
			node = node.Next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Список пуст!")
	}
}

func main() {
	list := NewIntList()

	var wg sync.WaitGroup

	// Adding elements concurrently
	wg.Add(4)
	go func() {
		defer wg.Done()
		fmt.Println("Adding 3")
		list.Add(3)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Adding 2")
		list.Add(2)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Inserting 5 at index 0")
		list.Insert(5, 0)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Inserting 0 at index 0")
		list.Insert(0, 0)
	}()

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Size after additions:", list.Size())
	list.Print()

	err := list.Remove(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Size after removal:", list.Size())
	list.Print()
}
