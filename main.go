package main

import "fmt"

func main() {
	names := []string{
		"Max",
		"Vika",
		"Nastia",
		"Artem",
		"Dima",
		"Vlad",
		"Kostia",
		"Vitaliy",
		"Stas",
	}

	names = append(names, "Vitaliy")
	names = append(names, "Victoria")

	fmt.Println(names)

	var (
		parni   []int
		neparni []int
	)

	for i := 0; i <= 21; i++ {
		if i%2 == 0 {
			parni = append(neparni, i)
		}

	}
	fmt.Println(parni)
}
