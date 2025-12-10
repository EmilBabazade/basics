package main

import (
	"log"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var screenWidth int32 = 1720
	var screenHeight int32 = 880
	rl.InitWindow(screenWidth, screenHeight, "Base")
	rl.SetExitKey(rl.KeyEscape)

	ship, err := loadTexture(path.Join("assets", "spaceship.png"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	pos := rl.Vector2{X: 0, Y: 0}
	var speed float32 = 500.0

	for !rl.WindowShouldClose() {
		// input
		direction := getDirection()

		// update
		dt := rl.GetFrameTime()

		if pos.X <= 0 {
			direction.X = 1
		} else if pos.X >= float32(screenWidth)-float32(ship.Width) {
			direction.X = -1
		}

		if pos.Y <= 0 {
			direction.Y = 1
		} else if pos.Y >= float32(screenHeight)-float32(ship.Height) {
			direction.Y = -1
		}
		direction = rl.Vector2Normalize(direction)

		pos.X += direction.X * speed * dt
		pos.Y += direction.Y * speed * dt

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTextureV(ship, pos, rl.White)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func getDirection() rl.Vector2 {
	direction := rl.Vector2{}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		direction.X = -1
	} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		direction.X = 1
	} else {
		direction.X = 0
	}

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		direction.Y = -1
	} else if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		direction.Y = 1
	} else {
		direction.Y = 0
	}

	return direction
}
