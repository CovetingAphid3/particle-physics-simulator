package physics

import (
    "fmt"
    "particle-physics-simulator/internal/particle"
    "testing"
)

// TestCollisionDetection runs a basic test for collision detection between two particles.
func TestCollisionDetection(t *testing.T) {
    // Initialize two particles close enough to collide
    p1 := &particle.Particle{X: 0, Y: 0, Radius: 10, Mass: 1, Vx: 1, Vy: 0}
    p2 := &particle.Particle{X: 15, Y: 0, Radius: 10, Mass: 1, Vx: -1, Vy: 0}

    // Run for a few simulation steps to check for collision
    for step := 0; step < 10; step++ {
        fmt.Printf("Step %d:\n", step)
        fmt.Printf("Particle 1 - Pos: (%.2f, %.2f) Vel: (%.2f, %.2f)\n", p1.X, p1.Y, p1.Vx, p1.Vy)
        fmt.Printf("Particle 2 - Pos: (%.2f, %.2f) Vel: (%.2f, %.2f)\n", p2.X, p2.Y, p2.Vx, p2.Vy)

        // Check if they collide
        if CheckCollision(p1, p2) {
            fmt.Println("Collision detected!")
            HandleCollision(p1, p2) // You’ll implement this to adjust velocities
        }

        // Update positions (assuming a fixed timestep, e.g., dt = 1 for simplicity)
        UpdateVelocity(p1, 1)
        UpdateVelocity(p2, 1)
        UpdatePosition(p1, 1)
        UpdatePosition(p2, 1)
    }
}

