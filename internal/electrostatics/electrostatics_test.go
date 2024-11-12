package electrostatics

import (
    "particle-physics-simulator/internal/particle"
    "testing"
)

func TestCalculateElectrostaticForce(t *testing.T) {
    p1 := &particle.Particle{
        X:      0.0,
        Y:      0.0,
        Charge: 1e-6, // Charge in Coulombs
    }
    p2 := &particle.Particle{
        X:      1.0,
        Y:      0.0,
        Charge: -1e-6, // Charge in Coulombs
    }

    // The expected force is based on Coulomb's law
    expectedForce := 8.9875e9 * 1e-6 * -1e-6 / 1.0 // N·m²/C² * C² / m²

    result := CalculateElectrostaticForce(p1, p2)
    if result != expectedForce {
        t.Errorf("Expected %v, got %v", expectedForce, result)
    }
}

