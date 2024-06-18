// Flatten a nested list structure
package main

import "fmt" 
import "reflect"

func flatten(list []interface{}) []interface{} {
	var r_list []interface{}
	for _, elem := range list {
		elem_type := reflect.TypeOf(elem);
		switch elem_type.Kind() {
		case reflect.Slice:
			r_list = append(r_list, flatten(elem.([]interface{}))...)
		default:
			r_list = append(r_list, elem)
		}
	}
	return r_list
}

func main() {
	fmt.Println(flatten([]any{"a", []any{"b", []any{"c", "d"}}, "e"}));
}
