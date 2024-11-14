package physics

import (
	"math"
	"testing"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/constants"
)

func TestApplyGravity(t *testing.T) {
	tests := []struct {
		name       string
		particle   *particle.Particle
		wantAy     float64
		isGrounded bool
	}{
		{
			name: "Ungrounded particle should have gravity applied",
			particle: &particle.Particle{
				IsGrounded: false,
				Ay:        0,
			},
			wantAy:     constants.Gravity,
			isGrounded: false,
		},
		{
			name: "Grounded particle should not have gravity applied",
			particle: &particle.Particle{
				IsGrounded: true,
				Ay:        0,
			},
			wantAy:     0,
			isGrounded: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ApplyGravity(tt.particle)
			if tt.particle.Ay != tt.wantAy {
				t.Errorf("ApplyGravity() got Ay = %v, want %v", tt.particle.Ay, tt.wantAy)
			}
		})
	}
}

func TestApplyAirFriction(t *testing.T) {
	initialVx := 10.0
	initialVy := -5.0
	initialVz := 3.0
	p := &particle.Particle{
		Vx: initialVx,
		Vy: initialVy,
		Vz: initialVz,
	}

	ApplyAirFriction(p)

	expectedVx := initialVx * (1 - constants.AirDragCoefficient)
	expectedVy := initialVy * (1 - constants.AirDragCoefficient)
	expectedVz := initialVz * (1 - constants.AirDragCoefficient)

	if math.Abs(p.Vx-expectedVx) > 1e-10 {
		t.Errorf("ApplyAirFriction() got Vx = %v, want %v", p.Vx, expectedVx)
	}
	if math.Abs(p.Vy-expectedVy) > 1e-10 {
		t.Errorf("ApplyAirFriction() got Vy = %v, want %v", p.Vy, expectedVy)
	}
	if math.Abs(p.Vz-expectedVz) > 1e-10 {
		t.Errorf("ApplyAirFriction() got Vz = %v, want %v", p.Vz, expectedVz)
	}
}

// func TestUpdateVelocity(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		particle *particle.Particle
// 		dt       float64
// 		want     *particle.Particle
// 	}{
// 		{
// 			name: "Movable ungrounded particle",
// 			particle: &particle.Particle{
// 				Movable:    true,
// 				IsGrounded: false,
// 				Ax:         2.0,
// 				Ay:         1.0,
// 				Az:         0.5,
// 				Vx:         1.0,
// 				Vy:         1.0,
// 				Vz:         1.0,
// 			},
// 			dt: 0.1,
// 			want: &particle.Particle{
// 				Movable:    true,
// 				IsGrounded: false,
// 				Vx:         1.2,
// 				Vy:         1.1,
// 				Vz:         1.05,
// 			},
// 		},
// 		{
// 			name: "Immovable particle",
// 			particle: &particle.Particle{
// 				Movable:    false,
// 				IsGrounded: false,
// 				Ax:         2.0,
// 				Ay:         1.0,
// 				Az:         0.5,
// 				Vx:         1.0,
// 				Vy:         1.0,
// 				Vz:         1.0,
// 			},
// 			dt: 0.1,
// 			want: &particle.Particle{
// 				Movable:    false,
// 				IsGrounded: false,
// 				Vx:         1.0,
// 				Vy:         1.0,
// 				Vz:         1.0,
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			UpdateVelocity(tt.particle, tt.dt)
// 			
// 			if math.Abs(tt.particle.Vx-tt.want.Vx) > 1e-10 {
// 				t.Errorf("UpdateVelocity() got Vx = %v, want %v", tt.particle.Vx, tt.want.Vx)
// 			}
// 			if math.Abs(tt.particle.Vy-tt.want.Vy) > 1e-10 {
// 				t.Errorf("UpdateVelocity() got Vy = %v, want %v", tt.particle.Vy, tt.want.Vy)
// 			}
// 			if math.Abs(tt.particle.Vz-tt.want.Vz) > 1e-10 {
// 				t.Errorf("UpdateVelocity() got Vz = %v, want %v", tt.particle.Vz, tt.want.Vz)
// 			}
// 		})
// 	}
// }


func TestUpdatePosition(t *testing.T) {
	tests := []struct {
		name     string
		particle *particle.Particle
		dt       float64
		want     *particle.Particle
	}{
		{
			name: "Movable ungrounded particle",
			particle: &particle.Particle{
				Movable:    true,
				IsGrounded: false,
				X:          0,
				Y:          0,
				Z:          0,
				Vx:         1.0,
				Vy:         2.0,
				Vz:         3.0,
			},
			dt: 0.5,
			want: &particle.Particle{
				X: 0.5,
				Y: 1.0,
				Z: 1.5,
			},
		},
		{
			name: "Movable grounded particle",
			particle: &particle.Particle{
				Movable:    true,
				IsGrounded: true,
				X:          0,
				Y:          10,
				Z:          0,
				Vx:         1.0,
				Vy:         2.0,
				Vz:         3.0,
			},
			dt: 0.5,
			want: &particle.Particle{
				X: 0.5,
				Y: 10,  // Y position shouldn't change when grounded
				Z: 1.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdatePosition(tt.particle, tt.dt)
			
			if math.Abs(tt.particle.X-tt.want.X) > 1e-10 {
				t.Errorf("UpdatePosition() got X = %v, want %v", tt.particle.X, tt.want.X)
			}
			if math.Abs(tt.particle.Y-tt.want.Y) > 1e-10 {
				t.Errorf("UpdatePosition() got Y = %v, want %v", tt.particle.Y, tt.want.Y)
			}
			if math.Abs(tt.particle.Z-tt.want.Z) > 1e-10 {
				t.Errorf("UpdatePosition() got Z = %v, want %v", tt.particle.Z, tt.want.Z)
			}
		})
	}
}


func TestApplyMagneticForces(t *testing.T) {
	particles := []*particle.Particle{
		{
			Mass:   1.0,
			Charge: 1.0,
			Vx:     1.0,
			Vy:     1.0,
			Vz:     1.0,
		},
	}

	magneticFieldX := 1.0
	magneticFieldY := 0.0
	magneticFieldZ := 0.0

	ApplyMagneticForces(particles, magneticFieldX, magneticFieldY, magneticFieldZ)

	// The magnetic force should be perpendicular to both velocity and magnetic field
	if particles[0].Ax == 0 && particles[0].Ay == 0 && particles[0].Az == 0 {
		t.Error("ApplyMagneticForces() did not apply any forces")
	}
}

func TestApplyElectrostaticForces(t *testing.T) {
	particles := []*particle.Particle{
		{
			X:      0,
			Y:      0,
			Z:      0,
			Charge: 1.0,
		},
		{
			X:      1.0,
			Y:      0,
			Z:      0,
			Charge: -1.0,
		},
	}

	ApplyElectrostaticForces(particles)

	// Opposite charges should attract
	if particles[0].Fx >= 0 {
		t.Error("ApplyElectrostaticForces() did not create attractive force between opposite charges")
	}
	if particles[1].Fx <= 0 {
		t.Error("ApplyElectrostaticForces() did not create attractive force between opposite charges")
	}
}
