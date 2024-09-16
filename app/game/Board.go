package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Board struct {
	snake  []Pos
	apples []Pos

	W int
	H int

	CellSize int
}

func (b *Board) Init() {
	b.snake = make([]Pos, 4)
	for i := range b.snake {
		b.snake[i] = Pos{X: b.W/2 - len(b.snake) + i, Y: b.H / 2}
	}

	b.apples = make([]Pos, 0)
	b.addApple()
}

func (b *Board) WindowSize() (int, int) {
	return b.W * b.CellSize, b.H*b.CellSize + headerSize
}

func (b *Board) Draw(screen *ebiten.Image) {
	w, h := b.WindowSize()
	vector.DrawFilledRect(screen, 0, 0, float32(w), float32(h), bgColor, false)

	vector.DrawFilledRect(screen, 2, 2, float32(w-4), float32(headerSize-3), color.RGBA{255, 255, 255, 255}, false)

	b.DrawCells(screen)

	for _, cell := range b.snake {
		vector.DrawFilledRect(screen, float32(cell.X*b.CellSize)+1, float32(cell.Y*b.CellSize)+1+float32(headerSize), float32(b.CellSize)-2, float32(b.CellSize)-2, color.RGBA{100, 200, 125, 255}, false)
	}

	for _, apple := range b.apples {
		vector.DrawFilledRect(screen, float32(apple.X*b.CellSize)+1, float32(apple.Y*b.CellSize)+1+float32(headerSize), float32(b.CellSize)-2, float32(b.CellSize)-2, color.RGBA{250, 75, 125, 255}, false)
	}
}

func (b *Board) DrawCells(screen *ebiten.Image) {
	for x := range b.W {
		for y := range b.H {
			vector.DrawFilledRect(screen, float32(x*b.CellSize)+1, float32(y*b.CellSize)+1+float32(headerSize), float32(b.CellSize)-2, float32(b.CellSize)-2, cellColor, false)
		}
	}
}

func (b *Board) addApple() {
	randomX := rand.Intn(b.W)
	randomY := rand.Intn(b.H)

	b.apples = append(b.apples, Pos{X: randomX, Y: randomY})
}

func (b *Board) CheckSnake(pos *Pos) bool {
	for i, cell := range b.snake {
		if i == len(b.snake)-1 {
			continue
		}

		if pos.X == cell.X && pos.Y == cell.Y {
			return true
		}
	}

	return false
}

func (b *Board) CheckApple(pos *Pos) int {
	for i, apple := range b.apples {
		if apple.X != pos.X || apple.Y != pos.Y {
			continue
		}

		b.apples = append(b.apples[:i], b.apples[i+1:]...)
		b.addApple()

		return i
	}

	return -1
}

func (b *Board) MoveApple(index int) {
	b.apples = append(b.apples[:index], b.apples[index+1:]...)
	b.addApple()
}

func (b *Board) MoveSnake(dirX, dirY int) *Pos {
	last := b.snake[len(b.snake)-1]
	newX, newY := last.X+dirX, last.Y+dirY
	newPos := Pos{X: newX, Y: newY}

	b.snake = append(b.snake, newPos)

	return &newPos
}
