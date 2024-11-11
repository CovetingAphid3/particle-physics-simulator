// internal/physics/physics.go
package physics

import (
	"math"
	"particle-physics-simulator/internal/particle"
)

const Gravity = 9.8 // Gravitational constant

// Apply gravity force to a particle
func ApplyGravity(p *particle.Particle) {
    p.Ay += Gravity // Apply downward acceleration due to gravity
}

// Update velocity based on acceleration
func UpdateVelocity(p *particle.Particle, dt float64) {
    p.Vx += p.Ax * dt
    p.Vy += p.Ay * dt
    p.Vz += p.Az * dt
}

// Update position based on velocity
func UpdatePosition(p *particle.Particle, dt float64) {
    p.X += p.Vx * dt
    p.Y += p.Vy * dt
    p.Z += p.Vz * dt
}

// internal/physics/physics.go (add to the existing file)
func CheckCollision(p1, p2 *particle.Particle) bool {
    // Calculate distance between the two particles
    dx := p1.X - p2.X
    dy := p1.Y - p2.Y
    dz := p1.Z - p2.Z
    distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

    // If distance is smaller than the sum of radii, a collision happens
    return distance < (p1.Radius + p2.Radius)
}

