package main

import "fmt"

func main() {

	slice := make([]string, 26, 30)

	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Printf("Lenght: %v\tCapacity: %v\n", len(slice), cap(slice))

	for x := 0; x < len(slice); x++ {
		fmt.Println(slice[x])
	}

}
