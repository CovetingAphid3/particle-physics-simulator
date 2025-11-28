// internal/particle/particle.go
package particle

type ShapeType int

const (
	ShapeCircle ShapeType = iota
	ShapeRectangle
)

type Particle struct {
	X, Y    float64
	Vx, Vy float64
	Ax, Ay float64
	Mass       float64
	Radius     float64 // Used for Circle
	Width, Height float64 // Used for Rectangle
	Shape      ShapeType
	Color      Color
	IsGrounded bool
	Charge     float64
	Fx, Fy float64
    Movable bool
}

type Color struct {
	R, G, B, A float32
}

func NewParticle(x, y,  vx, vy, ax, ay, mass, radius float64, color Color, movable bool) *Particle {
	return &Particle{
		X: x, Y: y,
		Vx: vx, Vy: vy, 
		Ax: ax, Ay: ay,
		Mass:   mass,
		Radius: radius,
		Shape:  ShapeCircle,
		Color:  color,
        Movable: movable,
	}
}

func NewRectangleParticle(x, y, width, height, mass float64, color Color, movable bool) *Particle {
	return &Particle{
		X: x, Y: y,
		Width: width, Height: height,
		Mass:   mass,
		Shape:  ShapeRectangle,
		Color:  color,
		Movable: movable,
	}
}

func NewCoulombParticle(x, y,  vx, vy, ax, ay,  mass, radius float64, color Color, charge float64, movable bool) *Particle {
    return &Particle{
        X: x, Y: y, 
        Vx: vx, Vy: vy, 
        Ax: ax, Ay: ay,
        Mass:   mass,
        Radius: radius,
		Shape:  ShapeCircle,
        Color:  color,
        Charge: charge,  
        Fx: 0.0, Fy: 0.0, 
        Movable: movable,
    }
}
