// Find the number of elements of a list

package main

import "fmt"

func length(list []any) any {
	length_list := 0;
	for _ = range list {
		length_list++;
	}
	return length_list;
}

func main() {
	fmt.Println(length([]any{"a", "b", "c"}));
	fmt.Println(length([]any{}));
}
