// internal/renderer/renderer.go
package renderer

import (
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  int = 800
	screenHeight int = 600
)

// Initialize the window
func InitWindow() {
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Particle Physics Simulator")
	rl.SetTargetFPS(60)
}

// Draw particles on the screen
func DrawParticle(p *particle.Particle) {
	rl.DrawCircle(int32(p.X), int32(p.Y), float32(p.Radius), rl.Color{R: 255, G: 0, B: 0, A: 255})
	physics.ApplyBoundryConditions(p, screenWidth, screenHeight)
}

// Close the window
func CloseWindow() {
	rl.CloseWindow()
}
