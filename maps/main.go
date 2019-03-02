package main 
import "fmt"

func  main()  {
	emails := make(map[string] string)
	name := map[string]string {"1":"vin","2":"mike","3":"joe"}

	emails["vin"] = "nagacoder@outlook.com"
	emails["mike"] = "mike@gmail.com"
	emails["joe"] = "joe@gmail.com"

	fmt.Println(emails)
	delete(emails,"joe")
	fmt.Println(emails["vin"])
	fmt.Println(name)
}