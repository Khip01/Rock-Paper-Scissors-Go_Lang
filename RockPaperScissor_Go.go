package main

import (
	"fmt"
	Console "fmt"
	"math"
	"math/rand"
	"os"
	System "os/exec"
	"strconv"
	"strings"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	textTitle := "---=--- ROCK PAPER SCISSORS ---=---"
	playAgain := true

	// Program Started
	clearScreen()

	generateTitleRandEffect(textTitle)

	for playAgain {
		showMenu("Choose Your Character!")

		inputUser()

		// Want to play again?
		Console.Print("Do you want to play again? :D (y/n)> ")
		var strInput string
		Console.Scan(&strInput)
		if !(strInput == "y" || strInput == "Y") {
			playAgain = false
			Console.Print("\n\nBye, thanks for playing!\n~Dev: Aakhif\n\n")
			time.Sleep(3 * time.Second)
		} else {
			clearScreen()
			Console.Println("========WELCOME TO THE GAME========")
			Console.Println("---=--- ROCK PAPER SCISSORS ---=---\n")
		}
	}

}

//// --- Method ZONE ! --- ////

// Clear Screen
func clearScreen() {
	cmd := System.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Title
func generateTitleRandEffect(text string) {
	Console.Println("=======WELCOME TO THE GAME========")

	for i := 0; i < len(text); i++ {
		isLoop := true

		// Menjalankan di goroutine
		go func() {
			time.Sleep(100 * time.Millisecond)
			isLoop = false
		}()

		for isLoop {
			randomInt := rand.Intn(26)

			Console.Printf("%c", charset[randomInt])
			time.Sleep(20 * time.Millisecond)
			Console.Print("\b \b")
		}

		Console.Printf("%c", text[i])
	}

	time.Sleep(1 * time.Second)
	Console.Print("\n\n")
}

// Show Menu
func showMenu(menuTitle string) {

	// Menu Title
	splitedWord := strings.Split(menuTitle, " ")
	// Check if string not null/empty
	if len(menuTitle) > 0 {
		for i := 0; i < len(splitedWord); i++ {
			Console.Print(splitedWord[i], " ")
			time.Sleep(300 * time.Millisecond)
		}
		time.Sleep(400 * time.Millisecond)
		Console.Print("\n")
	} else {
		Console.Println("Something Wrong with the code, check it again")
		return
	}

	// Menu Content
	Console.Println("1. Rock ")
	Console.Println("2. Paper")
	Console.Println("3. Scissors")
}

func inputUser() {
	devComment := [8]string{
		"Think carefully!",
		"Your choice determines your future!",
		"Hurry Choose!",
		"In RL this process happens in your head btw :D",
		"Give up? Press ctrl + C",
		"Never mind, you can't win",
		"Pray first before choosing :)",
		"Good luck!",
	}

	// Random Some Comment
	randomComment := rand.Intn(7)
	Console.Print("/", devComment[randomComment], " [Input Number, ex: 2]> ")

	// Input From User
	var inputStr string
	var inputInt int

	// Looping for input
	for {
		Console.Scan(&inputStr)

		// Check Input Error
		resInput, errCatch := strconv.Atoi(inputStr)
		if (errCatch != nil) || (1 > resInput || resInput > 3) {
			Console.Print("/You entered the wrong input! \n[Input MUST be a Number 1 - 3!, ex: 2]> ")
		} else {
			inputInt = resInput
			break
		}
	}
	Console.Print("\n")

	// The Bot/Opponent is randomizing his selection
	randomBotsChoice := rand.Intn(3)

	/*
		1 1 =  0 Draw	| 0	| =
		1 2 = -1 Lose	| 1	| <
		1 3 = -2 Win	| 1	| <

		2 1 =  1 Win	| 0	| >
		2 2 =  0 Draw	| 0	| =
		2 3 = -1 Lose	| 2	| <

		3 1 =  2 Lose	| 0	| >
		3 2 =  1 Win	| 1	| >
		3 3 =  0 Draw	| 0	| =

	*/

	// For Loading Animation
	loadingAnimation()

	if inputInt == (randomBotsChoice + 1) {
		Console.Println("YOUR Weapon: ", inputInt)
		printASCII(inputInt)
		Console.Println("BOT  Weapon: ", randomBotsChoice)
		printASCII(randomBotsChoice + 1)
		Console.Println("DRAW!!")
	} else {
		getTheWinner(inputInt, randomBotsChoice+1)
	}

}

func loadingAnimation() {
	loadingText := "And the result is"

	Console.Print(loadingText)
	for i := 1; i <= 3; i++ {
		for i := 1; i <= 3; i++ {
			Console.Print(". ")
			time.Sleep(400 * time.Millisecond)
		}
		for i := 1; i <= 6; i++ {
			Console.Print("\b \b")
			time.Sleep(10 * time.Millisecond)
		}
	}
	for i := 1; i <= len(loadingText); i++ {
		Console.Print("\b \b")
	}
}

func printASCII(weapon int) {
	switch weapon {
	case 1:
		Rock.Print()
		break
	case 2:
		Paper.Print()
		break
	case 3:
		Scissors.Print()
		break
	}
}

func getTheWinner(userInput int, botChoice int) {
	// Get Difference
	difference := int(math.Abs(float64(userInput - botChoice)))

	switch userInput {
	case 1:
		if difference > 1 {
			Console.Println("YOUR Weapon: ", userInput)
			printASCII(userInput)
			Console.Println("BOT  Weapon: ", botChoice)
			printASCII(botChoice)
			Console.Println("You WIN!!")
		} else {
			Console.Println("YOUR Weapon: ", userInput)
			printASCII(userInput)
			Console.Println("BOT  Weapon: ", botChoice)
			printASCII(botChoice)
			Console.Println("You LOSE!!")
		}
		break
	case 2:
		if userInput > botChoice {
			Console.Println("YOUR Weapon: ", userInput)
			printASCII(userInput)
			Console.Println("BOT  Weapon: ", botChoice)
			printASCII(botChoice)
			Console.Println("You WIN!!")
		} else {
			Console.Println("YOUR Weapon: ", userInput)
			printASCII(userInput)
			Console.Println("BOT  Weapon: ", botChoice)
			printASCII(botChoice)
			Console.Println("You LOSE!!")
		}
		break
	case 3:
		if difference > 1 {
			Console.Println("YOUR Weapon: ", userInput)
			printASCII(userInput)
			Console.Println("BOT  Weapon: ", botChoice)
			printASCII(botChoice)
			Console.Println("You LOSE!!")
		} else {
			Console.Println("YOUR Weapon: ", userInput)
			printASCII(userInput)
			Console.Println("BOT  Weapon: ", botChoice)
			printASCII(botChoice)
			Console.Println("You WIN!!")
		}
		break
	}
}

// For ASCII ilustration (ASCII Source: https://gist.github.com/wynand1004/b5c521ea8392e9c6bfe101b025c39abe)

type RpsAscii struct {
	ASCII string
}

var (
	// Rock
	Rock = RpsAscii{
		ASCII: `
	    _______
	---'   ____)
	      (_____)
	      (_____)
	      (____)
	---.__(___)
	`,
	}

	// Paper
	Paper = RpsAscii{
		ASCII: `
	     _______
	---'    ____)____
	           ______)
	          _______)
	         _______)
	---.__________)
	`,
	}

	// Scissors
	Scissors = RpsAscii{
		ASCII: `
	    _______
	---'   ____)____
	          ______)
	       __________)
	      (____)
	---.__(___)
	`,
	}
)

func (a RpsAscii) Print() {
	fmt.Println(a.ASCII)
}
