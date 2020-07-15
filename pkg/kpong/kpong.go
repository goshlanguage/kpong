package kpong

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Start initiates the main game loop
func Start(kubeconfig string, namespace string) {
	screenWidth := int32(1024)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "kPong")
	rl.InitAudioDevice()

	rl.SetTargetFPS(60)

	// startmsg := "Press space to serve"

	var kubeErr bool

	clientset, err := newK8SClient(kubeconfig)
	if err != nil {
		kubeErr = true
	}

	pod1, err := GetRandomPod(clientset, namespace)
	if err != nil {
		kubeErr = true
	}
	pod2, err := GetRandomPod(clientset, namespace)
	if err != nil {
		kubeErr = true
	}

	player1 := &Player{
		Paddle: &Paddle{rl.RayWhite, screenHeight, 10, 10, 10, 75},
		Pod:    pod1,
	}
	player2 := &Player{
		Paddle: &Paddle{rl.RayWhite, screenHeight, screenWidth - 20, screenHeight - 70, 10, 75},
		Pod:    pod2,
	}
	ball := &Ball{
		rl.RayWhite,
		0,
		0,
		screenWidth / 2,
		screenHeight / 2,
		10,
		10,
		false,
	}
	if kubeErr {
		// TODO: use a different ball texture to denote that an error has occurred
	}

	keybindings := make(map[int]func())
	keybindings[rl.KeyW] = player1.Paddle.Up
	keybindings[rl.KeyS] = player1.Paddle.Down
	keybindings[rl.KeySpace] = ball.Serve

	keybindings2 := make(map[int]func())
	keybindings2[rl.KeyUp] = player2.Paddle.Up
	keybindings2[rl.KeyDown] = player2.Paddle.Down
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
		KubeClient:    clientset,
		PodFontSize:   18,
		ScoreFontSize: 36,
		ScreenHeight:  screenHeight,
		ScreenWidth:   screenWidth,
	}
	game.Init()
	rl.PlaySound(game.SFX[3])

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		game.Render()
		game.Update()

		rl.EndDrawing()
	}
}
