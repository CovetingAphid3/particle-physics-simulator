package simulation

import (
	"math"
	"particle-physics-simulator/internal/collisions"
	// "particle-physics-simulator/internal/force"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"particle-physics-simulator/internal/renderer"
	"particle-physics-simulator/internal/state"
	"sync"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	TimeStep       = 1.0 / 120.0 // Target simulation time step (120 FPS)
	MagneticFieldX = 0.1      
	MagneticFieldY = 0.0     
)

func RunSimulation() {
	renderer.InitWindow()
	defer renderer.CloseWindow()

	simState := state.SimulationState{
		AppState:             state.AppStateMenu,
		Paused:               false,
		GravityEnabled:       true,
		ElectrostaticsEnabled: false,
		GravityStrength:      980.0,
		MouseMode:            state.MouseModeAdd,
		AttractionStrength:   5000.0,
		ParticleCount:        10, // Default
		SpawnType:            state.SpawnTypeParticle,
		ParticleSize:         10.0,
		SpawnMovable:         false,
	}
	var wg sync.WaitGroup
	var particles []*particle.Particle
	var springs []*particle.Spring

	for !rl.WindowShouldClose() {
		currentTime := time.Now()

		// Handle user input
		HandleUserInput(&particles, &springs, &simState)

		// Initialize particles if transitioning to running state
		if simState.AppState == state.AppStateRunning && len(particles) == 0 {
			particles = InitializeParticles(simState.ParticleCount)
		}

		if simState.AppState == state.AppStateRunning && !simState.Paused {
			// Apply global forces (electrostatic)
			if simState.ElectrostaticsEnabled {
				physics.ApplyElectrostaticForces(particles)
			}
			
			// Apply Spring Forces
			physics.ApplySpringForces(springs)

			// Parallelize particle updates
			wg.Add(len(particles))
			for _, p := range particles {
				go func(p *particle.Particle) {
					defer wg.Done()
					currentGravity := 0.0
					if simState.GravityEnabled {
						currentGravity = simState.GravityStrength
					}
					physics.UpdateVelocity(p, TimeStep, currentGravity)
					physics.UpdatePosition(p, TimeStep)
				}(p)
			}
			wg.Wait()

			// Check and handle collisions between particles in parallel
			wg.Add(len(particles) * (len(particles) - 1) / 2) // Maximum number of collisions
			for i := 0; i < len(particles); i++ {
				for j := i + 1; j < len(particles); j++ {
					go func(i, j int) {
						defer wg.Done()
						if collisions.WillCollide(particles[i], particles[j], TimeStep) {
							collisions.HandleCollision(particles[i], particles[j])
						}
					}(i, j)
				}
			}
			wg.Wait()
		}

		// Render the simulation
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if simState.AppState == state.AppStateRunning {
			renderer.DrawSprings(springs)
			for _, p := range particles {
				renderer.DrawParticle(p)
			}
			renderer.DrawParticleInfo(particles)
		}

		renderer.DrawUI(particles, &simState)
		renderer.DrawWindowButtons()

		rl.EndDrawing()

		// Sleep to maintain a consistent frame rate (120 FPS)
		sleepDuration := time.Second/120 - time.Since(currentTime)
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}

func InitializeParticles(count int) []*particle.Particle {
	color1 := particle.Color{R: 1, G: 0, B: 0, A: 1}
	color2 := particle.Color{R: 0, G: 1, B: 0, A: 1}
	color3 := particle.Color{R: 0, G: 0, B: 1, A: 1}
	color4 := particle.Color{R: 1, G: 1, B: 0, A: 1}
	color5 := particle.Color{R: 1, G: 0, B: 1, A: 1}

	particles := []*particle.Particle{}
	
	// Create particles based on count
	// Let's split count between normal and magnetic/charged for variety if we want, 
	// or just make them all charged for now as per original main.go logic.
	// Original had 100 normal (charged) + 100 magnetic (charged).
	// Let's just make them all charged for simplicity and variety.
	
	for i := 0; i < count; i++ {
		x := float64(i % 500)
		y := float64((i * 100) % 500)
		velocityX := math.Sin(float64(i)*0.1) * 100
		velocityY := math.Cos(float64(i)*0.1) * 100
		charge := float64(math.Cos(float64(i)*0.1) * 0.100)

		var color particle.Color
		switch i % 5 {
		case 0:
			color = color1
		case 1:
			color = color2
		case 2:
			color = color3
		case 3:
			color = color4
		default:
			color = color5
		}

		p := particle.NewCoulombParticle(x, y, velocityX, velocityY, 0.0, 0.0, 5.0, 5, color, charge, true)
		particles = append(particles, p)
	}

	// Always add the obstacle? Or maybe not if count is small? Let's keep it.
	// obstacle := particle.NewParticle(300, 300, 0.0, 0.0, 0.0, 0.0, 100.0, 50, particle.Color{R: 0.5, G: 0.5, B: 0.5, A: 1}, false)
	// particles = append(particles, obstacle)

	return particles
}
