package main

import (
	"fmt"

	"github.com/alphaaleph/qixg/data"
	"github.com/alphaaleph/qixg/game"
	"github.com/hajimehoshi/ebiten"
)

func init() {
	//load the game data
	data.Read()
}

func main() {

	//start running the 2D framework
	if err := ebiten.Run(game.Loop, int(data.Memory.ScreenWidth), int(data.Memory.ScreenHeight), 2, "QIX"); err != nil {
		fmt.Println("panic")
		panic(err)
	}

	//save the scores
	fmt.Println("hey")
	data.Save()
}
