package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Pong")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyEscape)

	// player rect
	playerWidth := float32(25)
	playerHeight := float32(100)
	//playerHeight := float32(rl.GetScreenHeight())
	player := newPlayer(rl.Rectangle{
		X:      0,
		Y:      (float32(rl.GetScreenHeight()) - playerHeight) / 2,
		Width:  playerWidth,
		Height: playerHeight,
	}, WASD)
	player2 := newPlayer(rl.Rectangle{
		X:      float32(rl.GetScreenWidth()) - playerWidth,
		Y:      (float32(rl.GetScreenHeight()) - playerHeight) / 2,
		Width:  playerWidth,
		Height: playerHeight,
	}, ARROWS)

	ball := newBall()

	for !rl.WindowShouldClose() {
		// update
		player.update()
		player2.update()
		ball.update()

		resolveCollisions(player, player2, ball)

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player.draw()
		player2.draw()
		ball.draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func resolveCollisions(p1 *Player, p2 *Player, b *Ball) {
	if rl.CheckCollisionCircleRec(b.coords, b.radius, p1.rect) { // colliding to left paddle
		rl.DrawText("Colliding p1", 100, 100, 50, rl.Red)
		b.coords.X = p1.rect.X + p1.rect.Width + b.radius
		b.direction.X = -b.direction.X
	} else if rl.CheckCollisionCircleRec(b.coords, b.radius, p2.rect) { // colliding to right paddle
		rl.DrawText("Colliding p2", 100, 100, 50, rl.Red)
		b.coords.X = p2.rect.X - b.radius
		b.direction.X = -b.direction.X
	}
}
