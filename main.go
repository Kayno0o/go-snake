package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kayno0o/go/snake/app/game"
)

func main() {
	app := &game.Game{
		Board: game.Board{
			W:        25,
			H:        25,
			CellSize: 30,
		},
		Delay: 250,
	}

	app.Init()

	w, h := app.Board.WindowSize()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Snake !")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
