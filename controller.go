package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Controller should provide a way to bind actions to different keys
// It stores a paddle so it can self reference
type Controller struct {
	Keybindings map[int]func()
}

// Update resolvings any key pressed and updates what its controlling
func (c *Controller) Update() {
	for key, function := range c.Keybindings {
		if rl.IsKeyDown(int32(key)) {
			function()
		}
	}
}
