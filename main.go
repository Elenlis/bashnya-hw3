package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = nil
}

func main() {
	stack := &Stack{}

	fmt.Println("Введите числа через пробел для добавления в стек:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Fields(input)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Пропущено некорректное значение: %s\n", part)
			continue
		}
		stack.Push(num)
	}

	fmt.Printf("Данные введены. Размер стека: %d\n", stack.Size())
	if val, ok := stack.Pop(); ok {
		fmt.Printf("Вытолкнули: %d\n", val)
	} else {
		fmt.Println("Стек пуст, выталкивать нечего.")
	}
}
