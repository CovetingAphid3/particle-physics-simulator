// cmd/main.go
package main

import (
    "particle-physics-simulator/internal/particle"
    "particle-physics-simulator/internal/simulation"
)

func main() {
    // Create some test particles
    particles := []*particle.Particle{
        {X: 0, Y: 500, Radius: 10, Mass: 1, Color: particle.Color{R: 1, G: 0, B: 0, A: 1}},
        {X: 300, Y: 200, Radius: 10, Mass: 1, Color: particle.Color{R: 0, G: 1, B: 0, A: 1}},
    }

    // Run the simulation
    simulation.RunSimulation(particles)
}

