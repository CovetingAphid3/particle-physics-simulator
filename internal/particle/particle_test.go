package particle

import (
	"testing"
	"reflect"
)

// Test for NewParticle function
func TestNewParticle(t *testing.T) {
	color := Color{R: 0.5, G: 0.5, B: 0.5, A: 1.0}
	p := NewParticle(1.0, 2.0, 3.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, color, true)

	if p.X != 1.0 || p.Y != 2.0 || p.Z != 3.0 {
		t.Errorf("Expected position (1.0, 2.0, 3.0), got (%v, %v, %v)", p.X, p.Y, p.Z)
	}
	if p.Vx != 0.0 || p.Vy != 0.0 || p.Vz != 0.0 {
		t.Errorf("Expected velocity (0.0, 0.0, 0.0), got (%v, %v, %v)", p.Vx, p.Vy, p.Vz)
	}
	if p.Mass != 1.0 || p.Radius != 1.0 {
		t.Errorf("Expected mass 1.0 and radius 1.0, got mass %v and radius %v", p.Mass, p.Radius)
	}
	if !reflect.DeepEqual(p.Color, color) {
		t.Errorf("Expected color %+v, got %+v", color, p.Color)
	}
	if !p.Movable {
		t.Errorf("Expected particle to be movable")
	}
}

// Test for NewCoulombParticle function
func TestNewCoulombParticle(t *testing.T) {
	color := Color{R: 0.5, G: 0.5, B: 0.5, A: 1.0}
	p := NewCoulombParticle(1.0, 2.0, 3.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, color, 1.0)

	if p.X != 1.0 || p.Y != 2.0 || p.Z != 3.0 {
		t.Errorf("Expected position (1.0, 2.0, 3.0), got (%v, %v, %v)", p.X, p.Y, p.Z)
	}
	if p.Vx != 0.0 || p.Vy != 0.0 || p.Vz != 0.0 {
		t.Errorf("Expected velocity (0.0, 0.0, 0.0), got (%v, %v, %v)", p.Vx, p.Vy, p.Vz)
	}
	if p.Mass != 1.0 || p.Radius != 1.0 {
		t.Errorf("Expected mass 1.0 and radius 1.0, got mass %v and radius %v", p.Mass, p.Radius)
	}
	if !reflect.DeepEqual(p.Color, color) {
		t.Errorf("Expected color %+v, got %+v", color, p.Color)
	}
	if p.Charge != 1.0 {
		t.Errorf("Expected charge 1.0, got %v", p.Charge)
	}
}

