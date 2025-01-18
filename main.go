package main

import (
	"fmt"
	"sync"
)

// Stack представляет стек с дженериками
type Stack[T any] struct {
	elements map[int]T
	top      int
	mu       sync.Mutex
}

// NewStack создает новый стек
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make(map[int]T),
		top:      -1, // Изначально стек пуст
	}
}

// Push добавляет элемент на вершину стека
func (s *Stack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.top++
	s.elements[s.top] = value
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.top < 0 {
		var zeroValue T // Создаем значение нулевого типа
		return zeroValue, fmt.Errorf("стек пуст")
	}

	value := s.elements[s.top]
	delete(s.elements, s.top) // Удаляем элемент из карты
	s.top--

	return value, nil
}

// Peek возвращает элемент с вершины стека без удаления
func (s *Stack[T]) Peek() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.top < 0 {
		var zeroValue T // Создаем значение нулевого типа
		return zeroValue, fmt.Errorf("стек пуст")
	}

	return s.elements[s.top], nil
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.top < 0
}

// Size возвращает количество элементов в стеке
func (s *Stack[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.top + 1
}

func main() {
	stack := NewStack[bool]() // Создаем стек для целых чисел

	stack.Push(true)
	stack.Push(false)
	stack.Push(true)
	stack.Push(true)
	stack.Push(true)
	stack.Push(false)

	fmt.Println("Size of stack:", stack.Size())
	fmt.Println("Stack:", stack.elements)
	topElement, _ := stack.Peek()
	fmt.Println("Top element:", topElement)

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println("Popped element:", value)
	}

	fmt.Println("Size of stack after popping:", stack.Size())
}
