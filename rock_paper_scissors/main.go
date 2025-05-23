package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const rounds = 3
	rand.Seed(time.Now().UnixNano()) // Add seed for random number generation
	fmt.Println("Let's play rock, paper and scissors... you have", rounds, "rounds")

	for i := 0; i < rounds; i++ {
		computerChoiceNum := rand.Intn(3) + 1 // Generates 1, 2, or 3
		var computerChoice string

		switch computerChoiceNum {
		case 1:
			computerChoice = "Rock"
		case 2:
			computerChoice = "Paper"
		case 3:
			computerChoice = "Scissors"
		}

		var playerChoice string
		fmt.Println("Enter your choice (Rock, Paper, Scissors):")
		fmt.Scanln(&playerChoice)

		fmt.Printf("Computer choice: %s\n", computerChoice)

	
		if playerChoice == computerChoice {
			fmt.Println("It's a tie")
		} else if (playerChoice == "Rock" && computerChoice == "Scissors") ||
			(playerChoice == "Paper" && computerChoice == "Rock") ||
			(playerChoice == "Scissors" && computerChoice == "Paper") {
			fmt.Println("You win")
		} else {
			fmt.Println("Computer wins")
		}
	}
}
