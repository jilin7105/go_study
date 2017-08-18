package main 
import "fmt"

func main()	{
	a := 10
	if a < 20 {
		fmt.Printf("a 小于 20\n")
		
	}
	fmt.Printf("a 的值为 ：%d\n", a)
	
	
	if a > 20{
		fmt.Printf("a 大于 20\n")
	} else {
		fmt.Printf("a 小于 20\n")
	}
	
}