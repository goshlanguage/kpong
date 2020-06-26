package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Paddle represents a pong paddle
type Paddle struct {
	ScreenHeight  int32
	X, Y          int32
	Width, Height int32
}

// Render draws the paddle to the screen
func (p *Paddle) Render() {
	rl.DrawRectangle(p.X, p.Y, p.Width, p.Height, rl.RayWhite)
}

// If you're looking for update, it doesn't exist because controllers update paddles

// Up moves a paddle up
func (p *Paddle) Up() {
	if p.Y > 0 {
		p.Y -= 8
	}
}

// Down moves a paddle down
func (p *Paddle) Down() {
	if p.Y < (p.ScreenHeight - p.Height) {
		p.Y += 8
	}
}

// GetPaddleCollisionRec returns a rectangle used to detect paddle collision with the ball
func GetPaddleCollisionRec(p *Paddle) rl.Rectangle {
	return rl.NewRectangle(
		float32(p.X),
		float32(p.Y),
		float32(p.Width),
		float32(p.Height),
	)
}
