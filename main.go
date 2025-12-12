package main

import (
	"image/color"
	"log"
	"math/rand"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  int32 = 1520
	screenHeight int32 = 680
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Base")
	rl.SetExitKey(rl.KeyEscape)
	rl.SetTargetFPS(60) // comment this out to start the jet engine

	// inputMovement()
	// collisions()
	//cameraStuff()
	//audioStuff()
	classStuff()
}

func classStuff() {
	texture, err := loadTexture(path.Join("assets", "spaceship.png"))
	if err != nil {
		log.Fatal(err)
	}

	player := newPlayer(rl.Vector2{}, texture)

	for !rl.WindowShouldClose() {
		player.update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player.draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func audioStuff() {
	rl.InitAudioDevice()
	laserSound, err := loadSound(path.Join("assets", "laser.wav"))
	if err != nil {
		log.Fatal(err)
	}

	musicStream, err := loadMusicStream(path.Join("assets", "music.wav"))
	if err != nil {
		log.Fatal(err)
	}
	rl.PlayMusicStream(musicStream)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(musicStream)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeySpace) {
			rl.PlaySound(laserSound)
		}

		rl.EndDrawing()
	}
	rl.UnloadMusicStream(musicStream)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func cameraStuff() {
	type Circle struct {
		position rl.Vector2
		radius   float32
		color    color.RGBA
	}
	circles := make([]Circle, 0, 100)
	for range 100 {
		x := getRandInt(-2000, 2000)
		y := getRandInt(-1000, 1000)
		radius := getRandInt(50, 200)
		clr := color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(rand.Intn(255)),
		}
		c := Circle{
			position: rl.Vector2{X: float32(x), Y: float32(y)},
			radius:   float32(radius),
			color:    clr,
		}
		circles = append(circles, c)
	}

	player := Circle{
		position: rl.Vector2{X: 0, Y: 0},
		radius:   50,
		color:    rl.Red,
	}
	playerSpeed := float32(400)

	camera := rl.Camera2D{
		Offset:   rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight / 2)},
		Target:   player.position,
		Rotation: 0,
		Zoom:     1,
	}

	for !rl.WindowShouldClose() {
		// update
		direction := getDirection()
		direction = rl.Vector2Normalize(direction)

		dt := rl.GetFrameTime()
		player.position.X += direction.X * playerSpeed * dt
		player.position.Y += direction.Y * playerSpeed * dt

		camera.Target = player.position

		rotateDirection := 0
		if rl.IsKeyDown(rl.KeyD) {
			rotateDirection = 1
		} else if rl.IsKeyDown(rl.KeyA) {
			rotateDirection = -1
		}
		camera.Rotation += dt * float32(rotateDirection) * 50
		camera.Zoom += dt * rl.GetMouseWheelMove() * 2
		if camera.Zoom <= 0.35 {
			camera.Zoom = 0.35
		} else if camera.Zoom >= 3 {
			camera.Zoom = 3
		}

		// draw
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.White)

		for _, c := range circles {
			rl.DrawCircleV(c.position, c.radius, c.color)
		}
		rl.DrawCircleV(player.position, player.radius, player.color)

		rl.DrawFPS(0, 0)

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func getRandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
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

	//if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
	if rl.IsKeyDown(rl.KeyLeft) {
		direction.X = -1
		//} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
	} else if rl.IsKeyDown(rl.KeyRight) {
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
