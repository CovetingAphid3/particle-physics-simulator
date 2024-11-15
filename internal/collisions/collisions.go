package collisions

import (
	// "log"
	"math"
	"particle-physics-simulator/internal/particle"
)

// CheckCollision checks if two particles are colliding based on their current positions and radii.
func CheckCollision(p1, p2 *particle.Particle) bool {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	distanceSq := dx*dx + dy*dy
	radiusSum := p1.Radius + p2.Radius
	if distanceSq < radiusSum*radiusSum {
		return true
	}
	return false
}

// WillCollide checks if two particles will collide based on their velocities and predicted positions.
func WillCollide(p1, p2 *particle.Particle, dt float64) bool {
	// Predict the future positions of both particles
	p1NextX := p1.X + p1.Vx*dt
	p1NextY := p1.Y + p1.Vy*dt
	p2NextX := p2.X + p2.Vx*dt
	p2NextY := p2.Y + p2.Vy*dt

	// Calculate the squared distance between the predicted positions
	dx := p1NextX - p2NextX
	dy := p1NextY - p2NextY
	distSq := dx*dx + dy*dy

	// Check if the distance is less than the sum of their radii squared (i.e., a collision)
	if distSq < (p1.Radius+p2.Radius)*(p1.Radius+p2.Radius) {
		return true
	}

	return false
}

// HandleCollision handles the actual collision between two particles.
func HandleCollision(p1, p2 *particle.Particle) {
	// Calculate the distance squared between the particles
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	distSq := dx*dx + dy*dy

	// If particles are exactly at the same position, skip collision handling
	if distSq == 0 {
		// log.Println("Particles are exactly overlapping, skipping collision handling.")
		return
	}

	// Normalize the collision direction (unit vector)
	invDist := 1.0 / math.Sqrt(distSq)
	nx := dx * invDist
	ny := dy * invDist

	// Relative velocity between particles
	vx := p1.Vx - p2.Vx
	vy := p1.Vy - p2.Vy

	// Dot product of relative velocity and collision normal
	dotProduct := vx*nx + vy*ny

	// If particles are moving apart, no collision is needed
	if dotProduct >= 0 {
		// log.Println("Particles are moving apart, no collision handled.")
		return
	}

	// Coefficient of Restitution for elastic collision
	impulse := 2 * dotProduct / (p1.Mass + p2.Mass)

	// Handle immovable particle (p2)
	if p1.Movable && !p2.Movable {
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
		// log.Printf("Updated velocities: p1(%.2f, %.2f), p2(%.2f, %.2f)", p1.Vx, p1.Vy, p2.Vx, p2.Vy)
		return
	}

	// Handle collision between two movable particles
	if p1.Movable && p2.Movable {
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
		p2.Vx += impulse * p1.Mass * nx
		p2.Vy += impulse * p1.Mass * ny
		// log.Printf("Updated velocities: p1(%.2f, %.2f), p2(%.2f, %.2f)", p1.Vx, p1.Vy, p2.Vx, p2.Vy)
	}
}

