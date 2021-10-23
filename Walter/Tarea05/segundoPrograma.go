package main

import (
	"fmt"
	"math/rand"
	"time"
)
const totalEstudiantes = 600
func main() {
	var (
		materia1 []int
		materia2 []int
		materia3 []int
		materia4 []int
		materia5 []int
	)
	//MATERIA 1 PROGRAMACIÓN
	var (
		sumatoriaReprobado1 int
		sumatoriaAprobado1  int
		promedio1           int
		reprobados1         int
		aprobados1          int
		promedioAprobado1   int
		promedioReprobado1  int
		mejoresAlumnos1     int
		peoresAlumnos1      int
		mejoresAlumnosCant1 int
		peoresAlumnosCant1  int
	)

	for i := 0; i < totalEstudiantes; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNumbers := rand.Intn(11)
		/*fmt.Print("[",i,"]",randomNumbers," ")*/
		materia1 := append(materia1, randomNumbers) //randomNumbers se queda igual
		for _, materia1 := range materia1 {
			promedio1 += materia1

			if materia1 < 6 {
				sumatoriaReprobado1 += materia1
				reprobados1++
			} else {
				aprobados1++
				sumatoriaAprobado1 += materia1
			}

			if materia1 == 10 {
				mejoresAlumnos1 = i
				/*fmt.Println(mejoresAlumnos)*/
				mejoresAlumnosCant1++ //En general sumar con todas
			}

			if materia1 == 0 {
				peoresAlumnos1 = i
				peoresAlumnosCant1++ //En general sumar con todas
			} else if materia1 < 2 {
				peoresAlumnos1 = i
				peoresAlumnosCant1++
			}

		}
	}
	sumatoriaReprobado1F := float64(sumatoriaReprobado1)
	sumatoriaAprobado1F := float64(sumatoriaAprobado1)
	aprobados1F := float64(aprobados1)
	reprobados1F := float64(reprobados1)
	promedioF1 := float64(promedio1)
	totalEstudiantesF1 := float64(totalEstudiantes) //Tampoco cambia
	promedioReprobadoF1 := float64(promedioReprobado1)
	promedioAprobadoF1 := float64(promedioAprobado1)
	promedioReprobadoF1 = sumatoriaReprobado1F / reprobados1F
	promedioAprobadoF1 = sumatoriaAprobado1F / aprobados1F
	promedioFinal1 := promedioF1/totalEstudiantesF1


	fmt.Println("\t----PROGRAMACIÓN---")
	fmt.Printf("1) El promedio de la materia Programación es de: %.1f \n", promedioFinal1)
	fmt.Println("2) Hay un total de", reprobados1, "reprobados y ", aprobados1, " aprobados en esta materia.")
	fmt.Println("3) El promedio de los alumnos reprobados es de:", promedioReprobadoF1, " y de los aprobados es de:", promedioAprobadoF1)
	fmt.Println("4) El alumno con el promedio mas alto fue el alumno numero:", mejoresAlumnos1)
	fmt.Println("5) El alumno con el promedio mas bajo fue el alumno numero:", peoresAlumnos1)

	fmt.Println("")

	//MATERIA 2 REDES
	var (
		sumatoriaReprobado2 int
		sumatoriaAprobado2  int
		promedio2           int
		reprobados2         int
		aprobados2          int
		promedioAprobado2   int
		promedioReprobado2  int
		mejoresAlumnos2     int
		peoresAlumnos2      int
		mejoresAlumnosCant2 int
		peoresAlumnosCant2  int
	)

	for i := 0; i < totalEstudiantes; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNumbers := rand.Intn(11)
		/*fmt.Print("[",i,"]",randomNumbers," ")*/
		materia2 := append(materia2, randomNumbers) //randomNumbers se queda igual
		for _, materia2 := range materia2 {
			promedio2 += materia2

			if materia2 < 6 {
				sumatoriaReprobado2 += materia2
				reprobados2++
			} else {
				aprobados2++
				sumatoriaAprobado2 += materia2
			}

			if materia2 > 9 {
				mejoresAlumnos2 = i
				/*fmt.Println(mejoresAlumnos)*/
				mejoresAlumnosCant2++ //En general sumar con todas
			} else if materia2 > 8 {
				mejoresAlumnos2 = i
				mejoresAlumnosCant2++
			}

			if materia2 < 1 {
				peoresAlumnos1 = i
				peoresAlumnosCant1++ //En general sumar con todas
			} else if materia2 < 2 {
				peoresAlumnos2 = i
				peoresAlumnosCant2++
			}

		}
	}
	sumatoriaReprobado2F := float64(sumatoriaReprobado2)
	sumatoriaAprobado2F := float64(sumatoriaAprobado2)
	aprobados2F := float64(aprobados2)
	reprobados2F := float64(reprobados2)
	promedioF2 := float64(promedio2)
	totalEstudiantesF2 := float64(totalEstudiantes) //Tampoco cambia
	promedioReprobadoF2 := float64(promedioReprobado2)
	promedioAprobadoF2 := float64(promedioAprobado2)
	promedioReprobadoF2 = sumatoriaReprobado2F / reprobados2F
	promedioAprobadoF2 = sumatoriaAprobado2F / aprobados2F
	promedioFinal2 := promedioF2/totalEstudiantesF2


	fmt.Println("\t----REDES---")
	fmt.Printf("1) El promedio de la materia Redes es de: %.1f \n",promedioFinal2)
	fmt.Println("2) Hay un total de", reprobados2, "reprobados y ", aprobados2, " aprobados en esta materia.")
	fmt.Println("3) El promedio de los alumnos reprobados es de:", promedioReprobadoF2, " y de los aprobados es de:", promedioAprobadoF2)
	fmt.Println("4) El alumno con el promedio mas alto fue el alumno numero:", mejoresAlumnos2)
	fmt.Println("5) El alumno con el promedio mas bajo fue el alumno numero:", peoresAlumnos2)

	fmt.Println("")

	//MATERIA 3 CALCULO
	var (
		sumatoriaReprobado3 int
		sumatoriaAprobado3  int
		promedio3           int
		reprobados3         int
		aprobados3          int
		promedioAprobado3   int
		promedioReprobado3  int
		mejoresAlumnos3     int
		peoresAlumnos3      int
		mejoresAlumnosCant3 int
		peoresAlumnosCant3  int
	)

	for i := 0; i < totalEstudiantes; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNumbers := rand.Intn(11)
		/*fmt.Print("[",i,"]",randomNumbers," ")*/
		materia3 := append(materia3, randomNumbers) //randomNumbers se queda igual
		for _, materia3 := range materia3 {
			promedio3 += materia3

			if materia3 < 6 {
				sumatoriaReprobado3 += materia3
				reprobados3++
			} else {
				aprobados3++
				sumatoriaAprobado3 += materia3
			}

			if materia3 > 9 {
				mejoresAlumnos3 = i
				/*fmt.Println(mejoresAlumnos)*/
				mejoresAlumnosCant3++ //En general sumar con todas
			} else if materia3 > 8 {
				mejoresAlumnos3 = i
				mejoresAlumnosCant3++
			}

			if materia3 < 1 {
				peoresAlumnos3 = i
				peoresAlumnosCant3++ //En general sumar con todas
			} else if materia3 < 2 {
				peoresAlumnos3 = i
				peoresAlumnosCant3++
			}

		}
	}
	sumatoriaReprobado3F := float64(sumatoriaReprobado3)
	sumatoriaAprobado3F := float64(sumatoriaAprobado3)
	aprobados3F := float64(aprobados3)
	reprobados3F := float64(reprobados3)
	promedioF3 := float64(promedio3)
	totalEstudiantesF3 := float64(totalEstudiantes) //Tampoco cambia
	promedioReprobadoF3 := float64(promedioReprobado3)
	promedioAprobadoF3 := float64(promedioAprobado3)
	promedioReprobadoF3 = sumatoriaReprobado3F / reprobados3F
	promedioAprobadoF3 = sumatoriaAprobado3F / aprobados3F
	promedioFinal3 := promedioF3/totalEstudiantesF3


	fmt.Println("\t----CALCULO---")
	fmt.Printf("1) El promedio de la materia Calculo es de: %.1f \n",promedioFinal3)
	fmt.Println("2) Hay un total de", reprobados3, "reprobados y ", aprobados3, " aprobados en esta materia.")
	fmt.Println("3) El promedio de los alumnos reprobados es de:", promedioReprobadoF3, " y de los aprobados es de:", promedioAprobadoF3)
	fmt.Println("4) El alumno con el promedio mas alto fue el alumno numero:", mejoresAlumnos3)
	fmt.Println("5) El alumno con el promedio mas bajo fue el alumno numero:", peoresAlumnos3)

	fmt.Println("")

	//MATERIA 4 PROBABILIDAD
	var (
		sumatoriaReprobado4 int
		sumatoriaAprobado4  int
		promedio4           int
		reprobados4         int
		aprobados4          int
		promedioAprobado4   int
		promedioReprobado4  int
		mejoresAlumnos4     int
		peoresAlumnos4      int
		mejoresAlumnosCant4 int
		peoresAlumnosCant4  int
	)

	for i := 0; i < totalEstudiantes; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNumbers := rand.Intn(11)
		/*fmt.Print("[",i,"]",randomNumbers," ")*/
		materia4 := append(materia4, randomNumbers) //randomNumbers se queda igual
		for _, materia4 := range materia4 {
			promedio4 += materia4

			if materia4 < 6 {
				sumatoriaReprobado4 += materia4
				reprobados4++
			} else {
				aprobados4++
				sumatoriaAprobado4 += materia4
			}

			if materia4 > 9 {
				mejoresAlumnos4 = i
				/*fmt.Println(mejoresAlumnos)*/
				mejoresAlumnosCant4++ //En general sumar con todas
			} else if materia4 > 8 {
				mejoresAlumnos4 = i
				mejoresAlumnosCant4++
			}

			if materia4 < 1 {
				peoresAlumnos4 = i
				peoresAlumnosCant4++ //En general sumar con todas
			} else if materia4 < 2 {
				peoresAlumnos4 = i
				peoresAlumnosCant4++
			}

		}
	}
	sumatoriaReprobado4F := float64(sumatoriaReprobado4)
	sumatoriaAprobado4F := float64(sumatoriaAprobado4)
	aprobados4F := float64(aprobados4)
	reprobados4F := float64(reprobados4)
	promedioF4 := float64(promedio4)
	totalEstudiantesF4 := float64(totalEstudiantes) //Tampoco cambia
	promedioReprobadoF4 := float64(promedioReprobado4)
	promedioAprobadoF4 := float64(promedioAprobado4)
	promedioReprobadoF4 = sumatoriaReprobado4F / reprobados4F
	promedioAprobadoF4 = sumatoriaAprobado4F / aprobados4F
	promedioFinal4 := promedioF4/totalEstudiantesF4


	fmt.Println("\t----PROBABILIDAD---")
	fmt.Printf("1) El promedio de la materia Probabilidad es de: %.1f \n",promedioFinal4)
	fmt.Println("2) Hay un total de", reprobados4, "reprobados y ", aprobados4, " aprobados en esta materia.")
	fmt.Println("3) El promedio de los alumnos reprobados es de:", promedioReprobadoF4, " y de los aprobados es de:", promedioAprobadoF4)
	fmt.Println("4) El alumno con el promedio mas alto fue el alumno numero:", mejoresAlumnos4)
	fmt.Println("5) El alumno con el promedio mas bajo fue el alumno numero:", peoresAlumnos4)

	fmt.Println("")

	//MATERIA 5 ESTRUCTURA DE DATOS
	var (
		sumatoriaReprobado5 int
		sumatoriaAprobado5  int
		promedio5           int
		reprobados5         int
		aprobados5          int
		promedioAprobado5   int
		promedioReprobado5  int
		mejoresAlumnos5     int
		peoresAlumnos5      int
		mejoresAlumnosCant5 int
		peoresAlumnosCant5  int
	)

	for i := 0; i < totalEstudiantes; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNumbers := rand.Intn(11)
		/*fmt.Print("[",i,"]",randomNumbers," ")*/
		materia5 := append(materia5, randomNumbers) //randomNumbers se queda igual
		for _, materia5 := range materia5 {
			promedio5 += materia5

			if materia5 < 6 {
				sumatoriaReprobado5 += materia5
				reprobados5++
			} else {
				aprobados5++
				sumatoriaAprobado5 += materia5
			}

			if materia5 > 9 {
				mejoresAlumnos5 = i
				/*fmt.Println(mejoresAlumnos)*/
				mejoresAlumnosCant5++ //En general sumar con todas
			} else if materia5 > 8 {
				mejoresAlumnos5 = i
				mejoresAlumnosCant5++
			}

			if materia5 < 1 {
				peoresAlumnos5 = i
				peoresAlumnosCant5++ //En general sumar con todas
			} else if materia5 < 2 {
				peoresAlumnos5 = i
				peoresAlumnosCant5++
			}

		}
	}
	sumatoriaReprobado5F := float64(sumatoriaReprobado5)
	sumatoriaAprobado5F := float64(sumatoriaAprobado5)
	aprobados5F := float64(aprobados5)
	reprobados5F := float64(reprobados5)
	promedioF5 := float64(promedio5)
	totalEstudiantesF5 := float64(totalEstudiantes) //Tampoco cambia
	promedioReprobadoF5 := float64(promedioReprobado5)
	promedioAprobadoF5 := float64(promedioAprobado5)
	promedioReprobadoF5 = sumatoriaReprobado5F / reprobados5F
	promedioAprobadoF5 = sumatoriaAprobado5F / aprobados5F
	promedioFinal5 := promedioF5/totalEstudiantesF5

	fmt.Println("\t----ESTRUCTURA DE DATOS---")
	fmt.Printf("1) El promedio de la materia Estructura de datos es de: %.1f \n", promedioFinal5)
	fmt.Println("2) Hay un total de", reprobados5, "reprobados y ", aprobados5, " aprobados en esta materia.")
	fmt.Println("3) El promedio de los alumnos reprobados es de:", promedioReprobadoF5, " y de los aprobados es de:", promedioAprobadoF5)
	fmt.Println("4) El alumno con el promedio mas alto fue el alumno numero:", mejoresAlumnos5)
	fmt.Println("5)El alumno con el promedio mas bajo fue el alumno numero:", peoresAlumnos5)

	fmt.Println("")

	fmt.Println("\t----ASPECTOS GENERALES----")

	fmt.Println("6) Cantidad de alumnos con el peor promedio de todas las materias ",mejoresAlumnosCant1+mejoresAlumnosCant2+mejoresAlumnosCant3+mejoresAlumnosCant4+mejoresAlumnosCant5,"" +
		"\nCantidad de alumnos con el mejor promedio de todas las materias:",peoresAlumnosCant1+peoresAlumnosCant2+peoresAlumnosCant3+peoresAlumnosCant4+peoresAlumnosCant5)

	fmt.Println("")

	calificacionMayor := [] float64 {promedioFinal1,promedioFinal2,promedioFinal3,promedioFinal4,promedioFinal5}
	numeroMayor := calificacionMayor[0] //Asignar variable al primer elemento
	for _, numero := range calificacionMayor {
		if numero > numeroMayor{
			numeroMayor = numero
		}
	}
	fmt.Printf("7) La calificación mayor fue de %f",numeroMayor)

	if promedioFinal1 == numeroMayor{
		fmt.Print(" La materia programación.")
	}

	if promedioFinal2 == numeroMayor{
		fmt.Print(" La materia de redes.")
	}

	if promedioFinal3 == numeroMayor{
		fmt.Print(" La materia de calculo.")
	}

	if promedioFinal4 == numeroMayor{
		fmt.Print(" La materia de probabilidad.")
	}

	if promedioFinal5 == numeroMayor{
		fmt.Print(" La materia de estructura de datos.")
	}
}