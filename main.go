package main

import (
	"fmt"
	"math"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	winScl = 1000
	radius = winScl / 2
)

var (
	cornerVertices    = 6
	intersectingEdges = 1
)

var corners []rl.Vector2 = make([]rl.Vector2, cornerVertices)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.InitWindow(winScl, winScl, "Polygon Star")
	rl.SetTargetFPS(60)

	generatePoints()
	for !rl.WindowShouldClose() {

		update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func update() {
	if rl.IsKeyPressed(rl.KeyW) {
		cornerVertices += 1
		corners = make([]rl.Vector2, cornerVertices)
		generatePoints()
	}
	if rl.IsKeyPressed(rl.KeyQ) && cornerVertices > 3 {
		cornerVertices -= 1
		corners = make([]rl.Vector2, cornerVertices)
		generatePoints()
	}
	if rl.IsKeyPressed(rl.KeyS) {
		intersectingEdges += 1
	}
	if rl.IsKeyPressed(rl.KeyA) && intersectingEdges > 1 {
		intersectingEdges -= 1
	}
	if cornerVertices%2 == 0 && intersectingEdges > (cornerVertices/2)-1 {
		intersectingEdges = (cornerVertices / 2) - 1
	}
	if cornerVertices%2 != 0 && intersectingEdges > (cornerVertices/2) {
		intersectingEdges = (cornerVertices / 2)
	}
}

func generatePoints() {
	deg := 360.0 / float32(cornerVertices)
	degAfter := float32(0.0)
	for i := 0; i < cornerVertices; i++ {
		angle := ((math.Pi * float64(degAfter)) / 180.0)
		x := winScl/2 + (radius * float32(math.Sin(angle)))
		y := winScl/2 + (radius * float32(math.Cos(angle)))

		corners[i] = rl.NewVector2(x, y)

		degAfter += deg
	}
}

func draw() {
	nextI := 0
	for i := range corners {
		nextI = i + intersectingEdges
		if nextI >= cornerVertices {
			extra := nextI - cornerVertices
			nextI = extra
		}
		rl.DrawLineEx(corners[i], corners[nextI], 1, rl.Black)
	}
	rl.DrawText(fmt.Sprintf("{%d/%d}", cornerVertices, intersectingEdges), 20, 20, 20, rl.Red)
}
