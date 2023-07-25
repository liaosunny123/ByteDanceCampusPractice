package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := rand.Intn(maxNum)

	fmt.Print("Please input your guess: ")
	for {
		var guess int
		_, err := fmt.Scanf("%d\r\n", &guess)
		if err != nil {
			fmt.Print("Invalid input. Please enter an integer value:")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
