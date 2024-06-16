// Second problem: find the last two elements of a list

package main

import "fmt"

func last_two(list []any) (any, any) {

	if (len(list) == 0 || len(list) == 1) {
		return nil, nil
	}
	return list[len(list)-2], list[len(list)-1]
}

func main() {
	fmt.Println(last_two([]any{"a", "b", "c", "d"}));
	fmt.Println(last_two([]any{"a"}));
}
