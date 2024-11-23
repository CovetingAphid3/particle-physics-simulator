package electrostatics

import (
	"math"
	"particle-physics-simulator/internal/constants"
	"particle-physics-simulator/internal/particle"
)

// CalculateElectrostaticForce calculates the magnitude of the electrostatic force between two charged particles.
func CalculateElectrostaticForce(p1, p2 *particle.Particle) float64 {
	if p1.Charge == 0 || p2.Charge == 0 {
		return 0
	}

	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	distSq := dx*dx + dy*dy

	minDistSq := math.Pow(p1.Radius+p2.Radius, 2) 
	if distSq < minDistSq {
		distSq = minDistSq
	}

	return constants.CoulombsConstant * p1.Charge * p2.Charge / distSq
}

// CalculateElectrostaticForceVector calculates the components of the electrostatic force between two particles.
func CalculateElectrostaticForceVector(p1, p2 *particle.Particle) (fx, fy float64) {
	if p1.Charge == 0 || p2.Charge == 0 {
		return 0, 0
	}

	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	distSq := dx*dx + dy*dy

	minDistSq := math.Pow(p1.Radius+p2.Radius, 2) // Dynamic threshold
	if distSq < minDistSq {
		distSq = minDistSq
	}

	forceMag := constants.CoulombsConstant * p1.Charge * p2.Charge / distSq

	// Calculate normalized direction and multiply by force magnitude
	invDist := 1.0 / math.Sqrt(distSq)
	fx = forceMag * dx * invDist
	fy = forceMag * dy * invDist

	return fx, fy
}

// BatchCalculateElectrostaticForces calculates electrostatic forces for all particles in the system.
func BatchCalculateElectrostaticForces(particles []*particle.Particle, forceX, forceY []float64) {
	n := len(particles)

	for i := 0; i < n-1; i++ {
		p1 := particles[i]
		if p1.Charge == 0 || !p1.Movable {
			continue
		}

		for j := i + 1; j < n; j++ {
			p2 := particles[j]
			if p2.Charge == 0 {
				continue
			}

			fx, fy := CalculateElectrostaticForceVector(p1, p2)

			// Apply Newton's third law
			if p1.Movable {
				forceX[i] += fx
				forceY[i] += fy
			}
			if p2.Movable {
				forceX[j] -= fx
				forceY[j] -= fy
			}
		}
	}
}

