// This will act as caller and printer to our game
package main

import (
	"battleship/gameService"
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

		fmt.Print("Where do you want to attack : ")
		fmt.Scanln(&h)
		fmt.Scanln(&w)

		s := gameService.GamePlay(h, w, plr, brd)

		if s == gameService.End {
			break
		}
	}

	fmt.Println("Congratulations")
	fmt.Println("You finished the game in :", plr.GetNoOfAttempts())
}
