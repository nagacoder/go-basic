package main 
import "fmt"

func  main()  {
	// array 
	var lang [2]string
	lang[0] = "Javascript"
	lang[1] = "Python"
	// also can like this
	arr := [3] string {"Darvin","Sinaga", "Developer"}
	// slice 
	var city = [] string {"Medan","Batam","Jakarta"}
	fmt.Println(lang)
	fmt.Println(lang[0],lang[1])
	fmt.Println(arr)
	fmt.Println(arr[2])
	fmt.Println(city)
	fmt.Println(len(city))
	fmt.Println(city[2:3])
}