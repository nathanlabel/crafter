package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nathanlabel/crafter/internal/assets"
)

const (
	GameTitle   = "An Autocrafting Game"
	GameVersion = "0.1.0"
	GameWidth   = 800
	GameHeight  = 600
	TargetFPS   = 60
)

func InitalizeGame() {
	rl.InitWindow(GameWidth, GameHeight, GameTitle)
	rl.SetTargetFPS(TargetFPS)
	assets.InitManager()
}

func main() {
	// Initialize window
	InitalizeGame()

	// Main game loop
	for !rl.WindowShouldClose() {
		// Update

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}

	// Close window and OpenGL context
	rl.CloseWindow()
}
