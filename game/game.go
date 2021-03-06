package game

import (
	"time"

	"github.com/sbiscigl/goconway/board"
	"github.com/sbiscigl/goconway/rules"
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
		time.Sleep(time.Millisecond * 200)
		for x, arr := range g.GameBoard.GameBoard {
			for y := range arr {
				g.RuleRunner.IsAlive(&g.GameBoard.GameBoard[x][y], g.GameBoard)
			}
		}
	}
}
