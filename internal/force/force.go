package force

import (
	"math"
	"particle-physics-simulator/internal/particle"
)

type Force struct {
    Value      float64
    XComponent float64
    YComponent float64
    ZComponent float64
}

func NewForce(value, xComponent, yComponent, zComponent float64) *Force {
    return &Force{
        Value:      value,
        XComponent: xComponent,
        YComponent: yComponent,
        ZComponent: zComponent,
    }
}

func (f *Force) Magnitude() float64 {
    return math.Sqrt(f.XComponent*f.XComponent + f.YComponent*f.YComponent + f.ZComponent*f.ZComponent)
}

func (f *Force) Direction() (float64, float64, float64) {
    mag := f.Magnitude()
    return f.XComponent / mag, f.YComponent / mag, f.ZComponent / mag
}

// ApplyForce directly updates the position of a particle based on the force components
func ApplyForce(p *particle.Particle, f *Force) {
	p.Vx += (f.Value * f.XComponent) / p.Mass
	p.Vy += (f.Value * f.YComponent) / p.Mass
}

