package main

import (
	// "particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/simulation"
)

func main() {
	color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}  // Red
	color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}  // Green
	color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}  // Blue
	color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}  // Yellow
	color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}  // Magenta
	colorObstacle := particle.Color{R: 0.5, G: 0.5, B: 0.5, A: 1} // Gray obstacle

	// Define particles with different initial positions and velocities
	p1 := particle.NewParticle(100, 500, 0, 250, -150, 0.0, 0.0, 0.0, 0.0, 30.0, 25, color1, true)
	p2 := particle.NewParticle(400, 300, 0, -300, 100, 0.0, 0.0, 0.0, 0.0, 20.0, 30, color2, true)
	p3 := particle.NewParticle(300, 400, 0, 150, -250, 0.0, 0.0, 0.0, 0.0, 25.0, 20, color3, true)
	p4 := particle.NewParticle(500, 200, 0, 100, 300, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color4, true)
	p5 := particle.NewParticle(600, 500, 0, -200, -300, 0.0, 0.0, 0.0, 0.0, 10.0, 15, color5, true)

	// Define a stationary obstacle
	obstacle := particle.NewParticle(500, 400, 0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 150.0, 50, colorObstacle, false)

	// Define forces to create dynamic interactions
	f1 := force.NewForce(200, 150, -50, 0)  // Force applied at an angle
	f2 := force.NewForce(-100, -200, 75, 0) // Opposite direction

	// Apply forces to specific particles for added dynamics
	force.ApplyForce(p1, f1)
	force.ApplyForce(p3, f2)

	// Add particles to the simulation
	particles := []*particle.Particle{p1, p2, p3, p4, p5, obstacle}

	// Run the simulation with these particles
	simulation.RunSimulation(particles)
}

// func main() {
//     // Colors
// color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}  // Red
// color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}  // Green
// color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}  // Blue
// color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}  // Yellow
// color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}  // Magenta
// colorObstacle := particle.Color{R: 0.5, G: 0.5, B: 0.5, A: 1} // Gray obstacle
// colorObstacle2 := particle.Color{R: 0.7, G: 0.2, B: 0.5, A: 1} // Gray obstacle
//
// colorObstacle3 := particle.Color{R: 0.8, G: 0.3, B: 0.8, A: 1} // Gray obstacle
// colorObstacle4 := particle.Color{R: 0.1, G: 0.8, B: 0.5, A: 1} // Gray obstacle
// // Define particles with different initial positions, velocities, and charges
// p1 := particle.NewCoulombParticle(100, 500, 0, 25, -15, 0.0, 0.0, 0.0, 0.0, 30.0, 25, color1, 0.010, true)  // Charge: +1
// p2 := particle.NewCoulombParticle(400, 300, 0, -30, 10, 0.0, 0.0, 0.0, 0.0, 20.0, 30, color2, -0.020, true) // Charge: -1
// p3 := particle.NewCoulombParticle(300, 400, 0, 15, -25, 0.0, 0.0, 0.0, 0.0, 25.0, 20, color3, 0.05, true)  // Charge: +0.5
// p4 := particle.NewCoulombParticle(500, 200, 0, 10, 30, 0.0, 0.0, 0.0, 0.0, 15.0, 15, color4, -0.05, true) // Charge: -0.5
// p5 := particle.NewCoulombParticle(600, 500, 0, -20, -30, 0.0, 0.0, 0.0, 0.0, 10.0, 15, color5, 0.02, true)  // Charge: +0.2
//
// // Adjusted obstacle positions and charges
// o1 := particle.NewCoulombParticle(650, 650, 0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 150.0, 50, colorObstacle, 0.063, false)  // + charge
// o2 := particle.NewCoulombParticle(400, 400, 0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 150.0, 50, colorObstacle2, -0.063, false) // - charge
//
// o3 := particle.NewCoulombParticle(1320, 500, 0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 150.0, 50, colorObstacle3, 0.163, false)  // + charge
// o4 := particle.NewCoulombParticle(1200, 500, 0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 150.0, 50, colorObstacle4, -0.163, false) // - charge
// // Add the Coulomb particles and the obstacles to the list
// particles := []*particle.Particle{p1, p2, p3, p4, p5, o1, o2,o3,o4}
//
// // Run the simulation with these particles
// simulation.RunSimulation(particles)
//
// }
//
