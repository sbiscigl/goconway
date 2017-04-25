package board

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/sbiscigl/conways/cell"
)

/*Board type for game board*/
type Board struct {
	Width     int
	Height    int
	GameBoard [][]cell.Cell
}

/*New returns new isntance of game board*/
func New(w int, h int) *Board {
	board := &Board{
		Height:    w,
		Width:     h,
		GameBoard: make([][]cell.Cell, w),
	}
	for i := range board.GameBoard {
		board.GameBoard[i] = make([]cell.Cell, h)
	}
	return board
}

/*Init gives a 30 percent chance a square lives*/
func (b *Board) Init(chance int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for x, arr := range b.GameBoard {
		for y := range arr {
			a := rand.Intn(100)
			if a <= chance {
				arr[y].Alive = true
				arr[y].Ypos = x
				arr[y].Xpos = y
			} else {
				arr[y].Alive = false
				arr[y].Ypos = x
				arr[y].Xpos = y
			}
		}
	}
}

/*Draw draws the game board*/
func (b *Board) Draw() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for i := 0; i < b.Width+2; i++ {
		fmt.Print("–")
	}
	fmt.Print("\n")
	for _, arr := range b.GameBoard {
		fmt.Print("|")
		for _, alive := range arr {
			if alive.Alive == true {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("|")
		fmt.Print("\n")
	}
	for i := 0; i < b.Width+2; i++ {
		fmt.Print("–")
	}
	fmt.Print("\n")
}
