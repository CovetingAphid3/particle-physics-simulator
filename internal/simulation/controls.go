package simulation

import (
	"math"
	"particle-physics-simulator/internal/particle"
	"github.com/gen2brain/raylib-go/raylib"
)

// HandleUserInput handles user interactions for the simulation.
func HandleUserInput(particles *[]*particle.Particle, paused *bool) {
    // Toggle pause with the space bar
    if rl.IsKeyPressed(rl.KeySpace) {
        *paused = !*paused
    }

    // Add particle at mouse position with left-click
    if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
        mouseX := float64(rl.GetMouseX())
        mouseY := float64(rl.GetMouseY())
        newParticle := particle.NewParticle(
            mouseX, mouseY, 
            0, 0,   // Starting velocity
            0, 0,   // Starting acceleration
            10.0,     // Mass
            10,       // Radius
            particle.Color{R: 0.5, G: 0.7, B: 1, A: 1}, // Color
            true,
        )
        *particles = append(*particles, newParticle)
    }

    // Remove particle near mouse position with right-click
    if rl.IsMouseButtonPressed(rl.MouseRightButton) {
        mouseX := float64(rl.GetMouseX())
        mouseY := float64(rl.GetMouseY())
        *particles = removeParticleNear(*particles, mouseX, mouseY, 15.0) // Radius for selection
    }
}

// removeParticleNear removes a particle within a certain distance from (x, y).
func removeParticleNear(particles []*particle.Particle, x, y, radius float64) []*particle.Particle {
    for i, p := range particles {
        dx, dy := p.X - x, p.Y - y
        distance := math.Sqrt(dx*dx + dy*dy)
        if distance <= radius {
            return append(particles[:i], particles[i+1:]...)
        }
    }
    return particles
}

