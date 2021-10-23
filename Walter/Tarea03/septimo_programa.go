package main
/*
7.-Escriba la cadena "Las estaciones son *Verano, e " preguntando desde el
teclado las estaciones del año.
*/
import "fmt"

func main(){
	var (
		estacion1,estacion2,estacion3,estacion4 string
	)
	fmt.Println("Ingresa las 4 estaciones del año:")
	fmt.Print("->")
	fmt.Scanln(&estacion1)
	fmt.Print("->")
	fmt.Scanln(&estacion2)
	fmt.Print("->")
	fmt.Scanln(&estacion3)
	fmt.Print("->")
	fmt.Scanln(&estacion4)

	fmt.Printf("Las estaciones del año son %s, %s, %s, e %s.",estacion1,estacion2,estacion3,estacion4)
}
