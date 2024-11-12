// internal/particle/particle.go
package particle

type Particle struct {
	X, Y, Z    float64
	Vx, Vy, Vz float64
	Ax, Ay, Az float64
	Mass       float64
	Radius     float64
	Color      Color
	IsGrounded bool
	Charge     float64
	Fx, Fy, Fz float64
}

type Color struct {
	R, G, B, A float32
}

func NewParticle(x, y, z, vx, vy, vz, ax, ay, az, mass, radius float64, color Color) *Particle {
	return &Particle{
		X: x, Y: y, Z: z,
		Vx: vx, Vy: vy, Vz: vz,
		Ax: ax, Ay: ay, Az: az,
		Mass:   mass,
		Radius: radius,
		Color:  color,
	}
}
