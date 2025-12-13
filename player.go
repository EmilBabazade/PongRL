package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MovementKeys int

const (
	WASD   MovementKeys = 0
	ARROWS MovementKeys = 1
)

type Player struct {
	rect      rl.Rectangle
	speed     int
	direction int
	controls  MovementKeys
}

func (p *Player) draw() {
	rl.DrawRectangleRec(p.rect, rl.White)
}

func (p *Player) update() {
	// set input keys
	var upKey int32 = rl.KeyW
	if p.controls == ARROWS {
		upKey = rl.KeyUp
	}
	var downKey int32 = rl.KeyS
	if p.controls == ARROWS {
		downKey = rl.KeyDown
	}

	// get direction input
	direction := float32(0) // pong paddles only need to move up and down
	if rl.IsKeyDown(upKey) {
		direction = -1
	} else if rl.IsKeyDown(downKey) {
		direction = 1
	} else {
		direction = 0
	}

	// update position
	dt := rl.GetFrameTime()
	p.rect.Y += direction * dt * float32(p.speed)

	// limit to screen bounds
	// loweBound is actually top of the screen and upperBound is actually bottom :D
	lowerBound := float32(0)
	upperBound := float32(rl.GetScreenHeight()) - p.rect.Height
	if p.rect.Y < lowerBound {
		p.rect.Y = lowerBound
	} else if p.rect.Y >= upperBound {
		p.rect.Y = upperBound
	}
}

func newPlayer(rect rl.Rectangle, controls MovementKeys) Player {
	return Player{
		rect:      rect,
		speed:     400,
		direction: 0,
		controls:  controls,
	}
}
