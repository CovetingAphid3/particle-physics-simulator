package simulation

import (
	"particle-physics-simulator/internal/collisions"
	"particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"particle-physics-simulator/internal/renderer"
	"sync"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	TimeStep       = 1.0 / 120.0 // Target simulation time step (120 FPS)
	MagneticFieldX = 0.1         // Example magnetic field in the X direction
	MagneticFieldY = 0.0         // Example magnetic field in the Y direction
)

// RunSimulation runs the main simulation loop
func RunSimulation(particles []*particle.Particle) {
	// Initialize rendering window
	renderer.InitWindow()
	defer renderer.CloseWindow()

	paused := false
	var wg sync.WaitGroup

	for !rl.WindowShouldClose() {
		currentTime := time.Now()

		// Handle user input (pause/unpause, add/remove particles)
		HandleUserInput(&particles, &paused)

		if !paused {
			// Apply global forces (electrostatic and magnetic) to all particles
			physics.ApplyElectrostaticForces(particles)
			magneticField := force.MagneticField{
				Strength:  1.0, // Set the magnitude of the B field
				Direction: 1,   // +1 for out of the plane, -1 for into the plane
			}
			physics.ApplyMagneticForces(particles, magneticField)

			// Parallelize particle updates
			wg.Add(len(particles))
			for _, p := range particles {
				go func(p *particle.Particle) {
					defer wg.Done()
					// Update each particle's velocity and position
					physics.UpdateVelocity(p, TimeStep)
					physics.UpdatePosition(p, TimeStep)
				}(p)
			}
			wg.Wait()

			// Check and handle collisions between particles in parallel
			wg.Add(len(particles) * (len(particles) - 1) / 2) // Maximum number of collisions
			for i := 0; i < len(particles); i++ {
				for j := i + 1; j < len(particles); j++ {
					go func(i, j int) {
						defer wg.Done()
						if collisions.WillCollide(particles[i], particles[j], TimeStep) {
							collisions.HandleCollision(particles[i], particles[j])
						}
					}(i, j)
				}
			}
			wg.Wait()
		}

		// Render the simulation
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Batch render particles
		// renderer.DrawParticle(particles)
		for _, p := range particles {
			renderer.DrawParticle(p)
		}

		// Draw UI (status, FPS, particle count, etc.)
		renderer.DrawUI(particles, paused)
		renderer.DrawParticleInfo(particles)

		rl.EndDrawing()

		// Sleep to maintain a consistent frame rate (120 FPS)
		sleepDuration := time.Second/120 - time.Since(currentTime)
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}
