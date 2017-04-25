package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sbiscigl/goconway/board"
	"github.com/sbiscigl/goconway/game"
	"github.com/sbiscigl/goconway/rules"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ")
		fmt.Println("goconway <Width> <Height> <Percent chance of initial life of cell (out of 100)>")
		panic(fmt.Sprintf("not enough args"))
	}
	w, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("width not number"))
	}
	h, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(fmt.Sprintf("height not number"))
	}
	chance, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(fmt.Sprintf("chance not number"))
	}
	if chance < 0 || chance > 100 {
		panic(fmt.Sprintf("chance is not in valid 0 - 100 range"))
	}
	b := board.New(w, h)
	b.Init(chance)
	game.New(b, rules.New()).Tick()
}
