// internal/simulation/simulation.go
package simulation

import (
	"particle-physics-simulator/internal/collisions"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"particle-physics-simulator/internal/renderer"
	"time"

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

		// Check for collisions between particles using swept collision detection
		for i := 0; i < len(particles); i++ {
			for j := i + 1; j < len(particles); j++ {
				if collisions.WillCollide(particles[i], particles[j], TimeStep) || collisions.CheckCollision(particles[i],particles[j]) {
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

		time.Sleep(time.Millisecond * 16) // Simulate 60 FPS
	}
}

func RunSimulationSingle(p *particle.Particle) {
	RunSimulation([]*particle.Particle{p})
}
