package main 
import "fmt"

func  main()  {
	num := []int{12,34,34,56,56,454,56}
	emails := map[string]string {"vin":"nagacoderGgmail.com","mike":"mike@gmail.com","joe":"joe@outlook.com"}

	for _, n := range num{
		fmt.Printf("id:%d\n",n)
	}
	for _,v := range emails {
		fmt.Printf("email: %s\n",v)
	}
	for k := range emails {
		fmt.Printf("name: %s\n",k)
	}
}