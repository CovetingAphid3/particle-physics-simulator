package force

import (
    "particle-physics-simulator/internal/particle"
    "testing"
    "math"
)

const epsilon = 1e-10 // Small value for floating point comparisons

func approxEqual(a, b float64) bool {
    return math.Abs(a-b) < epsilon
}

func TestNewForce(t *testing.T) {
    tests := []struct {
        name               string
        value             float64
        xComponent        float64
        yComponent        float64
        zComponent        float64
        expectedMagnitude float64
    }{
        {
            name:               "Unit force along x-axis",
            value:             1.0,
            xComponent:        1.0,
            yComponent:        0.0,
            zComponent:        0.0,
            expectedMagnitude: 1.0,
        },
        {
            name:               "Force with equal components",
            value:             1.0,
            xComponent:        1.0,
            yComponent:        1.0,
            zComponent:        1.0,
            expectedMagnitude: math.Sqrt(3),
        },
        {
            name:               "Zero force",
            value:             0.0,
            xComponent:        0.0,
            yComponent:        0.0,
            zComponent:        0.0,
            expectedMagnitude: 0.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := NewForce(tt.value, tt.xComponent, tt.yComponent, tt.zComponent)
            if f == nil {
                t.Fatal("NewForce returned nil")
            }
            if !approxEqual(f.Value, tt.value) {
                t.Errorf("Value = %v, want %v", f.Value, tt.value)
            }
            if !approxEqual(f.Magnitude(), tt.expectedMagnitude) {
                t.Errorf("Magnitude = %v, want %v", f.Magnitude(), tt.expectedMagnitude)
            }
        })
    }
}

func TestForceDirection(t *testing.T) {
    tests := []struct {
        name     string
        force    *Force
        expectedX float64
        expectedY float64
        expectedZ float64
    }{
        {
            name:      "Unit vector along x",
            force:     NewForce(1.0, 1.0, 0.0, 0.0),
            expectedX: 1.0,
            expectedY: 0.0,
            expectedZ: 0.0,
        },
        {
            name:      "45 degrees in xy-plane",
            force:     NewForce(1.0, 1.0, 1.0, 0.0),
            expectedX: 1.0/math.Sqrt(2),
            expectedY: 1.0/math.Sqrt(2),
            expectedZ: 0.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x, y, z := tt.force.Direction()
            if !approxEqual(x, tt.expectedX) || !approxEqual(y, tt.expectedY) || !approxEqual(z, tt.expectedZ) {
                t.Errorf("Direction() = (%v, %v, %v), want (%v, %v, %v)",
                    x, y, z, tt.expectedX, tt.expectedY, tt.expectedZ)
            }
        })
    }
}

func TestApplyForce(t *testing.T) {
    tests := []struct {
        name           string
        particle      *particle.Particle
        force         *Force
        expectedVx    float64
        expectedVy    float64
    }{
        {
            name: "Unit force on unit mass",
            particle: &particle.Particle{
                Mass: 1.0,
                Vx:   0.0,
                Vy:   0.0,
            },
            force:      NewForce(1.0, 1.0, 0.0, 0.0),
            expectedVx: 1.0,
            expectedVy: 0.0,
        },
        {
            name: "Force on larger mass",
            particle: &particle.Particle{
                Mass: 2.0,
                Vx:   0.0,
                Vy:   0.0,
            },
            force:      NewForce(2.0, 1.0, 1.0, 0.0),
            expectedVx: 1.0,
            expectedVy: 1.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ApplyForce(tt.particle, tt.force)
            if !approxEqual(tt.particle.Vx, tt.expectedVx) || !approxEqual(tt.particle.Vy, tt.expectedVy) {
                t.Errorf("Velocity = (%v, %v), want (%v, %v)",
                    tt.particle.Vx, tt.particle.Vy, tt.expectedVx, tt.expectedVy)
            }
        })
    }
}


