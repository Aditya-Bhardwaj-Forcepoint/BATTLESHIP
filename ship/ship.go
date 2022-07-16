// This file will contain info about ship structure and all it's functions
package ship

import (
	"battleship/board"
	"math/rand"
	"time"
)

type ShipInfo struct {
	orient uint8
	xCoord uint8
	yCoord uint8
}

// An array/struct to store the address of first compartment of each ship
var AllShips []ShipInfo

func GenerateRandomShips(b *board.Board) {
	// This array contains the sizes of the ships
	shipSizes := []uint8{1, 2, 3, 4, 5}

	for size := range shipSizes {
		size++ // As range() returns values in the interval [o, size)

		// We genereate a random number between 0 & 1
		// Here, 0 -> Horizontally Placed Ship
		// And,	 1 -> Verically Placed Ship

		rand.Seed(time.Now().UnixNano())
		orient := rand.Intn(2)

		if orient == 0 { // Horizontally placed ship
			// Creating a horizontally placed ship is simple as we just create an array/ship in the slice of the given index
			// We choose a random index of the breadth slice and length slice
		pickIndex:
			rand.Seed(time.Now().UnixNano())
			dx := rand.Intn(int(b.GetBoardHeight()))

			rand.Seed(time.Now().UnixNano())
			dy := rand.Intn(int(b.GetBoardWidth()))

			var flag1 uint8 = 0

			if (uint8(dx) + uint8(size)) < b.GetBoardWidth() { // Checking if the sum(size + indices) is out of bounds

				for i := 0; i < size; i++ {
					if (*b.PtrToBoard)[dy][dx+i] == board.Ship { // Checking if any indices of our chosen slice is already occupied by another ship.
						flag1 = 1
					}
				}

				if flag1 != 1 { // The choosen index and it's following cells are empty. So, we can create our new ship there.
					// Making the ship and storing it in AllShips slice
					// ship := make([]*board.Status, size)
					for i := 0; i < size; i++ {
						(*b.PtrToBoard)[dy][dx+i] = board.Ship
					}
					shp := ShipInfo{
						orient: 0,
						xCoord: uint8(dx),
						yCoord: uint8(dy),
					}
					AllShips = append(AllShips, shp)
				} else { // The choosen slice is already occupied by another ship
					goto pickIndex
				}
			} else { // The randomly choosen index is out of bounds with the size
				goto pickIndex
			}

		} else if orient == 1 { // Vertically places ship
			// Horizontal index where the head of our ship will be placed
		pickIndex2:
			rand.Seed(time.Now().UnixNano())
			dx := rand.Intn(int(b.GetBoardHeight()))

			rand.Seed(time.Now().UnixNano())
			dy := rand.Intn(int(b.GetBoardWidth()))

			if (dy + size) < int(b.GetBoardHeight()) { // There is enough size for us to create our ship
				// shipVert := make([]*board.Status, size)
				var flag2 uint8 = 0

				for i := 0; i < size; i++ {

					if (*b.PtrToBoard)[dy+i][dx] == board.Ship {
						flag2 = 1
					}

				}

				if flag2 != 1 {
					for i := 0; i < size; i++ {
						(*b.PtrToBoard)[dy+i][dx] = board.Ship
					}
					shp := ShipInfo{
						orient: 1,
						xCoord: uint8(dx),
						yCoord: uint8(dy),
					}
					AllShips = append(AllShips, shp)
				} else {
					goto pickIndex2
				}
			} else {
				goto pickIndex2
			}
		}
	}
}
