package simulation

import (
    "particle-physics-simulator/internal/collisions"
    "particle-physics-simulator/internal/particle"
    "particle-physics-simulator/internal/physics"
    "particle-physics-simulator/internal/renderer"
    "time"

    "github.com/gen2brain/raylib-go/raylib"
)

const TimeStep = 1.0 / 120.0 // 120 FPS

// Run the simulation with dynamic frame rate
func RunSimulation(particles []*particle.Particle) {
    renderer.InitWindow()
    defer renderer.CloseWindow()

    lastTime := time.Now()

    for !rl.WindowShouldClose() {
        currentTime := time.Now()
        dt := currentTime.Sub(lastTime).Seconds()  // Calculate time delta (in seconds)
        lastTime = currentTime

        // Update physics
        for _, p := range particles {
            physics.ApplyGravity(p)
            physics.UpdateVelocity(p, dt)
            physics.UpdatePosition(p, dt)
        }

        // Check for collisions with swept collision detection
        for i := 0; i < len(particles); i++ {
            for j := i + 1; j < len(particles); j++ {
                if collisions.WillCollide(particles[i], particles[j], dt) {
                    collisions.HandleCollision(particles[i], particles[j])
                }
            }
        }

        // Render particles
        rl.BeginDrawing()
        rl.ClearBackground(rl.RayWhite)

        for _, p := range particles {
            renderer.DrawParticle(p)
        }

        rl.EndDrawing()

        // Sleep to simulate 120 FPS
        time.Sleep(time.Millisecond * 8) // Simulate 120 FPS
    }
}

func RunSimulationSingle(p *particle.Particle) {
    RunSimulation([]*particle.Particle{p})
}

