// internal/particle/particle.go
package particle

import (
	"particle-physics-simulator/internal/force"
)

type Particle struct {
	X, Y, Z    float64
	Vx, Vy, Vz float64
	Ax, Ay, Az float64
	Mass       float64
	Radius     float64
	Color      Color
	IsGrounded bool
}

type Color struct {
	R, G, B, A float32
}

func NewParticle(x, y, z, vx, vy, vz, ax, ay, az, mass, radius float64, color Color) *Particle {
    return &Particle{
        X: x, Y: y, Z: z,
        Vx: vx, Vy: vy, Vz: vz,
        Ax: ax, Ay: ay, Az: az,
        Mass: mass,
        Radius: radius,
        Color: color,
    }
}


func (p *Particle)ApplyForce(f *force.Force){
    p.Ax += (f.Value * f.XDirection) / p.Mass
    p.Ay += (f.Value * f.YDirection) / p.Mass
}
