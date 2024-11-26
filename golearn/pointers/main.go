package main

import "fmt"

// & = Амперсанд = определяет адрес переменной в памяти
// * = Астериск

func main() {
	Ampersand()
}

func Ampersand() {
	answer := 42
	fmt.Println(&answer)
	address := &answer
	fmt.Println(*address)

	fmt.Printf("address это %T\n", address)

	fmt.Printf("Значение на которое указывает address: %d\n", *address)

	*address = 100
	fmt.Printf("Значение на которое указывает address: %d\n", *address)

	canada := "Canada"
	var home *string
	fmt.Printf("home is a %T\n", home)
	// Присваиваем указателю home адрес переменной canada
	home = &canada
	// Разыменовываем указатель home и выводим значение
	fmt.Println(*home)
}
