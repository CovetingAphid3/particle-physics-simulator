// internal/particle/particle.go
package particle

type Particle struct {
	X, Y    float64
	Vx, Vy float64
	Ax, Ay float64
	Mass       float64
	Radius     float64
	Color      Color
	IsGrounded bool
	Charge     float64
	Fx, Fy float64
    Movable bool
}

type Color struct {
	R, G, B, A float32
}

func NewParticle(x, y,  vx, vy, ax, ay, mass, radius float64, color Color,movable bool) *Particle {
	return &Particle{
		X: x, Y: y,
		Vx: vx, Vy: vy, 
		Ax: ax, Ay: ay,
		Mass:   mass,
		Radius: radius,
		Color:  color,
        Movable: movable,
	}
}

func NewCoulombParticle(x, y,  vx, vy, ax, ay,  mass, radius float64, color Color, charge float64,movable bool) *Particle {
    return &Particle{
        X: x, Y: y, 
        Vx: vx, Vy: vy, 
        Ax: ax, Ay: ay,
        Mass:   mass,
        Radius: radius,
        Color:  color,
        Charge: charge,  // Set the charge for this particle
        Fx: 0.0, Fy: 0.0, 
        Movable: movable,
    }
}
