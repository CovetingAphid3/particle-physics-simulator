// internal/simulation/simulation.go
package simulation

import (
	// "fmt"
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"particle-physics-simulator/internal/renderer"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

const TimeStep = 1.0 / 60.0 // 60 FPS

func WillCollide(p1, p2 *particle.Particle, dt float64) bool {
    // Predict the future position of both particles
    p1NextX := p1.X + p1.Vx*dt
    p1NextY := p1.Y + p1.Vy*dt
    p2NextX := p2.X + p2.Vx*dt
    p2NextY := p2.Y + p2.Vy*dt

    // Calculate the distance between the predicted positions
    dx := p1NextX - p2NextX
    dy := p1NextY - p2NextY
    distSq := dx*dx + dy*dy

    // Check if the distance is less than the sum of their radii (i.e., a collision)
    return distSq < (p1.Radius + p2.Radius)*(p1.Radius + p2.Radius)
}


// HandleCollision updates the velocities of two particles after collision
func HandleCollision(p1, p2 *particle.Particle) {
    dx := p1.X - p2.X
    dy := p1.Y - p2.Y
    dz := p1.Z - p2.Z

    // Distance squared to avoid the expensive sqrt calculation
    distSq := dx*dx + dy*dy + dz*dz

    if distSq == 0 {
        return // Avoid division by zero if particles overlap completely
    }

    // Inverse of distance for normalizing
    invDist := 1.0 / math.Sqrt(distSq)

    // Normalized direction vector of the collision
    nx := dx * invDist
    ny := dy * invDist
    nz := dz * invDist

    // Relative velocity between particles
    vx := p1.Vx - p2.Vx
    vy := p1.Vy - p2.Vy
    vz := p1.Vz - p2.Vz

    // Dot product of relative velocity and collision normal
    dotProduct := vx*nx + vy*ny + vz*nz

    // Only proceed if particles are moving toward each other
    if dotProduct > 0 {
        return
    }

    // Calculate impulse scalar for elastic collision
    impulse := 2 * dotProduct / (p1.Mass + p2.Mass)

    // Update velocities based on impulse
    p1.Vx -= impulse * p2.Mass * nx
    p1.Vy -= impulse * p2.Mass * ny
    p1.Vz -= impulse * p2.Mass * nz

    p2.Vx += impulse * p1.Mass * nx
    p2.Vy += impulse * p1.Mass * ny
    p2.Vz += impulse * p1.Mass * nz
}

// Run the simulation
// Run the simulation
func RunSimulation(particles []*particle.Particle) {
    renderer.InitWindow()
    defer renderer.CloseWindow()

    for !rl.WindowShouldClose() {
        // Update particle physics
        for _, p := range particles {
            physics.ApplyGravity(p)
            physics.UpdateVelocity(p, TimeStep)
            physics.UpdatePosition(p, TimeStep)
        }

        // Check for collisions between particles using swept collision detection
        for i := 0; i < len(particles); i++ {
            for j := i + 1; j < len(particles); j++ {
                if WillCollide(particles[i], particles[j], TimeStep) {
                    HandleCollision(particles[i], particles[j])
                }
            }
        }

        // Render particles
        rl.BeginDrawing()
        rl.ClearBackground(rl.RayWhite)

        for _, p := range particles {
            renderer.DrawParticle(p)
        }

        rl.EndDrawing()

        time.Sleep(time.Millisecond * 16) // Simulate 60 FPS
    }
}



func RunSimulationSingle(p *particle.Particle) {
    RunSimulation([]*particle.Particle{p})
}
