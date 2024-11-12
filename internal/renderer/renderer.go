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
    // Convert the particle color from float32 (0.0 to 1.0) to uint8 (0 to 255)
    color := rl.Color{
        R: uint8(p.Color.R * 255),
        G: uint8(p.Color.G * 255),
        B: uint8(p.Color.B * 255),
        A: uint8(p.Color.A * 255),
    }

    // Draw the particle as a circle with the particle's color
    rl.DrawCircle(
        int32(p.X),
        int32(p.Y),
        float32(p.Radius),
        color,
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

