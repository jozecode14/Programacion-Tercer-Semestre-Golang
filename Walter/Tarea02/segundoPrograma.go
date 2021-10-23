package main
/*
									----2DO PROGRAMA----
Escriba una función que reciba Nombre (s), Apellido Paterno, Apellido Materno y Año de Nacimiento.
Que escriba dentro de la función el nombre completo iniciando por los apelllidos y escriba la edad
de la persona.
*/
import (
	"bufio"
	"fmt"
	"os"
)
func main()  {
	datosPersonales()
}

func datosPersonales() {
	var (
		apellidoMa, apellidoPa string
		aNac                   int16
		aActual                int16 = 2021
	)
	fmt.Print("Ingresa tu nombre completo!\nApellido paterno:")
	fmt.Scanln(&apellidoPa)

	fmt.Print("Apellido materno:")
	fmt.Scanln(&apellidoMa)

	fmt.Print("Ahora ingresa tu nombre(s):")
	scanner := bufio.NewScanner(os.Stdin) //Añadimos la libreria bufio, con el metodo NewScanner(0s.Stdion) para escanear
	scanner.Scan()                        //Scaneamos el string
	nombre := scanner.Text()              //Guardamos en la variable y lo convertimos a texto

	fmt.Print("Ahora ingresa tu año de nacimiento:")
	fmt.Scanln(&aNac)

	fmt.Printf("Tu nombre es %s %s %s y tienes %d años de edad.", apellidoPa, apellidoMa, nombre, aActual-aNac)
}
