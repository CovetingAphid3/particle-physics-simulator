package particle

import (
	"testing"
	"reflect"
)

func TestNewParticle(t *testing.T) {
	color := Color{R: 1, G: 0, B: 0, A: 1}
	p := NewParticle(10, 20, 30, 1, 1, 1, 0, 0, 0, 5, 2, color)

	// Check that the particle's fields match the values we passed in
	if p.X != 10 || p.Y != 20 || p.Z != 30 {
		t.Errorf("NewParticle() failed to initialize position. Got: (%f, %f, %f)", p.X, p.Y, p.Z)
	}
	if p.Vx != 1 || p.Vy != 1 || p.Vz != 1 {
		t.Errorf("NewParticle() failed to initialize velocity. Got: (%f, %f, %f)", p.Vx, p.Vy, p.Vz)
	}
	if p.Mass != 5 || p.Radius != 2 {
		t.Errorf("NewParticle() failed to initialize mass or radius. Got: mass=%f, radius=%f", p.Mass, p.Radius)
	}
	if !reflect.DeepEqual(p.Color, color) {
		t.Errorf("NewParticle() failed to initialize color. Got: %+v", p.Color)
	}
}

