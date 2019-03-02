package main 
import "fmt"
func tambah() func(int) int{
	angka :=0
	return func(num int) int{
		angka += num
		return angka
	}
}
func  main()  {
	tambah := tambah()
	for i :=0; i <= 10; i++{
		fmt.Println(tambah(i))
	} 
}