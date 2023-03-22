// guess game - a game, where you need to guess a number to win. You have limited amount of attempts
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// pseudo random int generator
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.\nCan you guess it?")

	reader := bufio.NewReader(os.Stdin) // создаем bufio.reader для чтения ввода с клавиатуры.

	success := false

	for i := 1; i <= 10; i++ {
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // удаление символа новой строки "\n"
		quess, err := strconv.Atoi(input) // convert string to int

		if err != nil {
			log.Fatal(err)
		}

		// make an iterator to simulate life
		if quess < target {
			fmt.Println("Oops. Your guess was LOW.", 10-i, "Lives left.")
		} else if quess > target {
			fmt.Println("Oops. Your guess was HIGH.", 10-i, "Lives left.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("You lose! The number was", target)
	}
}
