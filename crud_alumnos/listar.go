//CARLOS ARMANDO ITURRIOS ALCARAZ ########## ITIC 10-1

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func listar() {
	dataDos, err := ioutil.ReadFile("alumnos.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(dataDos))
}

func listarUnoPorUno() {
	dataDos, err := ioutil.ReadFile("alumnos_UnoPorUno.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(dataDos))
}
