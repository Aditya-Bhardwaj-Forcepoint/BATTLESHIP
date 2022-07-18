package ship

import (
	"battleship/components/board"
	"testing"
)

func TestGenerateRandomShip(t *testing.T) {
	var flag uint8 = 0

	b := board.CreateBoard(10, 10)

	GenerateRandomShips(b)

	for _, ship := range AllShips {
		if (*b.PtrToBoard)[ship.yCoord][ship.xCoord] != board.Ship {
			flag = 1
			break
		}
	}

	if flag == 1 {
		t.Errorf("Ships were not created properly.")
	}

}
