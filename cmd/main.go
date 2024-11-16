package main

import (
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/simulation"
)

func main() {
	// Define colors for the particles
	color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}  // Red
	color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}  // Green
	color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}  // Blue
	color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}  // Yellow
	color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}  // Magenta
	colorObstacle := particle.Color{R: 0.5, G: 0.5, B: 0.5, A: 1} // Gray obstacle

	// Create a large number of smaller particles
	particles := []*particle.Particle{}
	for i := 0; i < 300; i++ { // Change 1000 to a larger number if needed
		// Random positions and velocities
		x := float64(i % 500) // Distribute particles randomly in x
		y := float64((i * 100) % 500) // Distribute particles randomly in y
		velocityX := math.Sin(float64(i) * 0.1) * 100
		velocityY := math.Cos(float64(i) * 0.1) * 100
		
		// Assign colors in a loop (or make it random)
		var color particle.Color
		switch i % 5 {
		case 0:
			color = color1
		case 1:
			color = color2
		case 2:
			color = color3
		case 3:
			color = color4
		default:
			color = color5
		}

		// Create a new particle with smaller size (e.g., 5 units)
		particle := particle.NewParticle(x, y, velocityX, velocityY,  0.0, 0.0, 5.0, 5, color, true)
		particles = append(particles, particle)
	}

	// Define a stationary obstacle
	obstacle := particle.NewParticle(500, 400,  0.0, 0.0,  0.0, 0.0, 150.0, 50, colorObstacle, false)
	particles = append(particles, obstacle)

	// Define forces to create dynamic interactions

	// Apply forces to specific particles for added dynamics

	// Run the simulation with the large number of particles
	simulation.RunSimulation(particles)
}
