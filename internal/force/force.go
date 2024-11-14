package force

import (
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/electrostatics"
	"particle-physics-simulator/internal/constants"
)
const GravitationalConstant = 6.67430e-11 // m^3 kg^−1 s^−2
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
// ApplyForces calculates all forces acting on a particle.
func ApplyForces(particles []*particle.Particle) {
    for i := range particles {
        totalForceX := 0.0
        totalForceY := 0.0
        totalForceZ := 0.0

        for j := range particles {
            if i != j {
                dx := particles[j].X - particles[i].X
                dy := particles[j].Y - particles[i].Y
                dz := particles[j].Z - particles[i].Z
                distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

                if distance < 1e-6 {
                    continue
                }

                electrostaticForce := electrostatics.CalculateElectrostaticForce(particles[i], particles[j])

                // Normalize the direction vector
                normX := dx / distance
                normY := dy / distance
                normZ := dz / distance

                totalForceX += electrostaticForce * normX
                totalForceY += electrostaticForce * normY
                totalForceZ += electrostaticForce * normZ
            }
        }

        // Apply the resulting force to the particle's velocity (change in velocity based on total force)
        particles[i].Vx += totalForceX / particles[i].Mass
        particles[i].Vy += totalForceY / particles[i].Mass
        particles[i].Vz += totalForceZ / particles[i].Mass
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

// ApplyGravitationalForces calculates gravitational forces between all particles
func ApplyGravitationalForces(particles []*particle.Particle) {
    for i := 0; i < len(particles); i++ {
        for j := i + 1; j < len(particles); j++ {
            p1 := particles[i]
            p2 := particles[j]

            // Calculate the distance between the two particles
            dx := p2.X - p1.X
            dy := p2.Y - p1.Y
            dz := p2.Z - p1.Z
            distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

            // Skip if distance is zero (particles can't interact with themselves)
            if distance == 0 {
                continue
            }

            // Calculate the gravitational force magnitude
            forceMagnitude := GravitationalConstant * p1.Mass * p2.Mass / (distance * distance)

            // Calculate the unit vector direction of the force
            forceX := forceMagnitude * dx / distance
            forceY := forceMagnitude * dy / distance
            forceZ := forceMagnitude * dz / distance

            // Apply the force to both particles (action and reaction)
            p1.Ax += forceX / p1.Mass
            p1.Ay += forceY / p1.Mass
            p1.Az += forceZ / p1.Mass

            p2.Ax -= forceX / p2.Mass
            p2.Ay -= forceY / p2.Mass
            p2.Az -= forceZ / p2.Mass
        }
    }
}
