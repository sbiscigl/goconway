package game

import (
	"time"

	"github.com/sbiscigl/conways/board"
	"github.com/sbiscigl/conways/rules"
)

/*Game type for a game runner*/
type Game struct {
	GameBoard  *board.Board
	RuleRunner *rules.Rules
}

/*New constructor for game*/
func New(b *board.Board, r *rules.Rules) *Game {
	return &Game{
		GameBoard:  b,
		RuleRunner: r,
	}
}

/*Tick run the game*/
func (g *Game) Tick() {
	for {
		g.GameBoard.Draw()
		time.Sleep(time.Second)
		for x, arr := range g.GameBoard.GameBoard {
			for y := range arr {
				g.RuleRunner.IsAlive(&g.GameBoard.GameBoard[x][y], g.GameBoard)
			}
		}
	}
}
