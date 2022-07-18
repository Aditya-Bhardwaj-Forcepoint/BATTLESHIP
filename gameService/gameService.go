package gameService

import (
	"battleship/board"
	"battleship/player"
	"battleship/ship"
	"fmt"
)

type GameStatus uint8

const (
	Continue GameStatus = 1
	End      GameStatus = 0
)

type Cordinate struct {
	ht uint8
	wd uint8
}

var printMap = make(map[Cordinate]bool)

var shipCount uint8 = 0

func NewGameService(name string, ht, wd uint8) (*player.Player, *board.Board) {
	p1 := player.CreatePlayer(name)

	brd := board.CreateBoard(ht, wd)

	ship.GenerateRandomShips(brd)

	PrintBoard(brd)

	return p1, brd
}

func GamePlay(h, w uint8, p *player.Player, b *board.Board) GameStatus {

	p.IncreaseNoOfAttempts()
	if (h > b.GetBoardHeight()-1) || (h < 0) || (w > b.GetBoardWidth()-1) || (w < 0) {
		fmt.Println("\nIndices out of bounds. Try again...")
	} else if b.GetStatus(h, w) == board.Ship {
		printMap[Cordinate{h, w}] = true
		shipCount++
		b.SetStatus(board.Hit, h, w)
	} else if b.GetStatus(h, w) == board.Blank {
		printMap[Cordinate{h, w}] = true

		b.SetStatus(board.Miss, h, w)
	} else if b.GetStatus(h, w) == board.Miss {
		fmt.Print("\nYou have already missed at this place...")
	} else if b.GetStatus(h, w) == board.Hit {
		fmt.Println("\nYou have already hit at this place...")
	}

	PrintBoard(b)
	if shipCount == 15 {
		return End
	} else {
		return Continue
	}
}

var i, j uint8

// Prints the board for player
func PrintBoard(b *board.Board) {
	fmt.Print("\nAt present the warfield looks like :\n\n")

	for i = 0; i < b.GetBoardHeight(); i++ {
		fmt.Printf("%d\t", i)

		for j = 0; j < b.GetBoardWidth(); j++ {
			if printMap[Cordinate{i, j}] == true {
				fmt.Print(b.GetStatus(i, j), "\t")
			} else {
				fmt.Print(board.Blank, "\t")

			}
		}
		fmt.Println()
	}

	fmt.Print("\n\n")

	for j = 0; j < b.GetBoardWidth(); j++ { // Printing indices along the breadth
		fmt.Printf("\t%d", j)
	}

	fmt.Printf("\n\n")
}
