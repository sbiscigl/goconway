package rules

import (
	"github.com/sbiscigl/goconway/board"
	"github.com/sbiscigl/goconway/cell"
)

/*Rules type for game rules*/
type Rules struct {
	//No members static class
}

/*New returns the rules instance*/
func New() *Rules {
	return &Rules{}
}

/*IsAlive tells us if a cell is alive*/
func (r *Rules) IsAlive(c *cell.Cell, b *board.Board) {
	num := numberAlive(*c, *b)
	//Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
	if num < 2 {
		c.Alive = false
	}
	//Any live cell with two or three live neighbours lives on to the next generation.
	if num == 2 || num == 3 { /*Nothing*/
	}
	//Any live cell with more than three live neighbours dies, as if by overpopulation.
	if num > 3 {
		c.Alive = false
	}
	//Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if num == 3 {
		c.Alive = true
	}
}

func numberAlive(c cell.Cell, b board.Board) int {
	alive := 0
	if c.Ypos-1 >= 0 && b.GameBoard[c.Ypos-1][c.Xpos].Alive == true {
		alive++
	}
	if c.Ypos-1 >= 0 && c.Xpos-1 >= 0 && b.GameBoard[c.Ypos-1][c.Xpos-1].Alive == true {
		alive++
	}
	if c.Ypos-1 >= 0 && c.Xpos+1 < b.Width && b.GameBoard[c.Ypos-1][c.Xpos+1].Alive == true {
		alive++
	}

	if c.Xpos-1 >= 0 && b.GameBoard[c.Ypos][c.Xpos-1].Alive == true {
		alive++
	}
	if c.Xpos+1 < b.Width && b.GameBoard[c.Ypos][c.Xpos+1].Alive == true {
		alive++
	}

	if c.Ypos+1 < b.Height && b.GameBoard[c.Ypos+1][c.Xpos].Alive == true {
		alive++
	}
	if c.Ypos+1 < b.Height && c.Xpos-1 >= 0 && b.GameBoard[c.Ypos+1][c.Xpos-1].Alive == true {
		alive++
	}
	if c.Ypos+1 < b.Height && c.Xpos+1 < b.Width && b.GameBoard[c.Ypos+1][c.Xpos+1].Alive == true {
		alive++
	}
	return alive
}
