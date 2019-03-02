package main 
import "fmt"

func  main()  {
	// using var
	// var name = " Darvin Sinaga"
	var age = 23 
	var isCool = true
	var long float32 = 2.2 //var with data type float32 we can set types if we want as well
	isCool = false
	// shorthand declare var


	// if inside a function
	address := "batam"
	size := 2.3 // the data type is default float64, if you want spesific types
	//you can see var long below var is isCooll

	// we can also define var with shorthand like this
	username,name, email := "nagacoder","Darvin Sinaga", "me@darvinsinaga.com"
	// log
	fmt.Println(name,age,username,email, isCool,address,size)
	fmt.Printf("%T\n", long)
}