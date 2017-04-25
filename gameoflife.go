package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sbiscigl/conways/board"
	"github.com/sbiscigl/conways/game"
	"github.com/sbiscigl/conways/rules"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ")
		fmt.Println("conways <Width> <Height>")
	}
	w, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("width not number"))
	}
	h, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(fmt.Sprintf("height not number"))
	}
	b := board.New(w, h)
	b.Init()
	game.New(b, rules.New()).Tick()
}
