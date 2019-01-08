package main

import "fmt"

func main() {

	/*fmt.Printf("***** mi programa con go *******\n")

	var edad int = 18

	if edad > 18 {
		fmt.Printf("Eres mayor de edad")
	} else {
		fmt.Printf("Eres menor de edad")
	}*/

	/*for i := 1; i < 15; i++ {
		fmt.Println(i, "es el número actual")
	}*/

	var numero1 float32 = 10
	var numero2 float32 = 6

	//Suma
	fmt.Print("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))

	//Resta
	fmt.Print("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))

	//Multiplicacion
	fmt.Print("La Multiplicacion es: ")
	fmt.Println(operacion(numero1, numero2, "*"))

	//La división
	fmt.Print("La división es: ")
	fmt.Println(operacion(numero1, numero2, "/"))

}

// Hacer una funcion con las operaciones que hay arriba

func operacion(n1 float32, n2 float32, operador string) float32 {
	var resultado float32

	if operador == "+" {
		resultado = n1 + n2
	}
	if operador == "-" {
		resultado = n1 - n2
	}
	if operador == "*" {
		resultado = n1 * n2
	}
	if operador == "/" {
		resultado = n1 / n2
	}

	return resultado
}
