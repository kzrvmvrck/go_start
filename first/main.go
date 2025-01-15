package main // Определяет название текущего пакета
// должен быть в первой строчке
// если main - это точка входа в выполнение программы, если его не будет программа не запуститься
//

import (
	"fmt" // библиотека содержащая необходимы нам инструменты
	"reflect"
)

func main() { // основная функция приложения
	// var message string - базовый способ инициализации переменной
	// const message string = "Я скоро стану писать на GO."
	message := "Я скоро стану писать на GO." // двоеточие это инициализация переменной и присваивание ей значения
	var zeroString string
	var zeroInt int       //int могут быть разые в зависимости от битности!
	var zeroFloat float32 // тоже завист
	var zeroBool bool

	xMassege := []byte("keizarov")
	var a rune = 'a' // ОДИНАРНЫЕ СКОБКИ!

	zeroFloat = 12.2
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))

	fmt.Println(zeroString)
	fmt.Println(zeroInt)
	fmt.Println(zeroFloat)
	fmt.Println(zeroBool)
	fmt.Println(xMassege)
	fmt.Println(a)

	// Мнодественное присвоение значения
	s, d, g := 1, 2, 3
	s, d = d, s
	fmt.Println(s, d, g)

	messageFromFunc := sayHello("Maksim", 34)
	printMessage("Вызов 1")
	printMessage("Вызов 2")
	printMessage("Вызов 3")
	printMessage(messageFromFunc)

	mes, entered := enterTheClub(19)
	fmt.Println(mes)
	fmt.Println(entered)

}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		response := "Welcome!"
		return response, true
	} else {
		response := "go out"
		return response, false
	}
}
