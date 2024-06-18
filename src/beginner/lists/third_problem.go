//Find the N'th element of a list

package main

import "fmt"

func nth(list []any, position int) any {
	if len(list) == 0 || len(list) < position { return nil;}
	return list[position];
}

func main() {
	fmt.Println(nth([]any{"a", "b", "c", "d", "e"}, 2));	
	fmt.Println(nth([]any{"a"}, 2));
}
