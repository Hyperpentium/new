package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"project_1/domain"
	"sort"
	"strconv"
	"time"
)

const (
	points            int = 10
	pointsPerQuestion int = 10
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у найкращій грі! Йде завантаження...")
	time.Sleep(1 * time.Second)

	for {
		menu()
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users := getUsers()
			users = append(users, u)
			sortAndSave(users)
		case "2":
			users := getUsers()
			for _, u := range users {
				fmt.Printf("Id: %v, Name: %s, Time: %v\n", u.Id, u.Name, u.Time)
			}
		case "3":
			return
		default:
			fmt.Println("Не коректний вибір")
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	fmt.Println("Підготуватись!")

	for i := 3; i > 0; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("ВПЕРЕД!")
	myPoints := 0
	pointsPerGame := points
	now := time.Now()
	for pointsPerGame > 0 {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			if ansInt == res {
				pointsPerGame -= pointsPerQuestion
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! У тебе: %v очок!\n", myPoints)
				fmt.Printf("Залишилось зібрати: %v очок!\n", pointsPerGame)
			} else {
				fmt.Println("Не правильно! Спробуй ще!")
			}
		}
	}
	then := time.Now()
	spent := then.Sub(now)

	fmt.Printf("Ти топ! Впорався за %s\n", spent)

	fmt.Println("Введіть своє ім'я: ")
	playerName := ""
	fmt.Scan(&playerName)

	// var user domain.User
	// user.Id = id
	// user.Name = playerName
	// user.Time = spent

	user := domain.User{
		Id:   id,
		Name: playerName,
		Time: spent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
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
		fmt.Printf("Error: %s", err)
		return
	}
}

func getUsers() []domain.User {
	info, err := os.Stat("users.json")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}

	var users []domain.User
	if info.Size() != 0 {
		file, err := os.Open("users.json")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}

		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil
		}
	}

	return users
}
