package physics

import (
    "math"
    "particle-physics-simulator/internal/particle"
)

const (
    Gravity           = 980.0  // Increased gravity (in pixels/second²)
    dampingFactor     = 0.7    // Reduced damping for more lively bounces
    velocityThreshold = 20.0   // Increased threshold for earlier stopping
)


// ApplyGravity force to a particle
func ApplyGravity(p *particle.Particle) {
    // Only apply gravity if particle is not grounded
    if !p.IsGrounded {
        p.Ay = Gravity
    }
}

func UpdateVelocity(p *particle.Particle, dt float64) {
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
func CheckCollision(p1, p2 *particle.Particle) bool {
	// Calculate distance between the two particles
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

	// If distance is smaller than the sum of radii, a collision happens
	return distance < (p1.Radius + p2.Radius)
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
