package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MovementKeys int

const (
	WASD   MovementKeys = 0
	ARROWS MovementKeys = 1
)

type AI struct {
	ball           *Ball
	trackingOffset float32
	followSpeed    float32
}

func newAI(ball *Ball) *AI {
	return &AI{
		ball:           ball,
		trackingOffset: 2,
		followSpeed:    15,
	}
}

func (a *AI) update(p *Player) {
	targetY := a.ball.coords.Y + a.trackingOffset
	dt := rl.GetFrameTime()
	p.rect.Y += (targetY - p.rect.Y) * a.followSpeed * dt
}

type Player struct {
	rect      rl.Rectangle
	speed     int
	direction int
	controls  MovementKeys
	ai        *AI
}

func (p *Player) draw() {
	rl.DrawRectangleRec(p.rect, rl.White)
}

func (p *Player) update() {
	if p.ai == nil {
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
	} else { // if AI exists let it deal handle the movement
		p.ai.update(p)
	}

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

func newPlayer(rect rl.Rectangle, controls MovementKeys, ai *AI) *Player {
	return &Player{
		rect:      rect,
		speed:     400,
		direction: 0,
		controls:  controls,
		ai:        ai,
	}
}
