package main

import (
	"fmt"
	"log"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1720, 880, "Base")

	ship, err := loadTexture(path.Join("assets", "spaceship.png"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	pos := rl.Vector2{
		X: 0,
		Y: 0,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		dt := rl.GetFrameTime()

		rl.DrawTextureV(ship, pos, rl.White)
		pos.X = pos.X + 10*dt

		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func loadImage(path string) (*rl.Image, error) {
	image := rl.LoadImage(path)
	if image.Data == nil {
		return nil, fmt.Errorf("file not found at path %s", path)
	}
	return image, nil
}

func loadFont(path string) (rl.Font, error) {
	font := rl.LoadFont(path)
	if font.Chars == nil {
		return rl.Font{}, fmt.Errorf("file not found at path %s", path)
	}
	return font, nil
}

func loadTexture(path string) (rl.Texture2D, error) {
	texture := rl.LoadTexture(path)
	if texture.Width == 0 && texture.Height == 0 {
		return rl.Texture2D{}, fmt.Errorf("file not found at path %s", path)
	}
	return texture, nil
}
