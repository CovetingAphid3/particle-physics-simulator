package electrostatics

import (
    "particle-physics-simulator/internal/constants"
    "particle-physics-simulator/internal/particle"
    "testing"
    "math"
)

const epsilon = 1e-10 // Small value for floating point comparisons

func TestCalculateElectrostaticForce(t *testing.T) {
    tests := []struct {
        name     string
        p1       *particle.Particle
        p2       *particle.Particle
        expected float64
    }{
        {
            name: "Equal positive charges at unit distance",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: 1.0,
            },
            p2: &particle.Particle{
                X: 1, Y: 0,
                Charge: 1.0,
            },
            expected: constants.CoulombsConstant, // k*1*1/1²
        },
        {
            name: "Equal negative charges at unit distance",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: -1.0,
            },
            p2: &particle.Particle{
                X: 1, Y: 0,
                Charge: -1.0,
            },
            expected: constants.CoulombsConstant, // k*(-1)*(-1)/1²
        },
        {
            name: "Opposite charges at unit distance",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: 1.0,
            },
            p2: &particle.Particle{
                X: 1, Y: 0,
                Charge: -1.0,
            },
            expected: -constants.CoulombsConstant, // k*1*(-1)/1²
        },
        {
            name: "Particles too close (less than 1e-9)",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: 1.0,
            },
            p2: &particle.Particle{
                X: 1e-10, Y: 0,
                Charge: 1.0,
            },
            expected: 0, // Should return 0 for very close particles
        },
        {
            name: "Verify inverse square law at 2 units",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: 1.0,
            },
            p2: &particle.Particle{
                X: 2, Y: 0,
                Charge: 1.0,
            },
            expected: constants.CoulombsConstant / 4, // k*1*1/2²
        },
        {
            name: "Diagonal separation (Pythagorean)",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: 1.0,
            },
            p2: &particle.Particle{
                X: 3, Y: 4,
                Charge: 1.0,
            },
            expected: constants.CoulombsConstant / 25, // k*1*1/5² (distance = 5 units)
        },
        {
            name: "Zero charge particles",
            p1: &particle.Particle{
                X: 0, Y: 0,
                Charge: 0,
            },
            p2: &particle.Particle{
                X: 1, Y: 0,
                Charge: 1.0,
            },
            expected: 0, // No force between uncharged particle
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            force := CalculateElectrostaticForce(tt.p1, tt.p2)
            if math.Abs(force-tt.expected) > epsilon {
                t.Errorf("CalculateElectrostaticForce() = %v, want %v", force, tt.expected)
            }
        })
    }
}

// TestSymmetry verifies Newton's third law - forces should be equal and opposite
func TestSymmetry(t *testing.T) {
    p1 := &particle.Particle{
        X: 0, Y: 0,
        Charge: 2.0,
    }
    p2 := &particle.Particle{
        X: 1, Y: 1,
        Charge: -3.0,
    }

    force12 := CalculateElectrostaticForce(p1, p2)
    force21 := CalculateElectrostaticForce(p2, p1)

    if math.Abs(force12-force21) > epsilon {
        t.Errorf("Forces not equal: F12=%v, F21=%v", force12, force21)
    }
}
