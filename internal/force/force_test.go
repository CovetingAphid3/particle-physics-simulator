package force

import (
	// "math"
	"particle-physics-simulator/internal/particle"
	"testing"
)

func TestNewForce(t *testing.T) {
	// Create a new force with value 10 and components (1, 0, 0)
	f := NewForce(10, 1, 0, 0)

	// Test if the force is initialized correctly
	if f.Value != 10 || f.XComponent != 1 || f.YComponent != 0 || f.ZComponent != 0 {
		t.Errorf("Expected Force Value: 10, XComponent: 1, YComponent: 0, ZComponent: 0, but got %v", f)
	}
}

func TestMagnitude(t *testing.T) {
	// Create a force with components (3, 4, 0)
	f := NewForce(1, 3, 4, 0)

	// Magnitude should be sqrt(3^2 + 4^2) = 5
	expectedMagnitude := 5.0
	if f.Magnitude() != expectedMagnitude {
		t.Errorf("Expected magnitude: %.2f, but got %.2f", expectedMagnitude, f.Magnitude())
	}
}

func TestDirection(t *testing.T) {
	// Create a force with components (3, 4, 0)
	f := NewForce(1, 3, 4, 0)

	// Direction should be the normalized vector (3/5, 4/5, 0)
	expectedX, expectedY, expectedZ := 3.0/5.0, 4.0/5.0, 0.0
	x, y, z := f.Direction()

	if x != expectedX || y != expectedY || z != expectedZ {
		t.Errorf("Expected direction: (%.2f, %.2f, %.2f), but got (%.2f, %.2f, %.2f)", expectedX, expectedY, expectedZ, x, y, z)
	}
}

func TestApplyForce(t *testing.T) {
	// Create a force with components (1, 0, 0)
	f := NewForce(10, 1, 0, 0)

	// Create a particle with mass 5 and initial position (0, 0, 0)
	p := particle.NewParticle(0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 1, particle.Color{})

	// Apply the force
	ApplyForce(p, f)

	// The new position should be updated based on the force applied: 
	// p.X += (f.Value * f.XComponent) / p.Mass = (10 * 1) / 5 = 2
	expectedX := 2.0
	if p.X != expectedX {
		t.Errorf("Expected particle X position: %.2f, but got %.2f", expectedX, p.X)
	}
}

func TestApplyForceEdgeCases(t *testing.T) {
	// Test for edge cases where the force components or mass is zero

	// Create a force with zero value
	f := NewForce(0, 0, 0, 0)

	// Create a particle with mass 5 and initial position (0, 0, 0)
	p := particle.NewParticle(0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 1, particle.Color{})

	// Apply the zero force
	ApplyForce(p, f)

	// The position should remain unchanged since the force is zero
	if p.X != 0 || p.Y != 0 {
		t.Errorf("Expected particle position to remain at (0, 0), but got (%.2f, %.2f)", p.X, p.Y)
	}
	
	// Test for a particle with zero mass
	zeroMassParticle := particle.NewParticle(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, particle.Color{})

	// Apply a force to the zero-mass particle
	ApplyForce(zeroMassParticle, f)

	// Since the mass is zero, this will create an issue in the force calculation. It should ideally throw an error or handle gracefully.
	// However, for this simple test, we can just ensure no panics occur. You could improve this test by handling division by zero explicitly in your code.
}


