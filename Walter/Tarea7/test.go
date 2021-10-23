package main

import (
	"fmt"
)
 //TESTING DEL ERROR
func main()  {
	n := 5
	suma,error := factorial(n)
	if error{
		fmt.Println("El factorial de !",n,"es:",suma)
	}else{
	fmt.Println("El número debe ser entero y positivo")
	}
}

func factorial(n int) (int,bool)  {
	suma,error := 0, true
	if n > 0 {
		for i := 0; i < n; i-- {
			suma *= n
		}
	}else {
		error = false
	}
	return suma,error
}
