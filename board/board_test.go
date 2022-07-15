package board

import (
	"fmt"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	brd := make([][]Status, 5) // Make outer slice of size == breadth
	// Make inner slice of size == length
	for i := range brd {
		brd[i] = make([]Status, 10)
	}
	// Assign each cell the status "Blank"
	for i := range brd {
		for j := range brd[i] {
			brd[i][j] = Blank
		}
	}

	expectedObject := Board{
		length:     10,
		breadth:    5,
		PtrToBoard: &brd,
	}

	outputObject := *CreateBoard(10, 5)

	if expectedObject.length != outputObject.length || expectedObject.breadth != outputObject.breadth {
		t.Errorf("The outputObject : %+v != %+v : The expectedObject \n", outputObject, expectedObject)
		t.Errorf("The ptr to board for outputObject contains : %+v \n", (*outputObject.PtrToBoard))
		t.Errorf("The ptr to board for expectedObject contains : %+v \n", (*expectedObject.PtrToBoard))

	} else {
		fmt.Println("Both the outputObject and expectedObject are the same.")
		fmt.Printf("The ptr to board for outputObject contains : %+v \n", (*outputObject.PtrToBoard))
	}
}

func TestGetBoardLength(t *testing.T) {
	brd := Board{
		length: 100,
	}

	var expectedLegth uint8 = 100

	actualLength := brd.GetBoardLength()

	if expectedLegth != actualLength {
		t.Errorf("The expectedLegth : %d != %d : The actualLength \n", expectedLegth, actualLength)
	} else {
		fmt.Printf("The expectedLegth : %d == %d : The actualLength \n", expectedLegth, actualLength)
		fmt.Println("Both the expectedLegth and actualLength are the same.")
	}
}

func TestGetBoardBreadth(t *testing.T) {
	brd := Board{
		breadth: 50,
	}

	var expectedBreadth uint8 = 50

	actualBreadth := brd.GetBoardBreadth()

	if expectedBreadth != actualBreadth {
		t.Errorf("The expectedBreadth : %d != %d : The actualBreadth \n", expectedBreadth, actualBreadth)
	} else {
		fmt.Printf("The expectedBreadth : %d == %d : The actualBreadth \n", expectedBreadth, actualBreadth)
		fmt.Println("Both the expectedBreadth and actualBreadth are the same.")
	}
}

func TestPrintBoard(t *testing.T) {
	var flag uint8 = 0
	// Creating our test board
	expectedBoard := make([][]Status, 5)
	for i := range expectedBoard {
		expectedBoard[i] = make([]Status, 10)
	}
	// Setting Status for the test board
	for i := range expectedBoard {
		for j := range expectedBoard[i] {
			expectedBoard[i][j] = Blank
		}
	}

	outputBoard := *CreateBoard(10, 5)

	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if expectedBoard[i][j] != (*outputBoard.PtrToBoard)[i][j] {
				flag = 1
			}
		}
	}

	if flag != 0 {
		t.Errorf("The Expected Board and Output Board aren't the same.")
		t.Errorf("Expected Board is : %+v \n", (expectedBoard))
		t.Errorf("Output Board is : %+v \n", (*outputBoard.PtrToBoard))
	} else {
		fmt.Println("Both the boards are same.")
		fmt.Printf("Expected Board is : %+v \n", (expectedBoard))
		fmt.Printf("Output Board is : %+v \n", (*outputBoard.PtrToBoard))
	}
}

func TestGetStatus(t *testing.T) {
	// Creating a 2D array and assigning every cell of it the status "Blank"
	arr := make([][]Status, 20)
	for i := range arr {
		arr[i] = make([]Status, 40)
	}
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = Blank
		}
	}

	brd := Board{
		length:     40,
		breadth:    20,
		PtrToBoard: &arr,
	}

	if brd.GetStatus(17, 31) != Blank {
		t.Errorf("The status returned is not the expected status.")
		t.Errorf("Resultant Status : %+v \n", brd.GetStatus(17, 31))
		t.Errorf("Expected Status : %+v \n", Blank)
	} else {
		fmt.Println("Both the Status are same.")
		fmt.Printf("Resultant Status : %+v \n", brd.GetStatus(17, 31))
		fmt.Printf("Expected Status : %+v \n", Blank)
	}
}

func TestSetStatus(t *testing.T) {
	// Creating a 2D array and assigning every cell of it the status "Blank"

	arr := make([][]Status, 20)
	for i := range arr {
		arr[i] = make([]Status, 40)
	}
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = Blank
		}
	}

	brd := Board{
		length:     40,
		breadth:    20,
		PtrToBoard: &arr,
	}

	brd.SetStatus(Hit, 19, 39)

	if brd.GetStatus(19, 39) != Hit {
		t.Errorf("The status set is not the expected status.")
		t.Errorf("Resultant Status : %+v \n", brd.GetStatus(19, 39))
		t.Errorf("Expected Status : %+v \n", Hit)
	} else {
		fmt.Println("Both the Status are same.")
		fmt.Printf("Resultant Status : %+v \n", brd.GetStatus(19, 39))
		fmt.Printf("Expected Status : %+v \n", Hit)
	}
}
