// This package will have information about our board and it's functions
package board

import "fmt"

// Structure of the Board
type Board struct {
	length     uint8
	breadth    uint8
	PtrToBoard *([][]Status)
}

// Status variable will define at any instance of time the status of current cell
type Status string

// Enum values taken by status
const (
	Miss  Status = "_"
	Hit   Status = "@"
	Blank Status = "-"
	Ship  Status = "S"
)

// Creates a new Board object
func CreateBoard(ln, bh uint8) *Board {

	brd := make([][]Status, bh)
	// Now we create the subslices inside each index of the main (breadth) slice.
	for i := range brd {
		brd[i] = make([]Status, ln)
	}

	// Now we initialize each cell of our board with "Blank" status
	for i := range brd {
		for j := range brd[i] {
			brd[i][j] = Blank
		}
	}

	return &Board{
		length:     ln,
		breadth:    bh,
		PtrToBoard: &brd,
	}
}

// Returns the length of the board object
func (b *Board) GetBoardLength() uint8 {
	return b.length
}

// Returns the breadth of the board object
func (b *Board) GetBoardBreadth() uint8 {
	return b.breadth
}

// Prints the board
func (b *Board) PrintBoard() {
	fmt.Println("At present the warfield looks like :")
	for i := 0; i < int(b.length); i++ {
		for j := 0; j < int(b.breadth); j++ {
			fmt.Print((*b.PtrToBoard)[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Gets the status of the ( dx, dy)th cell of our board
func (b *Board) GetStatus(dy, dx uint8) Status {
	return (*b.PtrToBoard)[dy][dx]
}

// Sets the status of the ( dx, dy)th cell of our board
func (b *Board) SetStatus(newStat Status, dy, dx uint8) {
	(*b.PtrToBoard)[dy][dx] = newStat
}
