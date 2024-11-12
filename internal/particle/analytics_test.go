// internal/particle/analytics_test.go
package particle

import "testing"

func TestGetInfo(t *testing.T) {
    tests := []struct {
        name      string
        particle  *Particle
        expected string
    }{
        {
            name: "Normal particle",
            particle: NewParticle(1.0, 2.0, 3.0, 0.5, 0.5, 0.5, 0.0, 0.0, 0.0, 10.0, 1.0, Color{R: 1.0, G: 0.0, B: 0.0, A: 1.0}),
            expected: "Mass: 10.00, Position: (1.000000, 2.000000), Velocity: (0.500000, 0.500000), Grounded: false",
        },
        {
            name: "Zero mass and velocity",
            particle: NewParticle(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.5, Color{R: 0.0, G: 1.0, B: 0.0, A: 1.0}),
            expected: "Mass: 0.00, Position: (0.000000, 0.000000), Velocity: (0.000000, 0.000000), Grounded: false",
        },
        {
            name: "Grounded particle",
            particle: NewParticle(5.0, 5.0, 5.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 20.0, 1.0, Color{R: 0.0, G: 0.0, B: 1.0, A: 1.0}),
            expected: "Mass: 20.00, Position: (5.000000, 5.000000), Velocity: (1.000000, 1.000000), Grounded: false",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.particle.GetInfo()
            if got != tt.expected {
                t.Errorf("GetInfo() = %v, want %v", got, tt.expected)
            }
        })
    }
}

