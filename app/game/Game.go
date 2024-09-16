package game

import (
	"image/color"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kayno0o/go/snake/app/utils"
)

var (
	headerSize   = 50
	bgColor      = color.RGBA{50, 0, 150, 255}
	cellColor    = color.RGBA{245, 230, 255, 255}
	dataDirPath  = "./data"
	dataFilePath = filepath.Join(dataDirPath, "data.json")
)

type Pos struct {
	X int
	Y int
}

type Game struct {
	Board Board

	Delay    int
	LastMove time.Time

	Score int

	dirX int
	dirY int

	data GameData

	// "menu" | "playing" | "lose"
	State string
}

func (g *Game) Init() {
	g.data.LoadData()

	g.Board.Init()

	utils.InitDraw()

	g.dirX = 1
	g.State = "play"
}

func (g *Game) Move() *Pos {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.dirX = 0
		g.dirY = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.dirX = 0
		g.dirY = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.dirX = -1
		g.dirY = 0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.dirX = 1
		g.dirY = 0
	}

	now := time.Now()
	if g.LastMove.UnixMilli()+int64(g.Delay) > now.UnixMilli() {
		return nil
	}

	g.LastMove = time.Now()

	return g.Board.MoveSnake(g.dirX, g.dirY)
}

func (g *Game) checkDeath(newPos *Pos) {
	if g.Board.CheckSnake(newPos) {
		g.State = "lose"
		g.data.SaveData()
	}
}

func (g *Game) checkApple(newPos *Pos) {
	if index := g.Board.CheckApple(newPos); index != -1 {
		g.Score++

		g.Board.MoveApple(index)
		return
	}

	g.Board.snake = g.Board.snake[1:]
}

func (g *Game) Play() {
	newPos := g.Move()
	if newPos == nil {
		return
	}

	g.checkApple(newPos)
	g.checkDeath(newPos)
}

func (g *Game) Update() error {
	g.Play()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == "play" {
		g.Board.Draw(screen)

		g.data.HighScore = utils.Max(g.data.HighScore, g.Score)

		utils.DrawText(screen, 6, 8, 16, "     Score:"+strconv.FormatInt(int64(g.Score), 10), bgColor)
		utils.DrawText(screen, 6, 30, 16, "High Score:"+strconv.FormatInt(int64(g.data.HighScore), 10), bgColor)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Board.WindowSize()
}
