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
	} else if p1.Shape == particle.ShapeRectangle && p2.Shape == particle.ShapeRectangle {
		return p1.X < p2.X+p2.Width &&
			p1.X+p1.Width > p2.X &&
			p1.Y < p2.Y+p2.Height &&
			p1.Y+p1.Height > p2.Y
	} else {
		// Circle-Rectangle
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
	// This is a bit expensive but accurate enough for now.
	// We can optimize later if needed.
	
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
	} else if p1.Shape == particle.ShapeRectangle && p2.Shape == particle.ShapeRectangle {
		// Rect-Rect collision normal is tricky, usually along the axis of least penetration
		// For now, let's skip complex rect-rect physics or approximate it.
		// Or just treat them as static walls mostly?
		// Let's implement a simple AABB response if needed, but usually walls don't move.
		// If both are movable rectangles, we need proper logic.
		// Let's assume for now walls (rects) are static.
		return 
	} else {
		// Circle-Rectangle
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

	// Impulse
	impulse := 2 * dotProduct / (p1.Mass + p2.Mass)

	if p1.Movable && !p2.Movable {
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
	} else if !p1.Movable && p2.Movable {
		p2.Vx += impulse * p1.Mass * nx
		p2.Vy += impulse * p1.Mass * ny
	} else if p1.Movable && p2.Movable {
		p1.Vx -= impulse * p2.Mass * nx
		p1.Vy -= impulse * p2.Mass * ny
		p2.Vx += impulse * p1.Mass * nx
		p2.Vy += impulse * p1.Mass * ny
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

