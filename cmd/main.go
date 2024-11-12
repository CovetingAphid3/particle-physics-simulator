package main

import (
    // "particle-physics-simulator/internal/force"
    "particle-physics-simulator/internal/particle"
    "particle-physics-simulator/internal/simulation"
)

// func main() {
//     color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}  // Red
//     color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}  // Green
//     color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}  // Blue
//     color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}  // Yellow
//     color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}  // Magenta
//
//     // Define particles with different initial positions and velocities
//     p1 := particle.NewParticle(100, 500, 0, 0, 0, 0.0, 0.0, 0.0, 0.0, 20.0, 20, color1)
//     p2 := particle.NewParticle(400, 300, 0, -50.0, 50.0, 0.0, 0.0, 0.0, 0.0, 25.0, 25, color2)
//     p3 := particle.NewParticle(300, 400, 0, 50.0, -50.0, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color3)
//     p4 := particle.NewParticle(500, 200, 0, 0.0, 100.0, 0.0, 0.0, 0.0, 0.0, 10.0, 10, color4)
//     p5 := particle.NewParticle(600, 500, 0, -100.0, -50.0, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color5)
//
//     // Define and apply some forces for demonstration
//     f1 := force.NewForce(100, 200, -50, 0)  // Force applied in an upward-right direction
//     f2 := force.NewForce(200, -100, 100, 0) // Force applied in an upward-left direction
//
//     // Apply forces to specific particles
//     force.ApplyForce(p1, f1)
//     force.ApplyForce(p3, f2)
//
//     // Add particles to the simulation
//     particles := []*particle.Particle{p1, p2, p3, p4, p5}
//
//     // Run the simulation with these particles
//     simulation.RunSimulation(particles)
// }

// func main() {
//     // Define colors for particles
//     color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}  // Red
//     color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}  // Green
//     color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}  // Blue
//     color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}  // Yellow
//     color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}  // Magenta
//
//     // Define particles with different initial positions, velocities, and charges
//     p1 := particle.NewParticle(100, 500, 0, 0, 0, 0.0, 0.0, 0.0, 0.0, 20.0, 20, color1)
//     p2 := particle.NewParticle(400, 300, 0, -50.0, 50.0, 0.0, 0.0, 0.0, 0.0, 25.0, 25, color2)
//     p3 := particle.NewParticle(300, 400, 0, 50.0, -50.0, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color3)
//     p4 := particle.NewParticle(500, 200, 0, 0.0, 100.0, 0.0, 0.0, 0.0, 0.0, 10.0, 10, color4)
//     p5 := particle.NewParticle(600, 500, 0, -100.0, -50.0, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color5)
//
//     // Assign charges to particles (positive or negative)
//     p1.Charge = 1.0   // Positive charge
//     p2.Charge = -1.0  // Negative charge
//     p3.Charge = 1.0   // Positive charge
//     p4.Charge = -1.0  // Negative charge
//     p5.Charge = 1.0   // Positive charge
//
//     // Add particles to the simulation
//     particles := []*particle.Particle{p1, p2, p3, p4, p5}
//
//     // Run the simulation with these particles
//     simulation.RunSimulation(particles)
// }
const GravitationalConstant = 6.67430e-11 // Gravitational constant (m^3 kg^−1 s^−2)

func main() {
    // Define colors for particles
    color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}  // Red
    color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}  // Green
    color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}  // Blue
    color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}  // Yellow
    color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}  // Magenta

    // Define particles with different initial positions, velocities, and charges
    p1 := particle.NewParticle(100, 500, 0, 0, 0, 0.0, 0.0, 0.0, 0.0, 20.0, 20, color1)
    p2 := particle.NewParticle(400, 300, 0, -50.0, 50.0, 0.0, 0.0, 0.0, 0.0, 25.0, 25, color2)
    p3 := particle.NewParticle(300, 400, 0, 50.0, -50.0, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color3)
    p4 := particle.NewParticle(500, 200, 0, 0.0, 100.0, 0.0, 0.0, 0.0, 0.0, 10.0, 10, color4)
    p5 := particle.NewParticle(600, 500, 0, -100.0, -50.0, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color5)

    // Assign masses to particles (in kilograms)
    p1.Mass = 5.972e24  // Earth mass (kg)
    p2.Mass = 4.867e24  // Venus mass (kg)
    p3.Mass = 6.39e23   // Mars mass (kg)
    p4.Mass = 1.989e30  // Sun mass (kg)
    p5.Mass = 1.898e27  // Jupiter mass (kg)

    // Add particles to the simulation
    particles := []*particle.Particle{p1, p2, p3, p4, p5}

    // Run the simulation with these particles
    simulation.RunSimulation(particles)
}
