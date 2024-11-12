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
    groundLevel   float64 = 580
    dampingFactor float64 = 0.9
)

func InitWindow() {
    rl.InitWindow(int32(screenWidth), int32(screenHeight), "Particle Physics Simulator")
    rl.SetTargetFPS(120)
}

func DrawParticle(p *particle.Particle) {
    drawParticleCircle(p)
    drawParticleInfo(p)
    physics.ApplyBoundaryConditions(p, screenWidth, screenHeight)
}

func drawParticleCircle(p *particle.Particle) {
    color := rl.Color{
        R: uint8(p.Color.R * 255),
        G: uint8(p.Color.G * 255),
        B: uint8(p.Color.B * 255),
        A: uint8(p.Color.A * 255),
    }
    rl.DrawCircle(int32(p.X), int32(p.Y), float32(p.Radius), color)
}

func drawParticleInfo(p *particle.Particle) {
    infoX := int32(p.X) + int32(p.Radius) + 5
    infoY := int32(p.Y) + int32(p.Radius) - 10
    rl.DrawText(p.GetInfo(), infoX, infoY, 10, rl.DarkGray)
}

func CloseWindow() {
    rl.CloseWindow()
}

