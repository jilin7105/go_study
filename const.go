package main 

import "fmt"
import "unsafe"
func main() {

	const LENGTH  =  10
	const WIDTH  = 5
	var  area int 
	const a, b, c = 1, false, "str"
	
	const (
		get = iota
		s = "abc"
		d = len(s)
		f = unsafe.Sizeof(s)
		set = 3<<iota
	)
	
	
	
	
	area = LENGTH * WIDTH
	fmt.Printf("面积 ： %d", area)
	println()
	println(a, b, c)
	
	println(s, d, f)
	
	
	
	println(get,set)
}