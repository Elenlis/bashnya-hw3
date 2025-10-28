package main

import "fmt"

// Stack — стек целых чисел
type Stack struct {
	data []int
}

// Push добавляет элемент на вершину
func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

// Pop удаляет и возвращает верхний элемент
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size возвращает количество элементов
func (s *Stack) Size() int {
	return len(s.data)
}

// Clear очищает стек
func (s *Stack) Clear() {
	s.data = nil
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	fmt.Println("Размер:", stack.Size()) // 2
	if val, ok := stack.Pop(); ok {
		fmt.Println("Вытолкнули:", val) // 20
	}
}
