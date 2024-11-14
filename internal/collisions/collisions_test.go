package collisions

import (
	"testing"
	"particle-physics-simulator/internal/particle"
	"github.com/stretchr/testify/assert"
)

func TestCheckCollision(t *testing.T) {
	// Create two particles that are supposed to collide
	p1 := &particle.Particle{
		X:      0,
		Y:      0,
		Radius: 10,
	}
	p2 := &particle.Particle{
		X:      15,
		Y:      0,
		Radius: 10,
	}

	// Check collision (should be true, particles are touching)
	collisionDetected := CheckCollision(p1, p2)
	assert.True(t, collisionDetected, "Collision should be detected")

	// Move p2 further away and check again (should be false)
	p2.X = 30
	collisionDetected = CheckCollision(p1, p2)
	assert.False(t, collisionDetected, "No collision should be detected")
}

func TestWillCollide(t *testing.T) {
	// Create particles with velocity
	p1 := &particle.Particle{
		X:     0,
		Y:     0,
		Vx:    10,
		Vy:    0,
		Radius: 5,
	}
	p2 := &particle.Particle{
		X:     15,
		Y:     0,
		Vx:    -5,
		Vy:    0,
		Radius: 5,
	}

	// Check if they will collide in 1 second (should be true)
	willCollide := WillCollide(p1, p2, 1)
	assert.True(t, willCollide, "Particles should collide in 1 second")

	// Check if they will collide in 2 seconds (should be false)
	willCollide = WillCollide(p1, p2, 2)
	assert.False(t, willCollide, "Particles should not collide in 2 seconds")
}

func TestHandleCollision(t *testing.T) {
	// Create two movable particles
	p1 := &particle.Particle{
		X:      0,
		Y:      0,
		Vx:     10,
		Vy:     0,
		Mass:   1,
		Radius: 5,
		Movable: true,
	}
	p2 := &particle.Particle{
		X:      15,
		Y:      0,
		Vx:     -5,
		Vy:     0,
		Mass:   1,
		Radius: 5,
		Movable: true,
	}

	// Handle the collision between movable particles
	HandleCollision(p1, p2)

	// Assert that their velocities have been updated
	assert.NotEqual(t, 10.0, p1.Vx, "p1's velocity should change after collision")
	assert.NotEqual(t, -5.0, p2.Vx, "p2's velocity should change after collision")

	// Create an immovable particle
	p3 := &particle.Particle{
		X:      15,
		Y:      0,
		Vx:     0,
		Vy:     0,
		Mass:   1,
		Radius: 5,
		Movable: false,
	}

	// Handle collision with an immovable particle
	HandleCollision(p1, p3)

	// Assert that p3's velocity should remain the same (p3 is immovable)
	assert.Equal(t, 0.0, p3.Vx, "Immovable particle's velocity should not change")
	assert.NotEqual(t, 10.0, p1.Vx, "p1's velocity should change after collision with immovable particle")
}


