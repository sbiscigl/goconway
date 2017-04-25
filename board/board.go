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
	Height    int
	Width     int
	GameBoard [][]cell.Cell
}

/*New returns new isntance of game board*/
func New(w int, h int) *Board {
	board := &Board{
		Width:     w,
		Height:    h,
		GameBoard: make([][]cell.Cell, w),
	}
	for i := range board.GameBoard {
		board.GameBoard[i] = make([]cell.Cell, h)
	}
	return board
}

/*Init gives a 30 percent chance a square lives*/
func (b *Board) Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	for x, arr := range b.GameBoard {
		for y := range arr {
			a := rand.Intn(9)
			if a <= 1 {
				arr[y].Alive = true
				arr[y].Xpos = x
				arr[y].Ypos = y
			} else {
				arr[y].Alive = false
				arr[y].Xpos = x
				arr[y].Ypos = y
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
