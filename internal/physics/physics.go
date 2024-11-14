package physics

import (
	"math"
	"particle-physics-simulator/internal/electrostatics"
	"particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/constants"
)

// ApplyGravity applies gravitational force to a particle
func ApplyGravity(p *particle.Particle) {
	if !p.IsGrounded {
		p.Ay = constants.Gravity
	}
}

// ApplyAirFriction applies air friction to a particle
func ApplyAirFriction(p *particle.Particle) {
	// Apply air friction in the opposite direction of velocity
	p.Vx -= p.Vx * constants.AirDragCoefficient
	p.Vy -= p.Vy * constants.AirDragCoefficient
	p.Vz -= p.Vz * constants.AirDragCoefficient
}

// applyFriction applies ground friction to a grounded particle
func applyFriction(p *particle.Particle) {
	// Friction coefficient adjusted based on particle mobility
	frictionCoef := constants.GroundFrictionCoefficient
	if p.Movable {
		frictionCoef *= 1.2
	} else {
		frictionCoef *= 0.8
	}

	// Apply friction to velocity (x-direction)
	if p.Vx > 0 {
		p.Vx -= frictionCoef * constants.Gravity
		if p.Vx < 0 {
			p.Vx = 0
		}
	} else if p.Vx < 0 {
		p.Vx += frictionCoef * constants.Gravity
		if p.Vx > 0 {
			p.Vx = 0
		}
	}
}

// UpdateVelocity updates particle velocity based on forces
func UpdateVelocity(p *particle.Particle, dt float64) {
	// Apply gravity if not grounded
	// ApplyGravity(p)

	// Apply air friction if not grounded
	if !p.IsGrounded {
		// ApplyAirFriction(p)
	}

	// Update velocity based on acceleration and time step
	if p.Movable {
		p.Vx += p.Ax * dt
		p.Vy += p.Ay * dt
		p.Vz += p.Az * dt
	}
}

// UpdatePosition updates the particle position based on its velocity
func UpdatePosition(p *particle.Particle, dt float64) {
	if p.Movable {
		p.X += p.Vx * dt
		if !p.IsGrounded {
			p.Y += p.Vy * dt
		}
		p.Z += p.Vz * dt
	}
}

// ApplyBoundaryConditions handles boundary constraints for particles
func ApplyBoundaryConditions(p *particle.Particle, screenWidth, screenHeight int) {
	// Right boundary
	if p.X+p.Radius > float64(screenWidth) {
		p.X = float64(screenWidth) - p.Radius
		p.Vx = -p.Vx * constants.DampingFactor
	}
	// Left boundary
	if p.X-p.Radius < 0 {
		p.X = p.Radius
		p.Vx = -p.Vx * constants.DampingFactor
	}

	// Bottom boundary (Ground level)
	groundY := float64(screenHeight - int(p.Radius))

	if p.Y >= groundY {
		if !p.IsGrounded {
			p.Y = groundY
			p.Vy = -p.Vy * constants.DampingFactor

			// Stop particle if velocity falls below threshold
			if abs(p.Vy) < constants.VelocityThreshold {
				p.IsGrounded = true
				p.Vy = 0
				p.Ay = 0
				p.Y = groundY
			}
		} else {
			// applyFriction(p) // Apply friction when grounded
			p.Y = groundY
			p.Vy = 0
			p.Ay = 0
		}
	} else {
		p.IsGrounded = false
	}

	// Top boundary
	if p.Y-p.Radius < 0 {
		p.Y = p.Radius
		p.Vy = -p.Vy * constants.DampingFactor
	}
}

// ApplyMagneticForces applies magnetic forces to particles
func ApplyMagneticForces(particles []*particle.Particle, magneticFieldX, magneticFieldY, magneticFieldZ float64) {
	for _, p := range particles {
		// Magnetic force can affect all three components of acceleration
		forceX, forceY, forceZ := force.MagneticForce(p, magneticFieldX, magneticFieldY, magneticFieldZ)
		p.Ax += forceX / p.Mass
		p.Ay += forceY / p.Mass
		p.Az += forceZ / p.Mass
	}
}

// ApplyElectrostaticForces calculates and applies electrostatic forces
func ApplyElectrostaticForces(particles []*particle.Particle) {
	for i := range particles {
		totalForceX := 0.0
		totalForceY := 0.0
		totalForceZ := 0.0

		for j := range particles {
			if i != j {
				electrostaticForce := electrostatics.CalculateElectrostaticForce(particles[i], particles[j])

				dx := particles[j].X - particles[i].X
				dy := particles[j].Y - particles[i].Y
				dz := particles[j].Z - particles[i].Z
				distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

				normX := dx / distance
				normY := dy / distance
				normZ := dz / distance

				totalForceX += electrostaticForce * normX
				totalForceY += electrostaticForce * normY
				totalForceZ += electrostaticForce * normZ
			}
		}

		particles[i].Fx = totalForceX
		particles[i].Fy = totalForceY
		particles[i].Fz = totalForceZ
	}
}

// abs returns the absolute value of a float
func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

