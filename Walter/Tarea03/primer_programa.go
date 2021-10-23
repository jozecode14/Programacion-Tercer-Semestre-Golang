package main

import (
	"fmt"
)

func main() {
	//1.-Generar las tablas de verdad de AND y OR
	var (
		verdad bool = true //Formato para bool %t
		falso bool = false
	)

	fmt.Println("Tabla de verdad de AND ")
	fmt.Printf("%t  && %t = %t\n",verdad,verdad,verdad && verdad)
	fmt.Printf("%t  && %t = %t\n",verdad,falso,verdad && falso)
	fmt.Printf("%t && %t = %t\n",falso,verdad,falso && verdad)
	fmt.Printf("%t && %t = %t\n",falso,falso,falso && falso)

	fmt.Println("")

	fmt.Println("Tabla de verdad de OR ")
	fmt.Printf("%t  || %t = %t\n",verdad,verdad,verdad || verdad)
	fmt.Printf("%t  || %t = %t\n",verdad,falso,verdad || falso)
	fmt.Printf("%t || %t = %t\n",falso,verdad,falso || verdad)
	fmt.Printf("%t || %t = %t\n",falso,falso,falso || falso)
}
