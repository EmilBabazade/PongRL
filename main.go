package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Pong")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyEscape)

	// player rect
	playerWidth := float32(25)
	playerHeight := float32(100)
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

	for !rl.WindowShouldClose() {
		// update
		player.update()
		player2.update()

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player.draw()
		player2.draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
