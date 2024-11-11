//internal/particle/analytics
package particle

import "fmt"

// internal/particle/analytics.go
func (p *Particle) GetInfo() string {
    return fmt.Sprintf("Mass: %.2f, Position: (%.6f, %.6f), Velocity: (%.6f, %.6f), Grounded: %v",
        p.Mass, p.X, p.Y, p.Vx, p.Vy, p.IsGrounded)
}
