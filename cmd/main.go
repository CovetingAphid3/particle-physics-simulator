package main

import (
	// "particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/simulation"
)

func main() {
	// Create particles with initial positions on opposite sides of the screen
	color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}
	color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}

	// Particle 1, starting on the left side, moving right
	p1 := particle.NewParticle(100, 600, 0, 500.0, 0.0, 0.0, 0.0, 0.0, 0.0, 100.0, 10, color1)

	// Particle 2, starting on the right side, moving left
	p2 := particle.NewParticle(700, 600, 0, -500.0, 0.0, 0.0, 0.0, 0.0, 0.0, 100.0, 10, color2)

	// Optional: Apply a force in the X direction to ensure they cross paths
	// f1 := force.NewForce(500.0, 0.0, -100.0)
	// f2 := force.NewForce(-50.0, 0.0, 0.0)
	// p1.ApplyForce(f1)
	// p2.ApplyForce(f2)

	// Add particles to the simulation
	particles := []*particle.Particle{p1, p2}

	// Run the simulation with these two particles
	simulation.RunSimulation(particles)
}
