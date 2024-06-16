// First problem: write a function that returns the last element of a list

package main

import "fmt"

func last(list []any) any {
	if (len(list) == 0) {
		return nil
	}
	return list[len(list)-1]
}

func main() {
	fmt.Println(last([]any{"a", "b", "c", "d"}));
	fmt.Println(last([]any{}));
}
