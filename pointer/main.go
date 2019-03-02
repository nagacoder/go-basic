package main 
import "fmt"

func  main()  {
	a := 5
	b := &a

	fmt.Println(b)
	// use '*' to get value of var b
	fmt.Println(*b)
	// change pointer value 
	*b = 10 // so value of var has been change as well
	fmt.Println(*b)
	fmt.Println(a)
}