// func TestApplyForces(t *testing.T) {
//     tests := []struct {
//         name      string
//         particles []*particle.Particle
//         checkFn   func([]*particle.Particle) bool
//     }{
//         {
//             name: "Two particles with opposite charges",
//             particles: []*particle.Particle{
//                 {
//                     X: 0, Y: 0, Z: 0,
//                     Mass: 1.0,
//                     Charge: 1.0,
//                     Vx: 0, Vy: 0, Vz: 0,
//                 },
//                 {
//                     X: 1, Y: 0, Z: 0,
//                     Mass: 1.0,
//                     Charge: -1.0,
//                     Vx: 0, Vy: 0, Vz: 0,
//                 },
//             },
//             checkFn: func(particles []*particle.Particle) bool {
//                 // First particle should move right (positive x) and second particle should move left (negative x)
//                 // Store initial velocities
//                 initialVx0 := particles[0].Vx
//                 initialVx1 := particles[1].Vx
//                 
//                 // Apply forces
//                 ApplyForces(particles)
//                 
//                 // Check if velocities changed in the expected directions
//                 velocityChange0 := particles[0].Vx - initialVx0
//                 velocityChange1 := particles[1].Vx - initialVx1
//                 
//                 // Debug output
//                 t.Logf("Particle 0 velocity change: %v", velocityChange0)
//                 t.Logf("Particle 1 velocity change: %v", velocityChange1)
//                 
//                 // For opposite charges, particles should attract
//                 return velocityChange0 > 0 && velocityChange1 < 0
//             },
//         },
//         {
//             name: "Two particles with same charge",
//             particles: []*particle.Particle{
//                 {
//                     X: 0, Y: 0, Z: 0,
//                     Mass: 1.0,
//                     Charge: 1.0,
//                     Vx: 0, Vy: 0, Vz: 0,
//                 },
//                 {
//                     X: 1, Y: 0, Z: 0,
//                     Mass: 1.0,
//                     Charge: 1.0,
//                     Vx: 0, Vy: 0, Vz: 0,
//                 },
//             },
//             checkFn: func(particles []*particle.Particle) bool {
//                 // Store initial velocities
//                 initialVx0 := particles[0].Vx
//                 initialVx1 := particles[1].Vx
//                 
//                 // Apply forces
//                 ApplyForces(particles)
//                 
//                 // Check if velocities changed in the expected directions
//                 velocityChange0 := particles[0].Vx - initialVx0
//                 velocityChange1 := particles[1].Vx - initialVx1
//                 
//                 // Debug output
//                 t.Logf("Particle 0 velocity change: %v", velocityChange0)
//                 t.Logf("Particle 1 velocity change: %v", velocityChange1)
//                 
//                 // For like charges, particles should repel
//                 return velocityChange0 < 0 && velocityChange1 > 0
//             },
//         },
//         {
//             name: "Particles with different masses",
//             particles: []*particle.Particle{
//                 {
//                     X: 0, Y: 0, Z: 0,
//                     Mass: 2.0,    // Heavier particle
//                     Charge: 1.0,
//                     Vx: 0, Vy: 0, Vz: 0,
//                 },
//                 {
//                     X: 1, Y: 0, Z: 0,
//                     Mass: 1.0,    // Lighter particle
//                     Charge: -1.0,
//                     Vx: 0, Vy: 0, Vz: 0,
//                 },
//             },
//             checkFn: func(particles []*particle.Particle) bool {
//                 initialVx0 := particles[0].Vx
//                 initialVx1 := particles[1].Vx
//                 
//                 ApplyForces(particles)
//                 
//                 velocityChange0 := particles[0].Vx - initialVx0
//                 velocityChange1 := particles[1].Vx - initialVx1
//                 
//                 t.Logf("Heavy particle velocity change: %v", velocityChange0)
//                 t.Logf("Light particle velocity change: %v", velocityChange1)
//                 
//                 // The lighter particle should experience more velocity change
//                 return math.Abs(velocityChange1) > math.Abs(velocityChange0)
//             },
//         },
//     }
//
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             // Make a deep copy of particles to preserve initial state
//             originalParticles := make([]*particle.Particle, len(tt.particles))
//             for i, p := range tt.particles {
//                 originalParticles[i] = &particle.Particle{
//                     X: p.X, Y: p.Y, Z: p.Z,
//                     Mass: p.Mass,
//                     Charge: p.Charge,
//                     Vx: p.Vx, Vy: p.Vy, Vz: p.Vz,
//                 }
//             }
//             
//             if !tt.checkFn(tt.particles) {
//                 t.Errorf("Particles did not move as expected")
//                 t.Logf("Original particles: %+v", originalParticles)
//                 t.Logf("Final particles: %+v", tt.particles)
//             }
//         })
//     }
// }

func TestGravitationalForce(t *testing.T) {
    tests := []struct {
        name     string
        p1       *particle.Particle
        p2       *particle.Particle
        expected float64
    }{
        {
            name: "Unit masses at unit distance",
            p1: &particle.Particle{
                X: 0, Y: 0, Z: 0,
                Mass: 1.0,
            },
            p2: &particle.Particle{
                X: 1, Y: 0, Z: 0,
                Mass: 1.0,
            },
            expected: GravitationalConstant,
        },
        {
            name: "Double mass at unit distance",
            p1: &particle.Particle{
                X: 0, Y: 0, Z: 0,
                Mass: 2.0,
            },
            p2: &particle.Particle{
                X: 1, Y: 0, Z: 0,
                Mass: 1.0,
            },
            expected: 2 * GravitationalConstant,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            force := CalculateGravitationalForce(tt.p1, tt.p2)
            if !approxEqual(force, tt.expected) {
                t.Errorf("CalculateGravitationalForce() = %v, want %v", force, tt.expected)
            }
        })
    }
}
