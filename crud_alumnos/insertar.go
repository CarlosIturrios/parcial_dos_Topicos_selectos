package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func obtenerDatosInsertar() Alumno {
	alumno := Alumno{
		Id:           0,
		Calificacion: "",
		Nombre:       "",
		Apellidos:    "",
	}

	fmt.Println("###### LLENE LOS DATOS DE MANERA CORRECTA ######")
	fmt.Println("Escriba los apellidos del alumno: ")
	scannerApellidos := bufio.NewScanner(os.Stdin)
	scannerApellidos.Scan()
	apellidos := scannerApellidos.Text()
	fmt.Println("Escriba el nombre o los nombres del alumno: ")
	scannerNombre := bufio.NewScanner(os.Stdin)
	scannerNombre.Scan()
	nombre := scannerNombre.Text()

	fmt.Println("Escriba la calificacion del alumno: ")
	scannerCalificacion := bufio.NewScanner(os.Stdin)
	scannerCalificacion.Scan()
	calificacion := scannerCalificacion.Text()
	alumno.Calificacion = calificacion
	alumno.Nombre = nombre
	alumno.Apellidos = apellidos
	return alumno
}

func EscribirDatos() {
	filename := "alumnos.txt"
	var err = os.Remove(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	res := []Alumno{}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	for i := 0; i < 5; i++ {
		alumno := obtenerDatosInsertar()
		alumno.Id = i + 1
		res = append(res, alumno)

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	tamanoArray := len(res)
	for i := 0; i < tamanoArray; i++ {
		fmt.Fprintf(f, "\n %d\t%s\t%s\t%s\t", res[i].Id, res[i].Apellidos, res[i].Nombre, res[i].Calificacion)
	}
	defer f.Close()
}

func EscribirDatosUnoPorUno() {
	filename := "alumnos_UnoPorUno.txt"
	res := []Alumno{}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	alumno := obtenerDatosInsertar()
	alumno.Id = rand.Intn(100)
	res = append(res, alumno)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	tamanoArray := len(res)
	for i := 0; i < tamanoArray; i++ {
		fmt.Fprintf(f, "\n %d\t%s\t%s\t%s\t", res[i].Id, res[i].Apellidos, res[i].Nombre, res[i].Calificacion)
	}
	defer f.Close()
}
