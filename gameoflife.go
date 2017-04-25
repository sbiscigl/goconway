package main

import (
	"github.com/sbiscigl/conways/board"
	"github.com/sbiscigl/conways/game"
	"github.com/sbiscigl/conways/rules"
)

func main() {
	b := board.New(40, 140)
	b.Init()
	game.New(b, rules.New()).Tick()
}
