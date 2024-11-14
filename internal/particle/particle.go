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
    Movable bool
}

type Color struct {
	R, G, B, A float32
}

func NewParticle(x, y, z, vx, vy, vz, ax, ay, az, mass, radius float64, color Color,movable bool) *Particle {
	return &Particle{
		X: x, Y: y, Z: z,
		Vx: vx, Vy: vy, Vz: vz,
		Ax: ax, Ay: ay, Az: az,
		Mass:   mass,
		Radius: radius,
		Color:  color,
        Movable: movable,
	}
}

func NewCoulombParticle(x, y, z, vx, vy, vz, ax, ay, az, mass, radius float64, color Color, charge float64,movable bool) *Particle {
    return &Particle{
        X: x, Y: y, Z: z,
        Vx: vx, Vy: vy, Vz: vz,
        Ax: ax, Ay: ay, Az: az,
        Mass:   mass,
        Radius: radius,
        Color:  color,
        Charge: charge,  // Set the charge for this particle
        Fx: 0.0, Fy: 0.0, Fz: 0.0,  // Initialize forces to zero
        Movable: movable,
    }
}
