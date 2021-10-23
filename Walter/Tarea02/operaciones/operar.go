package operaciones //Lleva el nombre de la carpeta donde se ubica el paquete

import "fmt"

func Operando(){
	var(
		num1, num2 int = 10, 5
	)
	fmt.Println("--RESULTADO--")

	fmt.Printf("%d + %d = %d\n",num1, num2, suma(num1,num2))
	fmt.Printf("%d - %d = %d\n",num1, num2, resta(num1,num2))
	fmt.Printf("%d x %d = %d\n",num1, num2, multiplicacion(num1,num2))
	fmt.Printf("%d / %d = %d\n",num1, num2, division(num1,num2))
}
func suma(n1, n2 int)int{
	return n1 + n2
}
func resta(n1, n2 int)int{
	return n1 - n2
}
func multiplicacion(n1, n2 int)int{
	return n1 * n2
}
func division(n1, n2 int)int{
	return n1 / n2
}