package solution

// you can also use imports, for example:
import (
	"errors"
	"fmt"
	"strconv"
)

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(expression string) int {
	// write your code in Go 1.4
	stack := Stack{}
	fmt.Println(expression)
	for _, element := range expression {
		value, err := strconv.Atoi(string(element))

		if err != nil {
			operand1, err := stack.Pop()
			if err != nil {
				return -1
			}
			operand2, err := stack.Pop()
			if err != nil {
				return -1
			}

			operation := strconv.QuoteRuneToASCII(element)
			if operation == "'+'" {
				total := operand1 + operand2
				stack.Push(total)
			} else if operation == "'*'" {
				total := operand1 * operand2
				stack.Push(total)
			} else {
				return -1
			}
		} else {
			stack.Push(value)
		}
	}

	finalValue, err := stack.Pop()
	if err != nil {
		return -1
	} else {
		return finalValue
	}
}

type Stack []int

func (s *Stack) Push(value int) {
	//fmt.Println(value)
	*s = append(*s, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Can't pop from an empty stack.")
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret, nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
