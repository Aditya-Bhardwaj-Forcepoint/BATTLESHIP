// This will act as caller and printer to our game
package main

import (
	"battleship/services/gameService"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Battle Ship")
	fmt.Println()

	var name string
	fmt.Print("Please enter the name of the player : ")
	fmt.Scanln(&name)
	fmt.Println()

	var ht uint8
	fmt.Print("Please enter the height of the board : ")
	fmt.Scanln(&ht)
	fmt.Println()

	var wd uint8
	fmt.Print("Please enter the width of the board : ")
	fmt.Scanln(&wd)
	fmt.Println()

	plr, brd := gameService.NewGameService(name, ht, wd)

	var h uint8
	var w uint8
	for { // Ask the player where he wants to attack untill all the ships are destroyed...

		fmt.Println("Where do you want to attack : ")
		fmt.Println("Enter height coord:")
		fmt.Scanln(&h)
		if h < 0 || h > 255 {
			fmt.Println("Invalid Input")
			continue
		}
		fmt.Println("Enter width coord:")
		fmt.Scanln(&w)
		if w < 0 || w > 255 {
			fmt.Println("Invalid Input")
			continue
		}

		s := gameService.GamePlay(h, w, plr, brd)

		if s == gameService.End {
			break
		}
	}

	fmt.Println("Congratulations")
	fmt.Println("You finished the game in :", plr.GetNoOfAttempts())
}
