package operaciones

import (
	"fmt"
	"strconv"
)

func EvaluarTemperaturas(data[][]string){
	println()
	diaMasCaluroso,_:= strconv.ParseFloat(data[0][0],64)
	diaMasFrio,_:= strconv.ParseFloat(data[0][0],64)

	var diaCalor, diaFrio, ubicacionCalor, ubicacionFrio int

	for i := 0; i<len(data[0]);i++{
		for j:= 0; j < len(data); j++{
			tempDay,_:=strconv.ParseFloat(data[j][i],64)
			if tempDay >= diaMasCaluroso{
				diaMasCaluroso,diaCalor,ubicacionCalor = tempDay, i+1, i+1
			}
			if tempDay <= diaMasFrio{
				diaMasFrio, diaFrio, ubicacionFrio = tempDay, i+1,i+1
			}
		}
	}

	fmt.Printf("2.2¿CUAL FUE EL DIA QUE SE REGISTRO LA TEMPERATURA MAS ALTA Y BAJA?\n")
	fmt.Printf("\t\tCALOR:\n")
	fmt.Printf("\t -TEMPERATURA MAS ALTA: %.2f \n",diaMasCaluroso)
	fmt.Printf("\t -DIA: %d\n",diaCalor)
	fmt.Printf("\t -MUESTRA: %d\n",ubicacionCalor)

	fmt.Printf("\n\t\tFRIO:\n")
	fmt.Printf("\t -TEMPERATURA MAS ALTA: %.2f \n",diaMasFrio)
	fmt.Printf("\t -DIA: %d\n",diaFrio)
	fmt.Printf("\t -MUESTRA: %d\n",ubicacionFrio)
}
