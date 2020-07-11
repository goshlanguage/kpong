package kpong

import (
	v1 "k8s.io/api/core/v1"
)

// Player models a player
type Player struct {
	Paddle *Paddle
	Pod    v1.Pod
}

// Render calls the player paddle render function to render it to the screen
func (p *Player) Render() {
	p.Paddle.Render()
}

// Update doesn't really do anything, does it?
func (p *Player) Update() {
}
