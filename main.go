package main

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Pong")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyEscape)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)

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

		// if window gets resized move right paddle since its at edge of screen not on coord 0
		if rl.IsWindowResized() {
			player2.rect.X = float32(rl.GetScreenWidth()) - playerWidth
		}

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

		if gamePaused {
			pauseMenu()
		}

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

func pauseMenu() {
	w := float32(rl.GetScreenWidth())
	h := float32(rl.GetScreenHeight())

	centerX := w * 0.5

	// ---- TITLE ----
	title := "PONG"
	titleSize := int32(72)
	titleY := h * 0.25

	titleX := float32(centerText(title, titleSize))
	rl.DrawText(title, int32(titleX), int32(titleY), titleSize, rl.White)

	// ---- BUTTONS ----
	btnW := float32(240)
	btnH := float32(56)
	btnX := centerX - btnW*0.5

	startY := titleY + float32(titleSize) + 40
	gap := float32(16)

	resumeRect := rl.Rectangle{X: btnX, Y: startY, Width: btnW, Height: btnH}
	restartRect := rl.Rectangle{X: btnX, Y: startY + btnH + gap, Width: btnW, Height: btnH}
	quitRect := rl.Rectangle{X: btnX, Y: startY + (btnH+gap)*2, Width: btnW, Height: btnH}

	// draw buttons (hook logic yourself)
	rg.Button(resumeRect, "RESUME")
	rg.Button(restartRect, "RESTART")
	rg.Button(quitRect, "QUIT")
}

func centerText(text string, fontSize int32) int32 {
	return int32(rl.GetScreenWidth()/2) - rl.MeasureText(text, fontSize)/2
}

func drawButton(text string, fontSize int32, spacing int, padding float32) {
	textWidth := float32(rl.MeasureText(text, fontSize) / 2)
	rect := rl.Rectangle{
		X:      float32(getHCenter(text, fontSize)) - float32(fontSize)*padding*2,
		Y:      float32(rl.GetScreenHeight()/10*spacing) - float32(fontSize)*padding,
		Width:  textWidth*2 + textWidth*2*padding,
		Height: float32(fontSize) + float32(fontSize)*padding*2,
	}
	rl.DrawRectangleLinesEx(rect, 1, rl.White)
	drawText(text, fontSize, spacing)
}

func drawText(text string, fontSize int32, spacing int) {
	xCord := getHCenter(text, fontSize)
	rl.DrawText(text, xCord, int32(rl.GetScreenHeight()/10*spacing), fontSize, rl.White)
}

func getHCenter(txt string, fontSize int32) int32 {
	return int32(rl.GetScreenWidth()/2) - rl.MeasureText(txt, fontSize)/2
}
