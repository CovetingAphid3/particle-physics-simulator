// cmd/main.go
package main

import (
	"particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/simulation"
)

func main() {
	// Create some test particles
	color := particle.Color{R: 1, G: 0, B: 0, A: 1}
	p := particle.NewParticle(200, 0.0, 10, 10, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 10, color)

	p2 := particle.NewParticle(0.0, 0.0, 10, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 10, color)

	f := force.NewForce(10.0, 1.0, 0.0)

	// Apply the force to the particle
	p.ApplyForce(f)
	//
	particles := []*particle.Particle{
	    {X: 0, Y: 500, Radius: 10, Mass: 1, Color: particle.Color{R: 1, G: 0, B: 0, A: 1}},
	    {X: 300, Y: 200, Radius: 10, Mass: 1, Color: particle.Color{R: 0, G: 1, B: 0, A: 1}},
	}
	particles2 := []*particle.Particle{p, p2}

	// Run the simulation
	// simulation.RunSimulationSingle(p2)
    simulation.RunSimulation(particles2)
    simulation.RunSimulation(particles)
}
