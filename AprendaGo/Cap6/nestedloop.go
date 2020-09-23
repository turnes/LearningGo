package main

import (
	"fmt"
)

var lastdayofmonth int

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora: ", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	}

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês: ", mes)
		if mes == 2 {
			lastdayofmonth = 28
		} else if mes%2 == 0 {
			lastdayofmonth = 30
		} else {
			lastdayofmonth = 31
		}
		for dia := 1; dia <= lastdayofmonth; dia++ {
			fmt.Printf("Dia: %v, ", dia)
		}
		fmt.Println("\n")
	}
}
