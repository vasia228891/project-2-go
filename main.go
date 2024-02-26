package main

import (
	"fmt"
	"math/rand"
	"project1/domain"
	"strconv"
	"time"
)

const (
	requiredPoints    = 10
	pointsPerQuestion = 5
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у найкращій грі всіх часів!")
	time.Sleep(1 * time.Second)

	var users []domain.User

	menu()
	for {
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			if u.Id != 0 {
				users = append(users, u)
			}
		case "2":
			for _, u := range users {
				fmt.Printf("id: %v, name: %s, time: %s",
					u.Id,
					u.Name,
					u.Time,
				)
			}
		case "3":
			return
		default:
			fmt.Println("Зробіть правильний вибір")
		}
		menu()
	}
}

func menu() {
	fmt.Println("1.Почати гру")
	fmt.Println("2.Переглянути рейтинг")
	fmt.Println("3.Залишити гру")

}

func play() domain.User {
	fmt.Println("Гра почнеться через...")
	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0

	points := requiredPoints

	now := time.Now()
	for points > 0 {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ?", x, y)

		var answer string
		fmt.Scan(&answer)

		res := x + y
		ansInt, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			if res == ansInt {
				points -= pointsPerQuestion
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! Залишилось набрати %v балів!\n", points)
				fmt.Printf("У вас зараз %v балів.\n", myPoints)
			} else {
				fmt.Println("Упс... Спробуй ще!")
			}
		}
	}
	then := time.Now()
	timeSpent := then.Sub(now)

	fmt.Printf("Ура, вітаємо, ви впорались за %v\n", timeSpent)
	fmt.Println("Введіть ім'я гравця:")

	name := ""

	fmt.Scan(&name)

	if name == " " {
		fmt.Println("не коректне імя дані не збережені")
		return domain.User{}
	}

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}
