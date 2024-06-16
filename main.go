// 99 Problems in Go, from 99 Problems in OCaml

package main

import "fmt"

// First problem: write a function that returns the last element of a list
func last(list []any) any {
	if (len(list) == 0) {
		return nil
	}
	return list[len(list)-1]
}

func main() {
	fmt.Println("Hello, World!");
	fmt.Println(last([]any{1,3,4}));
	fmt.Println(last([]any{}));
	fmt.Println(last([]any{"a","b","c"}));
}
