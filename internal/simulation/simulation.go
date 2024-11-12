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

// RunSimulation starts the simulation with user interaction.
func RunSimulation(particles []*particle.Particle) {
    renderer.InitWindow()
    defer renderer.CloseWindow()

    lastTime := time.Now()
    paused := false

    for !rl.WindowShouldClose() {
        currentTime := time.Now()
        dt := currentTime.Sub(lastTime).Seconds()
        lastTime = currentTime

        // Handle user input (pause, add/remove particles)
        HandleUserInput(&particles, &paused)

        if !paused {
            // Update physics if not paused
            for _, p := range particles {
                // Apply gravity to each particle
                // physics.ApplyGravity(p)

                // Apply electrostatic forces between particles
                physics.ApplyElectrostaticForces(particles)

                // Update velocity and position based on applied forces
                physics.UpdateVelocity(p, dt)
                physics.UpdatePosition(p, dt)
            }

            // Check for collisions between particles
            for i := 0; i < len(particles); i++ {
                for j := i + 1; j < len(particles); j++ {
                    if collisions.WillCollide(particles[i], particles[j], dt) {
                        collisions.HandleCollision(particles[i], particles[j])
                    }
                }
            }
        }

        // Render particles and UI
        rl.BeginDrawing()
        rl.ClearBackground(rl.Black)

        // Draw particles
        for _, p := range particles {
            renderer.DrawParticle(p)
        }

        // Draw UI and particle info
        renderer.DrawUI(particles, paused)
        renderer.DrawParticleInfo(particles)

        rl.EndDrawing()

        // Sleep to simulate 120 FPS
        time.Sleep(time.Millisecond * 8)
    }
}

