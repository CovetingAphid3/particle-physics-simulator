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

func HandleCollision(p1, p2 *particle.Particle) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	distance := math.Sqrt(dx*dx + dy*dy)

	// Normal vector components
	nx := dx / distance
	ny := dy / distance

	// Relative velocity
	dvx := p1.Vx - p2.Vx
	dvy := p1.Vy - p2.Vy

	// Dot product of relative velocity and normal vector
	dot := dvx*nx + dvy*ny

	// Calculate the impulse scalar for elastic collision
	impulse := 2 * dot / (p1.Mass + p2.Mass)

	// Update velocities based on the impulse and normal vector
	p1.Vx -= impulse * p2.Mass * nx
	p1.Vy -= impulse * p2.Mass * ny
	p2.Vx += impulse * p1.Mass * nx
	p2.Vy += impulse * p1.Mass * ny
}

func ApplyBoundryConditions(p *particle.Particle, screenWidth, screenHeight int) {
	// Right boundary
	if p.X+p.Radius > float64(screenWidth) {
		p.X = float64(screenWidth) - p.Radius // Reposition particle at the boundary
		p.Vx = -p.Vx                          // Reverse horizontal velocity
	}

	// Left boundary
	if p.X-p.Radius < 0 {
		p.X = p.Radius // Reposition particle at the boundary
		p.Vx = -p.Vx   // Reverse horizontal velocity
	}

	// Bottom boundary
	if p.Y+p.Radius > float64(screenHeight) {
		p.Y = float64(screenHeight) - p.Radius // Reposition particle at the boundary
		p.Vy = -p.Vy                           // Reverse vertical velocity
	}

	// Top boundary
	if p.Y-p.Radius < 0 {
		p.Y = p.Radius // Reposition particle at the boundary
		p.Vy = -p.Vy   // Reverse vertical velocity
	}
}
