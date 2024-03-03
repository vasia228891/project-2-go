package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"project1/domain"
	"sort"
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
	users = append(users, domain.User{Id: 1, Name: "Mykola", Time: 5 * time.Second})
	users = append(users, domain.User{Id: 2, Name: "Vasyl", Time: 3 * time.Second})
	users = append(users, domain.User{Id: 3, Name: "Sokrat", Time: 8 * time.Second})

	sortAndSaveUsers(users)

	// 	for {
	// 		menu()
	// 		var punct string
	// 		fmt.Scan(&punct)

	// 		switch punct {
	// 		case "1":
	// 			u := play()
	// 			if u.Id != 0 {
	// 				users = append(users, u)

	// 			}
	// 		case "2":
	// 			fmt.Println("Рейтинг користувачів:")
	// 			for _, u := range users {
	// 				fmt.Printf("id: %v, name: %s, time: %s\n",
	// 					u.Id,
	// 					u.Name,
	// 					u.Time,
	// 				)
	// 			}
	// 		case "3":
	// 			return
	// 		default:
	// 			fmt.Println("Зробіть правильний вибір")
	// 		}
	// 	}
}

func sortAndSaveUsers(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	func showUserRate() []domain.User {
		info, err := os.Stat("users.json")
		if err != nil && !os.IsNotExist(err){
			return nil 
		}
	}

	var users []domain.User
	if info != nil && info.Size() !=0{
		file, err := os.Open("users.json")
		if err != nil {
			return nil
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}(file)
	}
func showUserRate() []domain.User {
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil
	}
}

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Сталась помилка T_T: %s\n", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {

		fmt.Printf("Сталась помилка T_T: %s\n", err)
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
	var res int
	for points > 0 {
		mode := rand.Intn(3)
		switch mode {
		case 0:
			x, y := rand.Intn(100), rand.Intn(100)
			fmt.Printf("%v + %v = ", x, y)
			res = x + y

		case 1:
			a, b := rand.Intn(10), rand.Intn(10)
			fmt.Printf("%v * %v = ", a, b)
			res = a * b

		case 2:
			z, v := rand.Intn(10), rand.Intn(10)
			fmt.Printf("%v - %v = ", z, v)
			res = z - v

		}

		var answer string
		fmt.Scan(&answer)

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

	var name string
	fmt.Scan(&name)

	if name == "" {
		fmt.Println("Не коректне ім'я. Дані не збережені.")
		return domain.User{}
	}

	user := domain.User{
		Id:        id,
		Name:      name,
		Time:      timeSpent,
		LoginTime: time.Now(), // Додавання часу входу для сортування
	}
	id++

	return user
}
