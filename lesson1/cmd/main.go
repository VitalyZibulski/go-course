package main

import "fmt"

//задача 1
//обьявление двух переменных и вывод их суммы
//func main() {
//	a := 1
//	b := 2
//	fmt.Println(a + b)
//}

// задача 2
// обьявление 3 переменных, вывод произведения любых двух, за вычетом последнего
//func main() {
//	a := 2
//	b := 2
//	c := 1
//
//	fmt.Println(a * b - c)
//}

//задача 3
//ввод имени фамилии и отчества с возврастом, вывод на 1 строке

//func main() {
//	var name string
//	var surName string
//
//	fmt.Print("Enter your name: ")
//	fmt.Scanln(&name)
//
//	fmt.Print("Enter your surName: ")
//	fmt.Scanln(&surName)
//
//	fmt.Println(name + " " + surName)
//}

//
//задача 4
//(со звездочкой, не обязательна для выполнения)
//сканирование 2 чисел из терминала, а затем вывод их произведения

func main() {
	var number1 int
	var number2 int

	fmt.Print("Enter number1: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter number2: ")
	fmt.Scanln(&number2)

	fmt.Println(number1 * number2)
}
