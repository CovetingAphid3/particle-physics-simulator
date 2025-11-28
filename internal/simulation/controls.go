package simulation

import (
	"math"
	"particle-physics-simulator/internal/constants"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/state"

	"github.com/gen2brain/raylib-go/raylib"
)

// HandleUserInput handles user interactions for the simulation.
func HandleUserInput(particles *[]*particle.Particle, simState *state.SimulationState) {
	if simState.AppState == state.AppStateMenu {
		handleMenuInput(simState)
	} else {
		handleSimulationInput(particles, simState)
	}
}

func handleMenuInput(simState *state.SimulationState) {
	// Start simulation with Enter
	if rl.IsKeyPressed(rl.KeyEnter) {
		simState.AppState = state.AppStateRunning
	}

	// Adjust particle count
	if rl.IsKeyPressed(rl.KeyUp) {
		simState.ParticleCount += 50
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		if simState.ParticleCount > 50 {
			simState.ParticleCount -= 50
		}
	}
}

func handleSimulationInput(particles *[]*particle.Particle, simState *state.SimulationState) {
	// Reset to Menu with R
	if rl.IsKeyPressed(rl.KeyR) {
		simState.AppState = state.AppStateMenu
		*particles = []*particle.Particle{} // Clear particles
		simState.Paused = false
	}

	// Toggle pause with the space bar
	if rl.IsKeyPressed(rl.KeySpace) {
		simState.Paused = !simState.Paused
	}

	// Toggle Gravity with G
	if rl.IsKeyPressed(rl.KeyG) {
		simState.GravityEnabled = !simState.GravityEnabled
	}

	// Toggle Electrostatics with E
	if rl.IsKeyPressed(rl.KeyE) {
		simState.ElectrostaticsEnabled = !simState.ElectrostaticsEnabled
	}

	// Cycle Mouse Mode with M
	if rl.IsKeyPressed(rl.KeyM) {
		simState.MouseMode = (simState.MouseMode + 1) % 4
	}

	// Adjust Gravity Strength with Up/Down Arrows
	if rl.IsKeyDown(rl.KeyUp) {
		simState.GravityStrength += 10.0
	}
	if rl.IsKeyDown(rl.KeyDown) {
		simState.GravityStrength -= 10.0
	}

	// Adjust Attraction Strength with [ and ]
	if rl.IsKeyDown(rl.KeyLeftBracket) {
		simState.AttractionStrength -= 100.0
	}
	if rl.IsKeyDown(rl.KeyRightBracket) {
		simState.AttractionStrength += 100.0
	}

	mouseX := float64(rl.GetMouseX())
	mouseY := float64(rl.GetMouseY())

	// Handle Mouse Actions based on Mode
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		switch simState.MouseMode {
		case state.MouseModeAdd:
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) { // Only add on press, not hold
				newParticle := particle.NewParticle(
					mouseX, mouseY,
					0, 0, // Starting velocity
					0, 0, // Starting acceleration
					10.0, // Mass
					10,   // Radius
					particle.Color{R: 0.5, G: 0.7, B: 1, A: 1}, // Color
					true,
				)
				*particles = append(*particles, newParticle)
			}
		case state.MouseModeRemove:
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) { // Only remove on press
				*particles = removeParticleNear(*particles, mouseX, mouseY, 15.0)
			}
		case state.MouseModeAttract:
			applyMouseForce(*particles, mouseX, mouseY, simState.AttractionStrength)
		case state.MouseModeRepel:
			applyMouseForce(*particles, mouseX, mouseY, -simState.AttractionStrength)
		}
	}

	// Right click can still be a quick remove or context specific, but let's keep it simple for now
	// I'll keep Right Click as "Remove" for now as it's useful.
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		*particles = removeParticleNear(*particles, mouseX, mouseY, 15.0)
	}
}

// removeParticleNear removes a particle within a certain distance from (x, y).
func removeParticleNear(particles []*particle.Particle, x, y, radius float64) []*particle.Particle {
	for i, p := range particles {
		dx, dy := p.X-x, p.Y-y
		distance := math.Sqrt(dx*dx + dy*dy)
		if distance <= radius {
			return append(particles[:i], particles[i+1:]...)
		}
	}
	return particles
}

func applyMouseForce(particles []*particle.Particle, mouseX, mouseY, strength float64) {
	for _, p := range particles {
		dx := mouseX - p.X
		dy := mouseY - p.Y
		distSq := dx*dx + dy*dy
		dist := math.Sqrt(distSq)

		if dist < 5.0 { // Avoid singularity
			dist = 5.0
		}

		// F = strength / dist
		// Force direction is (dx/dist, dy/dist)
		// Fx = F * (dx/dist) = strength * dx / distSq
		forceMag := strength / dist 
		
		fx := forceMag * (dx / dist)
		fy := forceMag * (dy / dist)

		p.Vx += fx * constants.SecondsPerFrame 
		p.Vy += fy * constants.SecondsPerFrame
	}
}

