package physics

import (
	"particle-physics-simulator/internal/constants"
	"particle-physics-simulator/internal/electrostatics"
	"particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
)

func ApplyGravity(p *particle.Particle) {
	if !p.IsGrounded {
		p.Ay = constants.Gravity
	}
}

func ApplyAirFriction(p *particle.Particle) {
	p.Vx -= p.Vx * constants.AirDragCoefficient
	p.Vy -= p.Vy * constants.AirDragCoefficient
}

func ApplyFriction(p *particle.Particle) {
	frictionCoef := constants.GroundFrictionCoefficient
	if p.Movable {
		frictionCoef *= 1.2
	} else {
		frictionCoef *= 0.8
	}

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

func UpdateVelocity(p *particle.Particle, dt float64) {
	if p.Movable {
        ApplyGravity(p)
		p.Vx += p.Ax * dt
		p.Vy += p.Ay * dt
	}
}

func UpdatePosition(p *particle.Particle, dt float64) {
	if p.Movable {
		p.X += p.Vx * dt
		p.Y += p.Vy * dt
	}
}

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

	// Bottom boundary
	groundY := float64(screenHeight) - p.Radius
	if p.Y >= groundY {
		p.Y = groundY
		p.Vy = -p.Vy * constants.DampingFactor

		if abs(p.Vy) < constants.VelocityThreshold {
			p.IsGrounded = true
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

// ApplyMagneticForces applies magnetic forces to particles.
func ApplyMagneticForces(particles []*particle.Particle, magneticField force.MagneticField) {
	for _, p := range particles {
		fx, fy := force.MagneticForceWithDirection(p, magneticField)
		p.Ax += fx / p.Mass
		p.Ay += fy / p.Mass
	}
}

func ApplyElectrostaticForces(particles []*particle.Particle) {
	for i := range particles {
		totalFx, totalFy := 0.0, 0.0

		for j := range particles {
			if i != j {
				fx, fy := electrostatics.CalculateElectrostaticForceVector(particles[i], particles[j])
				totalFx += fx
				totalFy += fy
			}
		}

		particles[i].Fx = totalFx
		particles[i].Fy = totalFy
	}
}

// abs returns the absolute value of a float.
func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

