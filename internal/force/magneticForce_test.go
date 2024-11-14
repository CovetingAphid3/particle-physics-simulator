package force

import (
	"particle-physics-simulator/internal/particle"
	"testing"
)

func TestMagneticForce(t *testing.T) {
	// Create a particle with a known charge and velocity
	p := &particle.Particle{
		Charge: 1.0, // 1 Coulomb
		Vx:     2.0, // 2 m/s
		Vy:     3.0, // 3 m/s
		Vz:     4.0, // 4 m/s
	}

	// Define the magnetic field (Bx, By, Bz)
	magneticFieldX := 1.0 // 1 Tesla in the x-direction
	magneticFieldY := 0.0 // No magnetic field in the y-direction
	magneticFieldZ := 0.0 // No magnetic field in the z-direction

	// Calculate the expected magnetic force
	expectedForceX := p.Charge * (p.Vy*magneticFieldZ - p.Vz*magneticFieldY) // F_x = q * (v_y * B_z - v_z * B_y)
	expectedForceY := p.Charge * (p.Vz*magneticFieldX - p.Vx*magneticFieldZ) // F_y = q * (v_z * B_x - v_x * B_z)
	expectedForceZ := p.Charge * (p.Vx*magneticFieldY - p.Vy*magneticFieldX) // F_z = q * (v_x * B_y - v_y * B_x)

	// Call the function under test
	forceX, forceY, forceZ := MagneticForce(p, magneticFieldX, magneticFieldY, magneticFieldZ)

	// Check if the calculated magnetic forces are correct
	if forceX != expectedForceX {
		t.Errorf("Expected force in X direction: %.2f, but got: %.2f", expectedForceX, forceX)
	}
	if forceY != expectedForceY {
		t.Errorf("Expected force in Y direction: %.2f, but got: %.2f", expectedForceY, forceY)
	}
	if forceZ != expectedForceZ {
		t.Errorf("Expected force in Z direction: %.2f, but got: %.2f", expectedForceZ, forceZ)
	}
}

