package kpong

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func MainMenu(conf GameConfig) {
	rl.InitWindow(conf.ScreenWidth, conf.ScreenHeight, "kPong")
	rl.InitAudioDevice()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		title := "kPong"
		titleWidth := rl.MeasureText(title, 64) / 2
		rl.DrawText(title, (int32(rl.GetScreenWidth())/2)-titleWidth, 100, 64, rl.White)

		startButton := raygui.Button(rl.NewRectangle(float32(conf.ScreenWidth/2)-40, float32(conf.ScreenHeight)-float32(conf.ScreenHeight/5)*1.2, 100, 30), "Start Game")
		if startButton {
			Start(conf)
		}

		hostButton := raygui.Button(rl.NewRectangle(float32(conf.ScreenWidth/2)-40, float32(conf.ScreenHeight)-float32(conf.ScreenHeight/5), 100, 30), "Host Match")
		if hostButton {
			go Listen()
			Start(conf)
		}

		rl.EndDrawing()
	}
}
