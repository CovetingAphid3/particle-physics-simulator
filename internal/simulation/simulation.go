package simulation

import (
	"particle-physics-simulator/internal/collisions"
	// "particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"particle-physics-simulator/internal/renderer"
	"sync"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	TimeStep       = 1.0 / 120.0 // Target simulation time step (120 FPS)
	MagneticFieldX = 0.1      
	MagneticFieldY = 0.0     
)

func RunSimulation(particles []*particle.Particle) {
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
			// physics.ApplyElectrostaticForces(particles)
			// magneticField := force.MagneticField{
			// 	Strength:  1.0, 
			// 	Direction: 1,  
			// }
			// physics.ApplyMagneticForces(particles, magneticField)

			// Parallelize particle updates
			wg.Add(len(particles))
			for _, p := range particles {
				go func(p *particle.Particle) {
					defer wg.Done()
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

		for _, p := range particles {
			renderer.DrawParticle(p)
		}

		renderer.DrawUI(particles, paused)
        renderer.DrawWindowButtons()
		renderer.DrawParticleInfo(particles)

		rl.EndDrawing()

		// Sleep to maintain a consistent frame rate (120 FPS)
		sleepDuration := time.Second/120 - time.Since(currentTime)
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}
