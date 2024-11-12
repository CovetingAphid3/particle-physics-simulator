package collisions

import (
	"math"
	"particle-physics-simulator/internal/particle"
	"testing"
)

// Helper function to create a particle with given parameters
func NewParticle(x, y, z, vx, vy, vz, mass, radius float64) *particle.Particle {
	return &particle.Particle{
		X:      x,
		Y:      y,
		Z:      z,
		Vx:     vx,
		Vy:     vy,
		Vz:     vz,
		Mass:   mass,
		Radius: radius,
	}
}

func TestCheckCollision(t *testing.T) {
	// Test for collision
	p1 := NewParticle(0, 0, 0, 0, 0, 0, 1, 1)
	p2 := NewParticle(1, 0, 0, 0, 0, 0, 1, 1)
	if !CheckCollision(p1, p2) {
		t.Errorf("Expected collision, but got no collision.")
	}

	// Test for no collision
	p2 = NewParticle(3, 0, 0, 0, 0, 0, 1, 1)
	if CheckCollision(p1, p2) {
		t.Errorf("Expected no collision, but got collision.")
	}
}

func TestWillCollide(t *testing.T) {
	// Test for particles moving towards each other
	p1 := NewParticle(0, 0, 0, 10, 0, 0, 1, 1)
	p2 := NewParticle(15, 0, 0, -10, 0, 0, 1, 1)
	if !WillCollide(p1, p2, 1.0) {
		t.Errorf("Expected particles to collide, but they won't.")
	}

	// Test for particles moving apart
	p1 = NewParticle(0, 0, 0, 10, 0, 0, 1, 1)
	p2 = NewParticle(15, 0, 0, 5, 0, 0, 1, 1)
	if WillCollide(p1, p2, 1.0) {
		t.Errorf("Expected particles not to collide, but they will.")
	}
}

func TestHandleCollision(t *testing.T) {
	// Initialize particles for collision
	p1 := NewParticle(100, 600, 0, 200.0, 0.0, 0.0, 100.0, 10)
	p2 := NewParticle(700, 600, 0, -200.0, 0.0, 0.0, 100.0, 10)

	// Before collision, check velocities
	initialVx1 := p1.Vx
	initialVx2 := p2.Vx

	// Simulate collision
	HandleCollision(p1, p2)

	// Check if the velocities have been updated correctly (elastic collision)
	if math.Abs(p1.Vx+initialVx2) > 1e-6 || math.Abs(p2.Vx+initialVx1) > 1e-6 {
		t.Errorf("Velocities not updated correctly after collision. p1.Vx: %f, p2.Vx: %f", p1.Vx, p2.Vx)
	}
}

func TestHandleCollisionEdgeCase(t *testing.T) {
	// Test for zero distance (particles are already overlapping)
	p1 := NewParticle(0, 0, 0, 0, 0, 0, 1, 1)
	p2 := NewParticle(0, 0, 0, 0, 0, 0, 1, 1)

	// Check if no error occurs during collision handling
	HandleCollision(p1, p2)

	// Since the particles are at the same position, their velocities shouldn't change
	if p1.Vx != 0 || p2.Vx != 0 {
		t.Errorf("Expected velocities to stay the same. p1.Vx: %f, p2.Vx: %f", p1.Vx, p2.Vx)
	}
}

