package force

import (
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/electrostatics"
	"particle-physics-simulator/internal/constants"
)

// Force struct represents a force with its magnitude and components (in X, Y, Z directions).
type Force struct {
    Value      float64
    XComponent float64
    YComponent float64
    ZComponent float64
}

// NewForce creates a new Force object.
func NewForce(value, xComponent, yComponent, zComponent float64) *Force {
    return &Force{
        Value:      value,
        XComponent: xComponent,
        YComponent: yComponent,
        ZComponent: zComponent,
    }
}

// Magnitude returns the magnitude of the force vector.
func (f *Force) Magnitude() float64 {
    return math.Sqrt(f.XComponent*f.XComponent + f.YComponent*f.YComponent + f.ZComponent*f.ZComponent)
}

// Direction returns the direction of the force as normalized vector components (X, Y, Z).
func (f *Force) Direction() (float64, float64, float64) {
    mag := f.Magnitude()
    return f.XComponent / mag, f.YComponent / mag, f.ZComponent / mag
}

// ApplyForce directly updates the velocity of a particle based on the force components.
func ApplyForce(p *particle.Particle, f *Force) {
    // Apply the force to the particle's velocity (using components)
	p.Vx += (f.Value * f.XComponent) / p.Mass
	p.Vy += (f.Value * f.YComponent) / p.Mass
}

// ApplyForces calculates all forces acting on a particle.
// It includes gravitational, electrostatic, or other forces you may want to add.
func ApplyForces(particles []*particle.Particle) {
	for i := range particles {
        totalForceX := 0.0
        totalForceY := 0.0
        totalForceZ := 0.0

        // Apply gravitational forces (could be implemented separately if needed)
        // Apply electrostatic forces
        for j := range particles {
            if i != j {
                // Calculate electrostatic force between particles[i] and particles[j]
                electrostaticForce := electrostatics.CalculateElectrostaticForce(particles[i], particles[j])

                // Calculate direction of the electrostatic force
                dx := particles[j].X - particles[i].X
                dy := particles[j].Y - particles[i].Y
                dz := particles[j].Z - particles[i].Z
                distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

                // Normalize the direction vector
                normX := dx / distance
                normY := dy / distance
                normZ := dz / distance

                // Apply electrostatic force to X, Y, and Z components
                totalForceX += electrostaticForce * normX
                totalForceY += electrostaticForce * normY
                totalForceZ += electrostaticForce * normZ
            }
        }

        // Apply the resulting force to the particle (change in velocity based on total force)
        particles[i].Fx = totalForceX
        particles[i].Fy = totalForceY
        particles[i].Fz = totalForceZ
    }
}

// CalculateGravitationalForce calculates the gravitational force between two particles.
// If you want to include gravitational forces, you can define it here and include it in ApplyForces.
func CalculateGravitationalForce(p1, p2 *particle.Particle) float64 {
    // Gravitational constant
    G := constants.GravitationalConstant

    // Calculate the distance between the particles
    dx := p2.X - p1.X
    dy := p2.Y - p1.Y
    dz := p2.Z - p1.Z
    distanceSquared := dx*dx + dy*dy + dz*dz

    // Gravitational force: F = G * (m1 * m2) / r²
    forceMagnitude := G * p1.Mass * p2.Mass / distanceSquared

    return forceMagnitude
}

// ApplyGravitationalForce applies the gravitational force between two particles.
func ApplyGravitationalForce(p1, p2 *particle.Particle) *Force {
    // Calculate the force magnitude
    forceMagnitude := CalculateGravitationalForce(p1, p2)

    // Calculate the direction of the gravitational force
    dx := p2.X - p1.X
    dy := p2.Y - p1.Y
    dz := p2.Z - p1.Z
    distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

    // Normalize the direction vector
    normX := dx / distance
    normY := dy / distance
    normZ := dz / distance

    // Return the gravitational force as a new force object
    return NewForce(forceMagnitude, normX, normY, normZ)
}

