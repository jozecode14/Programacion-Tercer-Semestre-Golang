package main

import "fmt" //FUNCION QUE SUMA DOS NUMEROS CON ARGUMENTOS INTERFACE
func sumarNumeros(value1, value2 interface{})(int,bool){ //funcion que tiene como argumento dos variables de tipo interface
	suma,error := 0,true
	switch value1.(type) { //Comparar el tipo de dato
	case int:
		switch value2.(type) {
		case int:
			suma = value1.(int) + value2.(int) //Convertimos y sumamos
		}
	default:
		error = false
	}
	return suma,error
}
	//FUNCION QUE CONCATENA DOS STRINGS CON ARGUMENTOS INTERFACE
func concatenarStrings(string1, string2 interface{})(string,bool){ //funcion que tiene como argumento dos variables de tipo interface
	concatenar,error := "",true
	switch string1.(type) { //Comparar el tipo de dato .(type)
	case string:
		switch string2.(type) {
		case string:
			concatenar = string1.(string)+" "+string2.(string) //convertimos y concatenamos
		}
	default:
		error = false
	}
	return concatenar,error
}

func main(){
	var (
		num1 = 5.5
		num2 = 5
	)
	suma,err := sumarNumeros(num1,num2)
	if err{
		fmt.Printf("%d + %d = %d",num1,num2,suma)
	} else{
		fmt.Println("Error, los numeros deberían ser enteros")
	}

	fmt.Println("")

	var(
		string1 = 5
		string2 = "mundo"
	)
	texto, err2:=concatenarStrings(string1,string2)

	if err2{
		fmt.Printf("%s",texto)
	}else{
		fmt.Println("Error la entrada debería ser string")
	}
}
