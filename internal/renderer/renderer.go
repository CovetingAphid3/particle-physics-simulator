// internal/renderer/renderer.go
package renderer

import (
    "github.com/gen2brain/raylib-go/raylib"
    "particle-physics-simulator/internal/particle"
)

// Initialize the window
func InitWindow() {
    rl.InitWindow(800, 600, "Particle Physics Simulator")
    rl.SetTargetFPS(60)
}

// Draw particles on the screen
func DrawParticle(p *particle.Particle) {
    rl.DrawCircle(int32(p.X), int32(p.Y), float32(p.Radius), rl.Color{R: 255, G: 0, B: 0, A: 255})
}

// Close the window
func CloseWindow() {
    rl.CloseWindow()
}

