package force

import (
    "particle-physics-simulator/internal/particle"
)

// MagneticForce calculates the magnetic force on a charged particle moving through a magnetic field.
func MagneticForce(p *particle.Particle, magneticFieldX, magneticFieldY, magneticFieldZ float64) (float64, float64, float64) {
    // The magnetic force is given by F = q * (v × B)
    q := p.Charge
    vx, vy, vz := p.Vx, p.Vy, p.Vz
    
    // Magnetic field vector (Bx, By, Bz)
    Bx, By, Bz := magneticFieldX, magneticFieldY, magneticFieldZ
    
    // Cross product of velocity and magnetic field (v × B)
    forceX := q * (vy*Bz - vz*By)
    forceY := q * (vz*Bx - vx*Bz)
    forceZ := q * (vx*By - vy*Bx)
    
    return forceX, forceY, forceZ
}

