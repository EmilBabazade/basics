package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	position  rl.Vector2
	texture   rl.Texture2D
	width     int32
	height    int32
	direction rl.Vector2
	speed     float32
}

func newPlayer(pos rl.Vector2, texture rl.Texture2D) *Player {
	return &Player{
		position:  pos,
		texture:   texture,
		width:     texture.Width,
		height:    texture.Height,
		direction: rl.Vector2{},
		speed:     400,
	}
}

func (p *Player) draw() {
	rl.DrawTextureV(p.texture, p.position, rl.White)
}

func (p *Player) update() {
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		p.direction.X = -1
	} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		p.direction.X = 1
	} else {
		p.direction.X = 0
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		p.direction.Y = -1
	} else if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		p.direction.Y = 1
	} else {
		p.direction.Y = 0
	}

	if p.position.X <= 0 {
		p.direction.X = 1
	} else if p.position.X >= float32(rl.GetScreenWidth())-float32(p.width) {
		p.direction.X = -1
	}

	if p.position.Y <= 0 {
		p.direction.Y = 1
	} else if p.position.Y >= float32(rl.GetScreenHeight())-float32(p.height) {
		p.direction.Y = -1
	}
	p.direction = rl.Vector2Normalize(p.direction)

	dt := rl.GetFrameTime()
	p.position.X += p.direction.X * p.speed * dt
	p.position.Y += p.direction.Y * p.speed * dt
}
