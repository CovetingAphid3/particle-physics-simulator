package physics

import (
	"math"
	"particle-physics-simulator/internal/constants"
	"particle-physics-simulator/internal/electrostatics"
	"particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
)

func ApplyGravity(p *particle.Particle, gravityStrength float64) {
	if !p.IsGrounded {
		p.Ay = gravityStrength
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

func UpdateVelocity(p *particle.Particle, dt, gravityStrength float64) {
	if p.Movable {
        ApplyGravity(p, gravityStrength)
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
	if p.Shape == particle.ShapeCircle {
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
	} else {
		// Rectangle / Bumper
		// Right boundary
		if p.X+p.Width > float64(screenWidth) {
			p.X = float64(screenWidth) - p.Width
			p.Vx = -p.Vx * constants.DampingFactor
		}

		// Left boundary
		if p.X < 0 {
			p.X = 0
			p.Vx = -p.Vx * constants.DampingFactor
		}

		// Bottom boundary
		groundY := float64(screenHeight) - p.Height
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
		if p.Y < 0 {
			p.Y = 0
			p.Vy = -p.Vy * constants.DampingFactor
		}
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

func ApplySpringForces(springs []*particle.Spring) {
	for _, s := range springs {
		dx := s.P2.X - s.P1.X
		dy := s.P2.Y - s.P1.Y
		dist := math.Sqrt(dx*dx + dy*dy)

		if dist == 0 {
			continue
		}

		// Hooke's Law: F = -k * (currentLength - restLength)
		forceMag := s.Stiffness * (dist - s.RestLength)

		// Damping: Fd = -c * relativeVelocity
		vx := s.P2.Vx - s.P1.Vx
		vy := s.P2.Vy - s.P1.Vy
		
		// Project relative velocity onto the spring axis
		// Axis unit vector: (dx/dist, dy/dist)
		dot := (vx*dx + vy*dy) / dist
		dampingForce := s.Damping * dot

		totalForce := forceMag + dampingForce

		fx := totalForce * (dx / dist)
		fy := totalForce * (dy / dist)

		if s.P1.Movable {
			s.P1.Vx += fx / s.P1.Mass * constants.SecondsPerFrame
			s.P1.Vy += fy / s.P1.Mass * constants.SecondsPerFrame
		}
		if s.P2.Movable {
			s.P2.Vx -= fx / s.P2.Mass * constants.SecondsPerFrame
			s.P2.Vy -= fy / s.P2.Mass * constants.SecondsPerFrame
		}
	}
}
