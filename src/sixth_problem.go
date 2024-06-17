// Find out whether a list is a palindrome.

package main

import "fmt"

func is_palindrome(list []any) bool {
	mid := len(list) / 2;
	for i := 0; i < mid; i++ {
			if list[i] != list[len(list) - 1 - i] {
				return false;
			}
		}
	return true;
}


func main() {
	fmt.Println(is_palindrome([]any{"x", "a", "m", "a", "x"}));
	fmt.Println(!is_palindrome([]any{"a", "b"}));
}
