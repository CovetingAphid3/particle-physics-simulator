// internal/simulation/simulation.go
package simulation

import (
    "time"
    "particle-physics-simulator/internal/particle"
    "particle-physics-simulator/internal/physics"
    "particle-physics-simulator/internal/renderer"
    "github.com/gen2brain/raylib-go/raylib"
)

const TimeStep = 1.0 / 60.0 // 60 FPS

// Run the simulation
func RunSimulation(particles []*particle.Particle) {
    renderer.InitWindow()
    defer renderer.CloseWindow()

    for !rl.WindowShouldClose() {
        // Update particle physics
        for _, p := range particles {
            physics.ApplyGravity(p)
            physics.UpdateVelocity(p, TimeStep)
            physics.UpdatePosition(p, TimeStep)
        }

        // Render particles
        rl.BeginDrawing()
        rl.ClearBackground(rl.RayWhite)

        for _, p := range particles {
            renderer.DrawParticle(p)
        }

        rl.EndDrawing()

        time.Sleep(time.Millisecond * 16) // Simulate 60 FPS
    }
}

func RunSimulationSingle(p *particle.Particle) {
    RunSimulation([]*particle.Particle{p})
}
