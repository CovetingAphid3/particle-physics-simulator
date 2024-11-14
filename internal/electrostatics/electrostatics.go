package electrostatics

import (
    "math"
    "particle-physics-simulator/internal/constants"
    "particle-physics-simulator/internal/particle"
)

// CalculateElectrostaticForce calculates the electrostatic force between two charged particles using Coulomb's law.
func CalculateElectrostaticForce(p1, p2 *particle.Particle) float64 {
    // Ensure the particles have a non-zero distance between them
    dx := p2.X - p1.X
    dy := p2.Y - p1.Y
    // dz := p2.Z - p1.Z
    distanceSquared := dx*dx + dy*dy //+ dz*dz
    distance := math.Sqrt(distanceSquared)

    // Avoid division by zero (if particles are too close)
    if distance < 1e-9 {
        return 0
    }

    // Coulomb's Law: F = k * (q1 * q2) / r²
    forceMagnitude := constants.CoulombsConstant * p1.Charge * p2.Charge / distanceSquared

    // Return the magnitude of the force
    return forceMagnitude
}

