package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "Pong")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyEscape)

	playerWidth := float32(25)
	playerHeight := float32(100)
	//playerHeight := float32(rl.GetScreenHeight())
	player := newPlayer(rl.Rectangle{
		X:      0,
		Y:      (float32(rl.GetScreenHeight()) - playerHeight) / 2,
		Width:  playerWidth,
		Height: playerHeight,
	}, WASD, nil)

	scoreManager := &ScoreManager{}
	ball := newBall(scoreManager)

	ai := newAI(ball)
	player2 := newPlayer(rl.Rectangle{
		X:      float32(rl.GetScreenWidth()) - playerWidth,
		Y:      (float32(rl.GetScreenHeight()) - playerHeight) / 2,
		Width:  playerWidth,
		Height: playerHeight,
	}, ARROWS, ai)

	gamePaused := false
	for !rl.WindowShouldClose() {
		// update
		player.update()
		player2.update()
		ball.update()

		if rl.IsKeyPressed(rl.KeySpace) {
			if gamePaused {
				resumeAll(player, player2, ball)
				gamePaused = false
			} else {
				pauseAll(player, player2, ball)
				gamePaused = true
			}
		}

		resolveCollisions(player, player2, ball)

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player.draw()
		player2.draw()
		ball.draw()
		scoreText(scoreManager)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func resolveCollisions(p1 *Player, p2 *Player, b *Ball) {
	if rl.CheckCollisionCircleRec(b.coords, b.radius, p1.rect) { // colliding to left paddle
		//rl.DrawText("Colliding p1", 100, 100, 50, rl.Red)
		b.coords.X = p1.rect.X + p1.rect.Width + b.radius
		b.direction.X = -b.direction.X
		b.speed += b.speed * 0.1
	} else if rl.CheckCollisionCircleRec(b.coords, b.radius, p2.rect) { // colliding to right paddle
		//rl.DrawText("Colliding p2", 100, 100, 50, rl.Red)
		b.coords.X = p2.rect.X - b.radius
		b.direction.X = -b.direction.X
		b.speed += b.speed * 0.1
	}
}

func scoreText(scoreManager *ScoreManager) {
	score := fmt.Sprintf("%d:%d", scoreManager.p1, scoreManager.p2)
	textWidth := rl.MeasureText(score, 50)
	xCord := int32(rl.GetScreenWidth()/2 - int(textWidth/2))
	rl.DrawText(score, xCord, 0, 50, rl.White)
}

func pauseAll(items ...Pausable) {
	for _, item := range items {
		item.pause()
	}
}

func resumeAll(items ...Pausable) {
	for _, item := range items {
		item.resume()
	}
}
