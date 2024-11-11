// internal/particle/particle.go
package particle

type Particle struct {
    X, Y, Z      float64
    Vx, Vy, Vz   float64
    Ax, Ay, Az   float64
    Mass         float64
    Radius       float64
    Color         Color
}

type Color struct {
    R, G, B, A float32
}

