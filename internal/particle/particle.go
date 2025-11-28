// internal/particle/particle.go
package particle

type ShapeType int

const (
	ShapeCircle ShapeType = iota
	ShapeRectangle
	ShapeBumper
)

type Particle struct {
	X, Y    float64
	Vx, Vy float64
	Ax, Ay float64
	Mass       float64
	Radius     float64 // Used for Circle
	Width, Height float64 // Used for Rectangle/Bumper
	Shape      ShapeType
	Color      Color
	IsGrounded bool
	Charge     float64
	Fx, Fy float64
    Movable bool
	Bounciness float64
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
		Bounciness: 1.0, // Default elastic
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
		Bounciness: 1.0,
	}
}

func NewBumperParticle(x, y, width, height float64, color Color) *Particle {
	return &Particle{
		X: x, Y: y,
		Width: width, Height: height,
		Mass:   100000.0, // Static
		Shape:  ShapeBumper,
		Color:  color,
		Movable: false,
		Bounciness: 1.5, // High bounciness
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
		Bounciness: 1.0,
    }
}
