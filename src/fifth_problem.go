// Reverse a list

package main

import "fmt"

func rev(list []any) []any {
	var rev_list []any;
	for i := len(list)-1; i >= 0; i-- {
		rev_list = append(rev_list, list[i]);
	}
	return rev_list;
}

func main() {
	fmt.Println(rev([]any{"a", "b", "c"}));
}

