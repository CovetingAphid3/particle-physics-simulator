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
    groundLevel   float64 = 580   // Adjust ground level slightly above screen bottom
    dampingFactor float64 = 0.9   // Damping factor to reduce velocity on each bounce
)

// Initialize the window
func InitWindow() {
    rl.InitWindow(int32(screenWidth), int32(screenHeight), "Particle Physics Simulator")
    rl.SetTargetFPS(60)
}

// Draw particles on the screen
func DrawParticle(p *particle.Particle) {
    // Draw the particle as a red circle
    rl.DrawCircle(
        int32(p.X), 
        int32(p.Y), 
        float32(p.Radius), 
        rl.Color{R: 255, G: 0, B: 0, A: 255},
    )

    // Draw particle info slightly below the particle
    infoX := int32(p.X) + int32(p.Radius) + 5
    infoY := int32(p.Y) + int32(p.Radius) - 10
    rl.DrawText(p.GetInfo(), infoX, infoY, 10, rl.DarkGray)

    // Apply boundary conditions to keep the particle within screen bounds
    physics.ApplyBoundaryConditions(p, screenWidth, screenHeight)
}
// Close the window
func CloseWindow() {
    rl.CloseWindow()
}

