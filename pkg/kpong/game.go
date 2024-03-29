package kpong

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game is the state store for the game of pong
type Game struct {
	Ball                      *Ball
	Conf                      GameConfig
	Controller1, Controller2  *Controller
	K8S                       *K8SClient
	Player1, Player2          *Player
	PodFontSize               int32
	Score1, Score2            int
	ScoreFontSize             int32
	ScreenHeight, ScreenWidth int32
	SFX                       map[int]rl.Sound
}

//GameConfig allows parameters of the game to be easily configured
type GameConfig struct {
	ConnString                string
	Host                      bool
	HostIP                    string
	Kubeconfig                string
	Namespace                 string
	Port                      string
	ScreenHeight, ScreenWidth int32
}

// Init loads in Game assets
func (g *Game) Init() {
	g.SFX = make(map[int]rl.Sound)
	g.SFX[0] = rl.LoadSound("assets/sfx/collide.wav")
	g.SFX[1] = rl.LoadSound("assets/sfx/lose.wav")
	g.SFX[2] = rl.LoadSound("assets/sfx/bounce.wav")
	g.SFX[3] = rl.LoadSound("assets/sfx/start.wav")
}

// Render draws the things to the screen
func (g *Game) Render() {
	// Draw net at half screen
	rl.DrawLineEx(
		rl.NewVector2(float32(g.ScreenWidth)/2, 0),
		rl.NewVector2(float32(g.ScreenWidth)/2, float32(g.ScreenHeight)),
		2,
		rl.NewColor(245, 245, 245, 100),
	)

	// Render scores
	scoreSize1 := rl.MeasureText(string(g.Score1), g.ScoreFontSize)
	rl.DrawText(
		fmt.Sprintf("%d", g.Score1),
		(g.ScreenWidth/4)-(scoreSize1/2),
		g.ScreenHeight/5,
		g.ScoreFontSize,
		rl.RayWhite,
	)

	scoreSize2 := rl.MeasureText(string(g.Score2), g.ScoreFontSize)
	rl.DrawText(
		fmt.Sprintf("%d", g.Score2),
		3*(g.ScreenWidth/4)-(scoreSize2/2),
		g.ScreenHeight/5,
		g.ScoreFontSize,
		rl.RayWhite,
	)

	// draw Pod1 to screen
	podSize1 := rl.MeasureText(string(g.Player1.Pod.Name), g.PodFontSize)
	rl.DrawText(
		g.Player1.Pod.Name,
		(g.ScreenWidth/4)-(podSize1/2),
		g.ScreenHeight-(g.ScreenHeight/5),
		g.PodFontSize,
		rl.RayWhite,
	)

	// draw Pod2 to screen
	podSize2 := rl.MeasureText(string(g.Player2.Pod.Name), g.PodFontSize)
	rl.DrawText(
		g.Player2.Pod.Name,
		3*(g.ScreenWidth/4)-(podSize2/2),
		g.ScreenHeight-(g.ScreenHeight/5),
		g.PodFontSize,
		rl.RayWhite,
	)

	// Render Ball and Paddles last so they are on "top"
	g.Ball.Render()
	g.Player1.Render()
	g.Player2.Render()

	// If the game isn't in play, display the host string so the user can share
	// TODO: have this check for if the user is hosting a match, and only  display when no other user has joined
	if g.Ball.DX == 0 {
		if g.Conf.ConnString != "" {
			listenMsg := fmt.Sprintf("Listening on: %s", g.Conf.ConnString)
			msgSize := rl.MeasureText(listenMsg, 24)

			rl.DrawText(
				listenMsg,
				(g.ScreenWidth/2)-(msgSize/2),
				g.ScreenHeight-(g.ScreenHeight/5),
				24,
				rl.RayWhite,
			)
		}
	}

}

// Update is the game's main update cycle
func (g *Game) Update() {
	g.Controller1.Update()
	g.Controller2.Update()

	g.Ball.Update()

	g.CheckBounds()
	// TODO: The follow collision detection is buggy, and can result in the ball being "stuck" in the paddle
	// Should probably fix this for playability
	g.CheckCollisions()
}

// CheckBounds makes sure the ball isn't out of bounds, and if it is, reacts
// Checks X and Y axis
func (g *Game) CheckBounds() {
	if g.Ball.X > g.ScreenWidth {
		g.Score1++
		g.Ball.DX = 0
		g.Ball.DY = 0
		g.Ball.X = g.ScreenWidth / 2
		g.Ball.Y = g.ScreenHeight / 2
		g.Ball.Served = false

		newPod, err := g.K8S.CyclePod(g.Player2.Pod)
		if err != nil {
			// TODO: maybe display this error in game
			fmt.Printf("Uhoh: %s", err)
		}
		g.Player2.Pod = newPod
		rl.PlaySound(g.SFX[1])
	}

	if g.Ball.X < 0 {
		g.Score2++
		g.Ball.DX = 0
		g.Ball.DY = 0
		g.Ball.X = g.ScreenWidth / 2
		g.Ball.Y = g.ScreenHeight / 2
		g.Ball.Served = false

		newPod, err := g.K8S.CyclePod(g.Player1.Pod)
		if err != nil {
			// TODO: maybe display this error in game
			fmt.Printf("Uhoh: %s", err)
		}
		g.Player1.Pod = newPod
		rl.PlaySound(g.SFX[1])
	}

	if g.Ball.Y < 0 || g.Ball.Y > g.ScreenHeight {
		rl.PlaySound(g.SFX[0])
		g.Ball.DY = -g.Ball.DY
	}
}

// CheckCollisions is responsible for reacting to collisions
func (g *Game) CheckCollisions() {
	// if a ball collides with a paddle, reverse it's DX and keep it from colliding into the paddle
	if rl.CheckCollisionRecs(GetPaddleCollisionRec(g.Player1.Paddle), GetBallCollisionRec(g.Ball)) {
		rl.PlaySound(g.SFX[0])
		g.Ball.DX = -g.Ball.DX
		if rand.Intn(2) == 1 {
			g.Ball.DY = int32(rand.Intn(10))
		} else {
			g.Ball.DY = -int32(rand.Intn(10))
		}
	}

	if rl.CheckCollisionRecs(GetPaddleCollisionRec(g.Player2.Paddle), GetBallCollisionRec(g.Ball)) {
		rl.PlaySound(g.SFX[0])
		g.Ball.DX = -g.Ball.DX
		if rand.Intn(2) == 1 {
			g.Ball.DY = int32(rand.Intn(10))
		} else {
			g.Ball.DY = -int32(rand.Intn(10))
		}
	}
}
