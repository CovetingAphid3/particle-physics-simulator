package collisions
import (
	"math"
	"particle-physics-simulator/internal/particle"
)

func CheckCollision(p1, p2 *particle.Particle) bool {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	// Using squared distance to avoid unnecessary sqrt calculation
	distanceSq := dx*dx + dy*dy
	radiusSum := p1.Radius + p2.Radius
	// Compare squared values to avoid calculating the square root
	return distanceSq < radiusSum*radiusSum
}

// WillCollide checks if two particles will collide based on their velocities and predicted positions.
func WillCollide(p1, p2 *particle.Particle, dt float64) bool {
	// Predict the future positions of both particles (no need to recalculate velocity here)
	p1NextX := p1.X + p1.Vx*dt
	p1NextY := p1.Y + p1.Vy*dt
	p2NextX := p2.X + p2.Vx*dt
	p2NextY := p2.Y + p2.Vy*dt

	// Calculate the squared distance between the predicted positions
	dx := p1NextX - p2NextX
	dy := p1NextY - p2NextY
	distSq := dx*dx + dy*dy

	return distSq < (p1.Radius+p2.Radius)*(p1.Radius+p2.Radius)
}

// HandleCollision handles the actual collision between two particles.
func HandleCollision(p1, p2 *particle.Particle) {
	// Calculate the distance squared between the particles once
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	distSq := dx*dx + dy*dy

	// Early exit if particles are exactly at the same position (no collision handling)
	if distSq == 0 {
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
		return
	}

	// Coefficient of Restitution for elastic collision
	impulse := 2 * dotProduct / (p1.Mass + p2.Mass)

	// Handle immovable particle (p2)
	if p1.Movable && !p2.Movable {
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
		return
	}

	// Handle collision between two movable particles
	if p1.Movable && p2.Movable {
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
		p2.Vx += impulse * p1.Mass * nx
		p2.Vy += impulse * p1.Mass * ny
	}
}

