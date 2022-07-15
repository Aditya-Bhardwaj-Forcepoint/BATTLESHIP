// This will act as caller and printer to our game
package main

import "fmt"

func main() {
	fmt.Println("Welcome to Battle Ship")
	fmt.Println()
	var ln uint8
	fmt.Print("Please enter the length of the board :")
	fmt.Scanln(&ln)
	fmt.Println()

	var bh uint8
	fmt.Print("Please enter the breadth of the board :")
	fmt.Scanln(&bh)
	fmt.Println()
}
