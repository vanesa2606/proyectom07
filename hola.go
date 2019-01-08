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

	fmt.Print("Calculadora 1")
	calculadora(numero1, numero2)

	var numero3 float32 = 48
	var numero4 float32 = 28

	fmt.Print("------------------------------\n")
	fmt.Print("Calculadora 2\n")
	calculadora(numero3, numero4)

	/*	//Suma
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

	*/

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

func calculadora(numero1 float32, numero2 float32) {
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
