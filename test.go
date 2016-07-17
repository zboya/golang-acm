package main

import (
	"fmt"
)

func main() {
	var n int
	var str string
	var result string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		//match
		rst := matchPairBracket(&str)
		result = result + rst
	}
	fmt.Printf("----\n%s", result)
}

func matchPairBracket(str *string) string {
	var stack []byte
	r := []byte("()[]")
	strLen := len(*str) % 2
	strByte := []byte(*str)

	if strLen == 1 || r[0] != strByte[0] && r[2] != strByte[0] {
		return "NO\n"
	} else {
		for i := 0; i < len(*str); i++ {
			if r[0] == strByte[i] || r[2] == strByte[i] {
				push(&stack, strByte[i])
			} else {
				if r[1] == strByte[i] && r[0] == stack[len(stack)-1] {
					pop(&stack)
				} else if r[3] == strByte[i] && r[2] == stack[len(stack)-1] {
					pop(&stack)
				} else {
					push(&stack, strByte[i])
				}
			}
		}
		if len(stack) == 0 {
			return "YES\n"
		} else {
			return "NO\n"
		}
	}

}

func push(stack *[]byte, r byte) *[]byte {
	*stack = append(*stack, r)
	return stack

}

func pop(stack *[]byte) {
	*stack = (*stack)[:len(*stack)-1]
}
