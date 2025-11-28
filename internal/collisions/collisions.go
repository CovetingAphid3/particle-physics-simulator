package collisions
import (
	"math"
	"particle-physics-simulator/internal/particle"
)

// CheckCollision checks if two particles are overlapping.
func CheckCollision(p1, p2 *particle.Particle) bool {
	if p1.Shape == particle.ShapeCircle && p2.Shape == particle.ShapeCircle {
		dx := p1.X - p2.X
		dy := p1.Y - p2.Y
		distanceSq := dx*dx + dy*dy
		radiusSum := p1.Radius + p2.Radius
		return distanceSq < radiusSum*radiusSum
	} else if (p1.Shape == particle.ShapeRectangle || p1.Shape == particle.ShapeBumper) && 
              (p2.Shape == particle.ShapeRectangle || p2.Shape == particle.ShapeBumper) {
		return p1.X < p2.X+p2.Width &&
			p1.X+p1.Width > p2.X &&
			p1.Y < p2.Y+p2.Height &&
			p1.Y+p1.Height > p2.Y
	} else {
		// Circle-Rectangle/Bumper
		var circle, rect *particle.Particle
		if p1.Shape == particle.ShapeCircle {
			circle, rect = p1, p2
		} else {
			circle, rect = p2, p1
		}
		return checkCircleRectCollision(circle, rect)
	}
}

// WillCollide checks if two particles will collide based on their velocities and predicted positions.
func WillCollide(p1, p2 *particle.Particle, dt float64) bool {
	// Simple prediction: move both and check overlap
	p1NextX := p1.X + p1.Vx*dt
	p1NextY := p1.Y + p1.Vy*dt
	p2NextX := p2.X + p2.Vx*dt
	p2NextY := p2.Y + p2.Vy*dt

	// Temporarily update positions to check collision
	origX1, origY1 := p1.X, p1.Y
	origX2, origY2 := p2.X, p2.Y

	p1.X, p1.Y = p1NextX, p1NextY
	p2.X, p2.Y = p2NextX, p2NextY

	collision := CheckCollision(p1, p2)

	// Restore positions
	p1.X, p1.Y = origX1, origY1
	p2.X, p2.Y = origX2, origY2

	return collision
}

// HandleCollision handles the actual collision between two particles.
func HandleCollision(p1, p2 *particle.Particle) {
	var nx, ny float64
	var dist float64

	if p1.Shape == particle.ShapeCircle && p2.Shape == particle.ShapeCircle {
		dx := p1.X - p2.X
		dy := p1.Y - p2.Y
		distSq := dx*dx + dy*dy
		if distSq == 0 {
			return
		}
		dist = math.Sqrt(distSq)
		nx = dx / dist
		ny = dy / dist
	} else if (p1.Shape == particle.ShapeRectangle || p1.Shape == particle.ShapeBumper) && 
              (p2.Shape == particle.ShapeRectangle || p2.Shape == particle.ShapeBumper) {
		// Rect-Rect collision ignored for now
		return 
	} else {
		// Circle-Rectangle/Bumper
		var circle, rect *particle.Particle
		swapped := false
		if p1.Shape == particle.ShapeCircle {
			circle, rect = p1, p2
		} else {
			circle, rect = p2, p1
			swapped = true
		}

		nx, ny, dist = getCircleRectCollisionNormal(circle, rect)
		if swapped {
			nx = -nx
			ny = -ny
		}
	}

	// Relative velocity
	vx := p1.Vx - p2.Vx
	vy := p1.Vy - p2.Vy

	// Dot product
	dotProduct := vx*nx + vy*ny

	// If moving apart
	if dotProduct >= 0 {
		return
	}

	// Coefficient of Restitution
	restitution := math.Max(p1.Bounciness, p2.Bounciness)

	// Impulse
	// impulse := (1 + e) * dotProduct / (1/m1 + 1/m2)
	// If one is infinite mass (static), 1/m = 0.
	
	invMass1 := 1.0 / p1.Mass
	if !p1.Movable { invMass1 = 0 }
	invMass2 := 1.0 / p2.Mass
	if !p2.Movable { invMass2 = 0 }

	impulse := -(1 + restitution) * dotProduct / (invMass1 + invMass2)

	if p1.Movable {
		p1.Vx += impulse * invMass1 * nx
		p1.Vy += impulse * invMass1 * ny
	}
	if p2.Movable {
		p2.Vx -= impulse * invMass2 * nx
		p2.Vy -= impulse * invMass2 * ny
	}
}

func checkCircleRectCollision(circle, rect *particle.Particle) bool {
	closestX := math.Max(rect.X, math.Min(circle.X, rect.X+rect.Width))
	closestY := math.Max(rect.Y, math.Min(circle.Y, rect.Y+rect.Height))

	dx := circle.X - closestX
	dy := circle.Y - closestY

	return (dx*dx + dy*dy) < (circle.Radius * circle.Radius)
}

func getCircleRectCollisionNormal(circle, rect *particle.Particle) (float64, float64, float64) {
	closestX := math.Max(rect.X, math.Min(circle.X, rect.X+rect.Width))
	closestY := math.Max(rect.Y, math.Min(circle.Y, rect.Y+rect.Height))

	dx := circle.X - closestX
	dy := circle.Y - closestY
	dist := math.Sqrt(dx*dx + dy*dy)

	if dist == 0 {
		return 0, 0, 0 // Should not happen if checked properly
	}

	return dx / dist, dy / dist, dist
}

