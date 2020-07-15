package kpong

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Ball represents the ball object
type Ball struct {
	Color         rl.Color
	DX, DY        int32
	X, Y          int32
	Width, Height int32
	Served        bool
}

// Update keeps track of the ball
func (b *Ball) Update() {
	b.X += b.DX
	b.Y += b.DY
}

// Render draws the ball to the screen
func (b *Ball) Render() {
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, b.Color)
}

// Serve sets the ball in motion
func (b *Ball) Serve() {
	if !b.Served {
		die := rand.Intn(2)
		b.DX = 8
		if die == 1 {
			b.DX = -b.DX
		}
		b.DY = int32(rand.Intn(10))
		b.Served = true
	}
}

// GetBallCollisionRec returns a rectangle used to detect ball collision
func GetBallCollisionRec(b *Ball) rl.Rectangle {
	return rl.NewRectangle(
		float32(b.X),
		float32(b.Y),
		float32(b.Width),
		float32(b.Height),
	)
}
