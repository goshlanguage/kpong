package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1024)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "kPong")
	rl.SetTargetFPS(60)

	// startmsg := "Press space to serve"

	player1 := &Paddle{screenHeight, 10, 10, 10, 60}
	player2 := &Paddle{screenHeight, screenWidth - 20, screenHeight - 70, 10, 60}
	ball := &Ball{
		0,
		0,
		screenWidth / 2,
		screenHeight / 2,
		5,
		5,
		false,
	}

	keybindings := make(map[int]func())
	keybindings[rl.KeyW] = player1.Up
	keybindings[rl.KeyS] = player1.Down
	keybindings[rl.KeySpace] = ball.Serve

	keybindings2 := make(map[int]func())
	keybindings2[rl.KeyUp] = player2.Up
	keybindings2[rl.KeyDown] = player2.Down
	keybindings2[rl.KeySpace] = ball.Serve

	// Plug in your controller player 1
	controller1 := &Controller{keybindings}
	controller2 := &Controller{keybindings2}

	game := Game{
		Ball:          ball,
		Player1:       player1,
		Player2:       player2,
		Controller1:   controller1,
		Controller2:   controller2,
		ScoreFontSize: 36,
		ScreenHeight:  screenHeight,
		ScreenWidth:   screenWidth,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		game.Render()
		game.Update()

		rl.EndDrawing()
	}
}
