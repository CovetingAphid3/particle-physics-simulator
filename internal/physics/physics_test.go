package physics_test

import (
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"testing"
	"math"
)

// Helper function to create a test particle
func newTestParticle(x, y, vx, vy float64) *particle.Particle {
	return &particle.Particle{
		X: x,
		Y: y,
		Vx: vx,
		Vy: vy,
		Radius: 10,
		Mass: 1.0,
		IsGrounded: false,
	}
}

func TestApplyGravity(t *testing.T) {
	tests := []struct {
		name      string
		particle  *particle.Particle
		wantAy    float64
		isGrounded bool
	}{
		{
			name:     "Falling particle",
			particle: newTestParticle(100, 100, 0, 0),
			wantAy:   physics.Gravity,
			isGrounded: false,
		},
		{
			name:     "Grounded particle",
			particle: &particle.Particle{X: 100, Y: 100, IsGrounded: true},
			wantAy:   0,
			isGrounded: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			physics.ApplyGravity(tt.particle)
			if tt.particle.Ay != tt.wantAy {
				t.Errorf("ApplyGravity() got Ay = %v, want %v", tt.particle.Ay, tt.wantAy)
			}
		})
	}
}

func TestUpdateVelocity(t *testing.T) {
	tests := []struct {
		name     string
		particle *particle.Particle
		dt       float64
		wantVx   float64
		wantVy   float64
	}{
		{
			name:     "Update with acceleration",
			particle: &particle.Particle{Ax: 1, Ay: 2, IsGrounded: false},
			dt:       1.0,
			wantVx:   1.0, // dt * Ax
			wantVy:   2.0, // dt * Ay
		},
		{
			name:     "Grounded particle",
			particle: &particle.Particle{Ax: 1, Ay: 2, IsGrounded: true},
			dt:       1.0,
			wantVx:   0,
			wantVy:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			physics.UpdateVelocity(tt.particle, tt.dt)
			if math.Abs(tt.particle.Vx-tt.wantVx) > 0.001 {
				t.Errorf("UpdateVelocity() got Vx = %v, want %v", tt.particle.Vx, tt.wantVx)
			}
			if math.Abs(tt.particle.Vy-tt.wantVy) > 0.001 {
				t.Errorf("UpdateVelocity() got Vy = %v, want %v", tt.particle.Vy, tt.wantVy)
			}
		})
	}
}

func TestParticleSimulation(t *testing.T) {
	p := newTestParticle(400, 100, 0, 0)
	dt := 1.0 / 60.0 // Typical frame time
	
	// Record initial state
	initialY := p.Y
	
	// Simulate 10 frames
	numFrames := 10
	for i := 0; i < numFrames; i++ {
		physics.ApplyGravity(p)
		physics.UpdateVelocity(p, dt)
		physics.UpdatePosition(p, dt)
	}
	
	// Particle should have fallen
	if p.Y <= initialY {
		t.Errorf("Particle did not fall: initial Y = %v, final Y = %v", initialY, p.Y)
	}
	
	// Calculate expected velocity:
	// v = v₀ + at, where a is gravity and t is total time
	totalTime := dt * float64(numFrames)
	expectedVy := physics.Gravity * totalTime
	velocityTolerance := physics.Gravity * dt // Allow some tolerance for numerical errors
	
	if math.Abs(p.Vy-expectedVy) > velocityTolerance {
		t.Errorf("Unexpected vertical velocity: got %v, want approximately %v (±%v)", 
			p.Vy, expectedVy, velocityTolerance)
	}
}

// Rest of the tests remain the same...
