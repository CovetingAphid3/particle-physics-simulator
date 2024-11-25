package main

import (
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/simulation"
)

func main() {
	color1 := particle.Color{R: 1, G: 0, B: 0, A: 1} 
	color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}
	color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}
	color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}
	color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}

	particles := []*particle.Particle{}
	for i := 0; i < 100; i++ { 
		x := float64(i % 500)         
		y := float64((i * 100) % 500) 
		velocityX := math.Sin(float64(i)*0.1) * 100
		velocityY := math.Cos(float64(i)*0.1) * 100
		charge := float64(math.Cos(float64(i)*0.1) * 0.100) 

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

		particle := particle.NewCoulombParticle(x, y, velocityX, velocityY, 0.0, 0.0, 5.0, 5, color, charge, true)
		particles = append(particles, particle)
	}

	magnetic_particles := []*particle.Particle{}
	for i := 0; i < 100; i++ { 
		x := float64(i % 500)
		y := float64((i * 100) % 500)
		velocityX := math.Sin(float64(i)*0.1) * 100
		velocityY := math.Cos(float64(i)*0.1) * 100
		charge := float64(math.Cos(float64(i)*0.1) * 0.100)

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

		particle := particle.NewCoulombParticle(x, y, velocityX, velocityY, 0.0, 0.0, 5.0, 5, color, charge, true)
		magnetic_particles = append(magnetic_particles, particle)
	}

	particles = append(particles, magnetic_particles...)

	obstacle := particle.NewParticle(300, 300, 0.0, 0.0, 0.0, 0.0, 100.0, 50, particle.Color{R: 0.5, G: 0.5, B: 0.5, A: 1}, false) // Obstacle at a fixed location
	particles = append(particles, obstacle)

    // run simulation
	simulation.RunSimulation(particles)
}
