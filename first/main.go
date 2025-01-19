package main // Определяет название текущего пакета
// должен быть в первой строчке
// если main - это точка входа в выполнение программы, если его не будет программа не запуститься
//

import (
	"errors"
	"fmt" // библиотека содержащая необходимы нам инструменты
)

func main() { // основная функция приложения
	// // var message string - базовый способ инициализации переменной
	// // const message string = "Я скоро стану писать на GO."
	// message := "Я скоро стану писать на GO." // двоеточие это инициализация переменной и присваивание ей значения
	// var zeroString string
	// var zeroInt int       //int могут быть разые в зависимости от битности!
	// var zeroFloat float32 // тоже завист
	// var zeroBool bool

	// xMassege := []byte("keizarov")
	// var a rune = 'a' // ОДИНАРНЫЕ СКОБКИ!

	// zeroFloat = 12.2
	// fmt.Println(message)
	// fmt.Println(reflect.TypeOf(message))

	// fmt.Println(zeroString)
	// fmt.Println(zeroInt)
	// fmt.Println(zeroFloat)
	// fmt.Println(zeroBool)
	// fmt.Println(xMassege)
	// fmt.Println(a)

	// // Мнодественное присвоение значения
	// s, d, g := 1, 2, 3
	// s, d = d, s
	// fmt.Println(s, d, g)

	// messageFromFunc := sayHello("Maksim", 34)
	// printMessage("Вызов 1")
	// printMessage("Вызов 2")
	// printMessage("Вызов 3")
	// printMessage(messageFromFunc)

	// mes, entered := enterTheClub(19)
	// fmt.Println(mes)
	// fmt.Println(entered)

	// mes2, error := enterTheClubWithErr(14)
	// if error != nil {
	// 	log.Fatal(error)
	// 	return //	Если функция ничего не возвращает и имеет return это значит, что дойдя до
	// 	// места где return, функция просто остоновит свое выполнение
	// }
	// fmt.Println(mes2)

	fmt.Println(prediction("Pn"))
	fmt.Println(prediction("qqq"))

}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age)
}

// Множественные возвращаемые значения
func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "Welcome", true
	}
	return "Go out", false
}

func enterTheClubWithErr(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Welcome", nil
	} else if age >= 45 && age < 65 {
		return "This music for not for you", nil
	} else if age >= 65 {
		return "You are old", errors.New("Vwry old man")
	}
	return "Go out", errors.New("Very young")
}

func prediction(dayOfWeek string) (string, error) {
	// if dayOfWeek == "Понедельник" {
	// 	return "Monday"
	// } else if dayOfWeek == "Thusday" {
	// 	return "Thusday"
	// } else if dayOfWeek == "thirday" {
	// 	return "Sreda"
	// } else if dayOfWeek == "Chetverg" {
	// 	return "Chetverg"
	// }
	// return "No  day"

	switch dayOfWeek {
	case "Pn":
		return "Monday", nil
	case "vt":
		return "Thusday", nil
	case "Sr":
		return "Sreda", nil
	case "Scht":
		return "Chetverv", nil
	default:
		return "No day", errors.New("Not a day select")
	} // После записи default пропала ошибка того что не возвращается значение из функции

}
