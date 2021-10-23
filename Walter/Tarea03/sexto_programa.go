package main
/* 6.-Escriba la cadena "El verano se da en los messes de, , y parte de ",
preguntando desde el teclado los meses.
*/
import "fmt"

func main(){
	var mes1,mes2,mes3 string

 	fmt.Println("Ingresa los meses de verano:")
	fmt.Print("->")
	fmt.Scanln(&mes1)
	fmt.Print("->")
	fmt.Scanln(&mes2)
	fmt.Print("->")
	fmt.Scanln(&mes3)

	fmt.Printf("El verano se da en los meses de, %s, %s y parte de %s.",mes1,mes2,mes3)
}
