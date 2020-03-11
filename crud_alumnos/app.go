//CARLOS ARMANDO ITURRIOS ALCARAZ ########## ITIC 10-1

package main

import "fmt"

func app() {
	bandera := 0
	for bandera < 1 {
		var value int
		fmt.Println("\nElige una de las opciones:\nCapturar Alumnos Masivo - Presiona 1 \nCapturar Alumnos uno por uno - Presiona 2 \nObtener Alumnos Masivo - Presiona 3\nObtener Alumnos uno por uno - Presiona 4")
		fmt.Println("Salir - presiona 6")
		_, err := fmt.Scanf("%d", &value)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err.Error())
			value = 0
		}

		switch {
		case value == 1:
			EscribirDatos()
		case value == 2:
			EscribirDatosUnoPorUno()
		case value == 3:
			listar()
		case value == 4:
			listarUnoPorUno()
		case value == 6:
			bandera = 1
		default:
			fmt.Println("Opcion invalida vuelva a intentarlo")
		}
	}
}
