package main

import (
	"flag"
	"fmt"
)

func main() {

	//Define flags
	op := flag.String("op", "", "Opration to perfrom: add or sub")
	n1 := flag.Float64("n1", 0, "Frist Operation")
	n2 := flag.Float64("n2", 0, "Second Operation")

	//Parse the flags
	flag.Parse()

	//Validate and Execute Operation:
	switch *op {
	case "add":
		fmt.Printf("Result: %.2f\n", *n1+*n2)
	case "sub":
		fmt.Printf("Result: %.2f\n", *n1-*n2)
	default:
		fmt.Println("Erro: Please enter the valid number and operation")
	}
}
