package main
/*
9.-Escriba la cadena "Trabajamos de a", preguntando desde el teclado los días de la semana.
*/
import "fmt"
func main() {
	var (
		diaTrabajo, diaTrabajo2 string
	)
	fmt.Println("Ingresa de que a que dia trabajas:")
	fmt.Print("->")
	fmt.Scanln(&diaTrabajo)
	fmt.Print("->")
	fmt.Scanln(&diaTrabajo2)

	fmt.Printf("Trabajamos de %s a %s.", diaTrabajo, diaTrabajo2)
}
