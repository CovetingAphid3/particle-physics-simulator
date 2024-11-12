package collisions_test

import (
	"testing"
	"log"

	"github.com/stretchr/testify/assert"
	"particle-physics-simulator/internal/collisions"
	"particle-physics-simulator/internal/particle"
)

func TestCheckCollision(t *testing.T) {
	// Create two particles
	p1 := particle.NewParticle(0, 0, 0, 0, 0, 0, 0, 0, 0, 1.0, 1.0, particle.Color{R: 1, G: 0, B: 0, A: 1})
	p2 := particle.NewParticle(1.5, 0, 0, 0, 0, 0, 0, 0, 0, 1.0, 1.0, particle.Color{R: 0, G: 1, B: 0, A: 1})

	// Test for collision (distance is exactly the sum of their radii)
	collides := collisions.CheckCollision(p1, p2)
	assert.True(t, collides, "Particles should collide at distance 1.5")

	// Test for no collision (particles are far apart)
	p2.X = 3.0 // Move p2 further apart
	collides = collisions.CheckCollision(p1, p2)
	assert.False(t, collides, "Particles should not collide at distance 3.0")
}

func TestWillCollide(t *testing.T) {
	// Create two particles moving towards each other
	p1 := particle.NewParticle(0, 0, 0, 2.0, 0, 0, 0, 0, 0, 1.0, 1.0, particle.Color{R: 1, G: 0, B: 0, A: 1})
	p2 := particle.NewParticle(5, 0, 0, -2.0, 0, 0, 0, 0, 0, 1.0, 1.0, particle.Color{R: 0, G: 1, B: 0, A: 1})

	// Test for collision prediction (will collide in 1 second)
	collides := collisions.WillCollide(p1, p2, 1.0)
	assert.True(t, collides, "Particles should collide in 1 second")

	// Test for no collision prediction (particles moving apart)
	p2.X = 10 // Move p2 further away
	collides = collisions.WillCollide(p1, p2, 1.0)
	assert.False(t, collides, "Particles should not collide in 1 second")
}

func TestHandleCollision(t *testing.T) {
	// Create two particles
	p1 := particle.NewParticle(0, 0, 0, 1.0, 0, 0, 0, 0, 0, 1.0, 1.0, particle.Color{R: 1, G: 0, B: 0, A: 1})
	p2 := particle.NewParticle(2, 0, 0, -1.0, 0, 0, 0, 0, 0, 1.0, 1.0, particle.Color{R: 0, G: 1, B: 0, A: 1})

	// Log initial velocities
	log.Printf("Initial velocities: p1(%.2f, %.2f, %.2f), p2(%.2f, %.2f, %.2f)", p1.Vx, p1.Vy, p1.Vz, p2.Vx, p2.Vy, p2.Vz)

	// Handle collision between p1 and p2
	collisions.HandleCollision(p1, p2)

	// Check if the velocities were updated after collision
	assert.NotEqual(t, p1.Vx, 1.0, "Particle 1's velocity should have changed after collision")
	assert.NotEqual(t, p2.Vx, -1.0, "Particle 2's velocity should have changed after collision")

	// Log final velocities
	log.Printf("Updated velocities: p1(%.2f, %.2f, %.2f), p2(%.2f, %.2f, %.2f)", p1.Vx, p1.Vy, p1.Vz, p2.Vx, p2.Vy, p2.Vz)
}

