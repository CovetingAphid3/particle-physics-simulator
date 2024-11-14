package collisions

import (
	"fmt"
	"log"
	"math"
	"particle-physics-simulator/internal/particle"
)

func CheckCollision(p1, p2 *particle.Particle) bool {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

	if distance < (p1.Radius + p2.Radius) {
		fmt.Printf("Collision detected between particles at distance: %f\n", distance)
		return true
	}
	return false
}

// WillCollide checks if two particles will collide based on their velocities and predicted positions
func WillCollide(p1, p2 *particle.Particle, dt float64) bool {
	// Predict the future positions of both particles
	p1NextX := p1.X + p1.Vx*dt
	p1NextY := p1.Y + p1.Vy*dt
	p2NextX := p2.X + p2.Vx*dt
	p2NextY := p2.Y + p2.Vy*dt

	// Log predicted positions
	log.Printf("Predicted positions: p1(%.2f, %.2f), p2(%.2f, %.2f)", p1NextX, p1NextY, p2NextX, p2NextY)

	// Calculate the squared distance between the predicted positions
	dx := p1NextX - p2NextX
	dy := p1NextY - p2NextY
	distSq := dx*dx + dy*dy

	// Log the distance squared
	log.Printf("Distance squared between predicted positions: %.2f", distSq)

	// Check if the distance is less than the sum of their radii (i.e., a collision)
	if distSq < (p1.Radius+p2.Radius)*(p1.Radius+p2.Radius) {
		log.Println("Collision predicted!")
		return true
	}

	log.Println("No collision predicted.")
	return false
}

// HandleCollision handles the actual collision between two particles
// HandleCollision handles the actual collision between two particles in 2D
func HandleCollision(p1, p2 *particle.Particle) {
	// Ensure p2 is immovable (adjust accordingly based on your particle system)
	if p1.Movable && !p2.Movable {
		// Calculate the distance between the particles in 2D
		dx := p1.X - p2.X
		dy := p1.Y - p2.Y
		distSq := dx*dx + dy*dy

		// Avoid division by zero if particles are exactly at the same position
		if distSq == 0 {
			log.Println("Particles are exactly overlapping, skipping collision handling.")
			return
		}

		// Normalize the collision direction (unit vector)
		invDist := 1.0 / math.Sqrt(distSq)
		nx := dx * invDist
		ny := dy * invDist

		// Relative velocity of the particles
		vx := p1.Vx - p2.Vx
		vy := p1.Vy - p2.Vy

		// Dot product of relative velocity and collision normal
		dotProduct := vx*nx + vy*ny

		// If particles are moving apart, no collision is needed
		if dotProduct >= 0 {
			log.Println("Particles are moving apart, no collision handled.")
			return
		}

		// Coefficient of Restitution for elastic collision
		impulse := 2 * dotProduct / (p1.Mass + p2.Mass)

		// Apply COR for immovable particles (no movement for p2, only p1's velocity changes)
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny

		// Log the updated velocity
		log.Printf("Updated velocities: p1(%.2f, %.2f), p2(%.2f, %.2f)", p1.Vx, p1.Vy, p2.Vx, p2.Vy)
	} else if p1.Movable && p2.Movable {
		// Handle collisions between two movable particles
		// Calculate distance squared between particles
		dx := p1.X - p2.X
		dy := p1.Y - p2.Y
		distSq := dx*dx + dy*dy

		// Avoid division by zero if particles are exactly at the same position
		if distSq == 0 {
			log.Println("Particles are exactly overlapping, skipping collision handling.")
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
			log.Println("Particles are moving apart, no collision handled.")
			return
		}

		// Coefficient of Restitution for elastic collision
		impulse := 2 * dotProduct / (p1.Mass + p2.Mass)

		// Update velocities based on impulse and masses
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
		p2.Vx += impulse * p1.Mass * nx
		p2.Vy += impulse * p1.Mass * ny

		// Log the updated velocity
		log.Printf("Updated velocities: p1(%.2f, %.2f), p2(%.2f, %.2f)", p1.Vx, p1.Vy, p2.Vx, p2.Vy)
	}
}

