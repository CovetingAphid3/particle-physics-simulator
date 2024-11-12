// internal/physics/physics.go
package physics

import (
	"particle-physics-simulator/internal/particle"
)
const (
    Gravity           = 980.0  // Increased gravity (in pixels/second²)
    dampingFactor     = 0.7    // Reduced damping for more lively bounces
    velocityThreshold = 20.0   // Increased threshold for earlier stopping
    frictionCoef      = 0.0001   // Reduced friction coefficient
    airFrictionCoefficient = 0.05
)

// ApplyGravity force to a particle
func ApplyGravity(p *particle.Particle) {
    // Only apply gravity if particle is not grounded
    if !p.IsGrounded {
        p.Ay = Gravity
    }
}
func ApplyAirFriction(p *particle.Particle) {
    // Apply air friction in the opposite direction of velocity
    p.Vx -= p.Vx * airFrictionCoefficient
    p.Vy -= p.Vy * airFrictionCoefficient
    p.Vz -= p.Vz * airFrictionCoefficient
}


// applyFriction applies friction to a particle on the ground
func applyFriction(p *particle.Particle) {
    if p.Vx > 0 {
        // Apply friction force to the right-moving particle
        p.Vx -= frictionCoef * p.Mass * Gravity / p.Mass
        if p.Vx < 0 {
            p.Vx = 0 // Stop particle if it reverses direction
        }
    } else if p.Vx < 0 {
        // Apply friction force to the left-moving particle
        p.Vx += frictionCoef * p.Mass * Gravity / p.Mass
        if p.Vx > 0 {
            p.Vx = 0 // Stop particle if it reverses direction
        }
    }
}

func UpdateVelocity(p *particle.Particle, dt float64) {
        // ApplyAirFriction(p)
    if !p.IsGrounded {
        p.Vx += p.Ax * dt
        p.Vy += p.Ay * dt
        p.Vz += p.Az * dt
    }
}

// UpdatePosition based on velocity
func UpdatePosition(p *particle.Particle, dt float64) {
    // Scale dt to make physics faster
    scaledDt := dt * 2.0
    p.X += p.Vx * scaledDt
    if !p.IsGrounded {
        p.Y += p.Vy * scaledDt
    }
    p.Z += p.Vz * scaledDt
}

// ApplyBoundaryConditions applies boundary conditions for window edges and ground level
func ApplyBoundaryConditions(p *particle.Particle, screenWidth, screenHeight int) {
    // Right boundary
    if p.X+p.Radius > float64(screenWidth) {
        p.X = float64(screenWidth) - p.Radius
        p.Vx = -p.Vx * dampingFactor
    }
    // Left boundary
    if p.X-p.Radius < 0 {
        p.X = p.Radius
        p.Vx = -p.Vx * dampingFactor
    }

    // Bottom boundary (Ground level)
    groundY := float64(screenHeight - int(p.Radius))  // Ground level accounting for radius

    if p.Y >= groundY {
        if !p.IsGrounded {
            // Bounce
            p.Y = groundY
            p.Vy = -p.Vy * dampingFactor

            // Check if particle should stop
            if abs(p.Vy) < velocityThreshold {
                p.IsGrounded = true
                p.Vy = 0
                p.Ay = 0
                p.Y = groundY  // Lock to ground
            }
        } else {
            // Apply friction when grounded
            applyFriction(p)

            // Keep particle at ground level if grounded
            p.Y = groundY
            p.Vy = 0
            p.Ay = 0
        }
    } else {
        // If particle is above ground level, it's not grounded
        p.IsGrounded = false
    }

    // Top boundary
    if p.Y-p.Radius < 0 {
        p.Y = p.Radius
        p.Vy = -p.Vy * dampingFactor
    }
}


// Helper function to calculate the absolute value of a float
func abs(value float64) float64 {
    if value < 0 {
        return -value
    }
    return value
}
