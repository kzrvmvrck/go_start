package main

import "fmt"

func main() {
	fmt.Println(findMin(1, 2, 3, -4, -6, 233, -55))
	fmt.Println(findMax(1, 2, 3, -4, -6, 233, -55))

	// Анонимные функции
	func() {
		fmt.Println("Anonimus func")
	}()

	// Замыкания и практическое применение анонимныъ функций
	// замыкание сохраняет состояние функции в который был вызван
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())

	fmt.Println(inc())
}

func findMin(numbers ...int) int { // numbers по сути это массив, будем работать с ним
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
func findMax(numbers ...int) int { // numbers по сути это массив, будем работать с ним
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}

// Это замыкание
func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}
