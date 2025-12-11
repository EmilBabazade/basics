package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  int32 = 1720
	screenHeight int32 = 880
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Base")
	rl.SetExitKey(rl.KeyEscape)
	rl.SetTargetFPS(60) // comment this out to start the jet engine

	// inputMovement()
	// collisions()
	camera()
}

func camera() {
	type Circle struct {
		position rl.Vector2
		radius   float32
		color    color.RGBA
	}
	circles := make([]Circle, 0, 100)
	for range 100 {
		x := rand.Intn(2000-(-2000)+1) + (-2000)
		y := rand.Intn(1000-(-1000)+1) + (-1000)
		radius := rand.Intn(200-50+1) + 50
		color := color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: uint8(rand.Intn(255))}
		c := Circle{
			position: rl.Vector2{X: float32(x), Y: float32(y)},
			radius:   float32(radius),
			color:    color,
		}
		circles = append(circles, c)
	}

	player := Circle{
		position: rl.Vector2{X: 0, Y: 0},
		radius:   50,
		color:    rl.Red,
	}
	playerSpeed := float32(400)

	for !rl.WindowShouldClose() {
		// update
		direction := getDirection()
		direction = rl.Vector2Normalize(direction)

		dt := rl.GetFrameTime()
		player.position.X += direction.X * playerSpeed * float32(dt)
		player.position.Y += direction.Y * playerSpeed * float32(dt)
		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		for _, c := range circles {
			rl.DrawCircleV(c.position, c.radius, c.color)
		}
		rl.DrawCircleV(player.position, player.radius, player.color)

		rl.DrawFPS(0, 0)

		rl.EndDrawing()
	}
}

func collisions() {

	obstaclePos := rl.Vector2{X: 500, Y: 400}
	playerRadius := float32(50)
	obstacleRadius := float32(30)
	rect := rl.NewRectangle(0, 0, 100, 200)
	rect2 := rl.NewRectangle(800, 500, 200, 300)

	for !rl.WindowShouldClose() {
		// update
		playerPos := rl.GetMousePosition()
		rect.X = float32(rl.GetMouseX())
		rect.Y = float32(rl.GetMouseY())
		collisionRect := rl.GetCollisionRec(rect, rect2)

		// draw
		rl.BeginDrawing()

		// rl.DrawCircleV(playerPos, playerRadius, rl.White)
		rl.DrawCircleV(obstaclePos, obstacleRadius, rl.Red)

		rl.DrawRectangleRec(rect, rl.Blue)
		rl.DrawRectangleRec(rect2, rl.Green)

		if rl.CheckCollisionCircles(playerPos, playerRadius, obstaclePos, obstacleRadius) {
			// rl.DrawText("COLLIDING", 200, 300, 100, rl.Red)
		}

		if collisionRect.Width > 0 && collisionRect.Height > 0 {
			rl.DrawRectangleRec(collisionRect, rl.Red)
		}

		// if rl.CheckCollisionCircleRec(playerPos, playerRadius, rect) {
		// 	rl.DrawText("COLLIDING", 200, 300, 100, rl.Blue)
		// }

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}

func inputMovement() {
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
