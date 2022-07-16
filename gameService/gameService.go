package gameService

import (
	"battleship/board"
	"battleship/player"
	"battleship/ship"
)

type GameStatus uint8

const (
	Continue GameStatus = 1
	End      GameStatus = 0
)

// var game GameStatus = Continue

func NewGameService(name string, ht, wd uint8) (*player.Player, *board.Board) {
	p1 := player.CreatePlayer(name)

	brd := board.CreateBoard(ht, wd)

	ship.GenerateRandomShips(brd)

	brd.PrintBoard()

	return p1, brd
}

func GamePlay(h, w uint8, p *player.Player, b *board.Board) GameStatus {

	p.IncreaseNoOfAttempts()
	if b.GetStatus(h, w) == board.Ship {
		b.SetStatus(board.Hit, h, w)
	} else if b.GetStatus(h, w) == board.Blank {
		b.SetStatus(board.Miss, h, w)
	}

	b.PrintBoard()

	return Continue
}
