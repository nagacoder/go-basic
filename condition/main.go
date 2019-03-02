package main 
import "fmt"

func  main()  {
	// if
	x := 10
	y := 20
	if x < y {
		fmt.Printf("%d Lebih Kecil dari %d \n",x,y)
	}
	// if else
	lang := "rb"

	if lang == "js"{
		fmt.Println("This Javascript format")
	}else if lang == "py"{
		fmt.Println("This Python format")
	}else {
		fmt.Println("This not python and Javascript")
	}
	// switch
	switch lang {
	case "js":
		fmt.Println("This Javascript format")
	case "py":
		fmt.Println("This Python format")
	default :
	fmt.Println("This not python and Javascript")
	}
}