package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ufguff/config"
	"github.com/ufguff/types"
)

func main() {

	angle := (math.Pi * 3) / 180

	sq := types.NewSquare(types.NewPoint(float64(400), float64(300)), float64(100), rl.Black)

	rl.InitWindow(config.W, config.H, "LAB 3")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		sq.DrawSquare()
		sq.Changes(angle)

		rl.EndDrawing()
	}
}
