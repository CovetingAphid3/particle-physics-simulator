package force

import (
	"particle-physics-simulator/internal/particle"
)


func MagneticForce(p *particle.Particle, B float64) (float64, float64) {
	if p.Charge == 0 || (p.Vx == 0 && p.Vy == 0) {
		// No force if the particle is uncharged or stationary
		return 0, 0
	}

	q := p.Charge
	// Fx = q * vy * B, Fy = -q * vx * B
	return q * p.Vy * B, -q * p.Vx * B
}

// MagneticField represents a uniform magnetic field in a 2D simulation.
type MagneticField struct {
	Strength  float64 // Magnitude of the B field
	Direction int     // +1 for out of the plane, -1 for into the plane
}

func MagneticForceWithDirection(p *particle.Particle, field MagneticField) (float64, float64) {
	B := field.Strength * float64(field.Direction)
	return MagneticForce(p, B)
}

// MagneticField2D models a non-uniform magnetic field as a function of position.
type MagneticField2D struct {
	FieldFunc func(x, y float64) (float64, int) // Function to get (magnitude, direction) at a given position
}

// MagneticForceNonUniform calculates the magnetic force for a particle in a non-uniform field.
func MagneticForceNonUniform(p *particle.Particle, field MagneticField2D) (float64, float64) {
	B, direction := field.FieldFunc(p.X, p.Y)
	B *= float64(direction)
	return MagneticForce(p, B)
}

// ApplyMagneticForce applies the magnetic force to a particle by updating its force components.
func ApplyMagneticForce(p *particle.Particle, B float64) {
	fx, fy := MagneticForce(p, B)
	p.Fx += fx
	p.Fy += fy
}

