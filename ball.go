package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	radius    float32
	coords    rl.Vector2
	speed     float32
	direction rl.Vector2
}

func (b *Ball) update() {
	// TODO move in random dir
	dt := rl.GetFrameTime()
	direction := rl.Vector2Normalize(b.direction)
	b.coords.X += direction.X * b.speed * dt
	b.coords.Y += direction.Y * b.speed * dt

	// TODO change direction when colliding with upper and lower bounds of the screen
}

func (b *Ball) draw() {
	rl.DrawCircleV(b.coords, b.radius, rl.White)
}

func newBall() Ball {
	coords := rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight() / 2)}
	direction := rl.Vector2{X: getRandFloat(-1, 1), Y: getRandFloat(-1, 1)}
	return Ball{
		radius:    10,
		coords:    coords,
		speed:     float32(400),
		direction: direction,
	}
}
