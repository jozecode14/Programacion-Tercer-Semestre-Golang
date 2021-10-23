package main
//8.-Escriba la cadena "Descansamos y" preguntando desde el teclado los dias de la semana.
import "fmt"

func main()  {
	var (
		descanso,descanso2 string
	)
	fmt.Println("Ingresa los dias de descanso:")
	fmt.Print("->")
	fmt.Scanln(&descanso)
	fmt.Print("->")
	fmt.Scanln(&descanso2)

	fmt.Printf("Descanzamos %s y %s.",descanso,descanso2)
}