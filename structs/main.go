package main 
import (
	"fmt"
	"strconv")
// define a struct
type Person struct {
	name string
	username string
	address string
	age int
}
// define method (value receiver)
func (p Person) greet() string{
	return "Hello, i am " +p.name+" "+"i am"+" "+strconv.Itoa(p.age) +" and also living in " +p.address
}
// define method( pointer receiver)
func (p *Person) changeName(newName string) string{
	 p.name = newName
	 return "."
}
func  main()  {
	darvin := Person{"Darvin","nagacoder","Batam",23}
	fmt.Println(darvin.greet())
	fmt.Println(darvin.changeName("Sinaga"))
	fmt.Println(darvin.greet())
}