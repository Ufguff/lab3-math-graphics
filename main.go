package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ufguff/types"
)

const (
	maxLen = int32(500)
	minLen = int32(400)
)

func main() {
	// white := color.RGBA{0, 0, 0, 0}
	angle := (math.Pi * 3) / 180

	sq := types.NewSquare(types.NewPoint(float64(400), float64(300)), float64(100), rl.Black)

	rl.InitWindow(800, 600, "LAB 3")
	defer rl.CloseWindow()

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		sq.DrawSquare()
		sq.Changes(angle)

		rl.EndDrawing()
	}
}
