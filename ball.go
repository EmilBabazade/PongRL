package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	radius       float32
	coords       rl.Vector2
	speed        float32
	direction    rl.Vector2
	scoreManager *ScoreManager
}

func (b *Ball) update() {
	dt := rl.GetFrameTime()
	direction := rl.Vector2Normalize(b.direction)
	b.coords.X += direction.X * b.speed * dt
	b.coords.Y += direction.Y * b.speed * dt

	// change direction when colliding with upper and lower bounds of the screen
	lowerBound := float32(0) + b.radius
	upperBound := float32(rl.GetScreenHeight()) - b.radius
	if b.coords.Y <= lowerBound {
		b.direction.Y = 1
	} else if b.coords.Y >= upperBound {
		b.direction.Y = -1
	}

	// reset when leaving screen bounds horizontally ( that's player 1 or 2 scoring )
	if b.coords.X <= 0 {
		b.reset()
		b.scoreManager.p1Score()
	} else if b.coords.X >= float32(rl.GetScreenWidth()) {
		b.reset()
		b.scoreManager.p2Score()
	}
}

func (b *Ball) reset() {
	b.coords = rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight() / 2)}
	b.direction = rl.Vector2{X: getRandFloat(-1, 1), Y: getRandFloat(-1, 1)}
}

func (b *Ball) draw() {
	rl.DrawCircleV(b.coords, b.radius, rl.White)
}

func newBall(scoreManager *ScoreManager) *Ball {
	coords := rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight() / 2)}
	direction := rl.Vector2{X: getRandFloat(-1, 1), Y: getRandFloat(-1, 1)}
	return &Ball{
		radius:       10,
		coords:       coords,
		speed:        float32(400),
		direction:    direction,
		scoreManager: scoreManager,
	}
}
