package simulation

import (
	"math"
	"particle-physics-simulator/internal/constants"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/state"

	"github.com/gen2brain/raylib-go/raylib"
)

// HandleUserInput handles user interactions for the simulation.
func HandleUserInput(particles *[]*particle.Particle, springs *[]*particle.Spring, simState *state.SimulationState) {
	if simState.AppState == state.AppStateMenu {
		handleMenuInput(simState)
	} else {
		handleSimulationInput(particles, springs, simState)
	}
}

func handleMenuInput(simState *state.SimulationState) {
	// Start simulation with Enter
	if rl.IsKeyPressed(rl.KeyEnter) {
		simState.AppState = state.AppStateRunning
	}

	// Adjust particle count
	if rl.IsKeyPressed(rl.KeyUp) {
		simState.ParticleCount += 5
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		if simState.ParticleCount > 1 {
			simState.ParticleCount -= 5
		}
	}
}

func handleSimulationInput(particles *[]*particle.Particle, springs *[]*particle.Spring, simState *state.SimulationState) {
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

	// Toggle Spawn Type with T
	if rl.IsKeyPressed(rl.KeyT) {
		if simState.SpawnType == state.SpawnTypeParticle {
			simState.SpawnType = state.SpawnTypeWall
		} else if simState.SpawnType == state.SpawnTypeWall {
			simState.SpawnType = state.SpawnTypeSpring
		} else if simState.SpawnType == state.SpawnTypeSpring {
			simState.SpawnType = state.SpawnTypeBumper
		} else {
			simState.SpawnType = state.SpawnTypeParticle
		}
	}

	// Toggle Movable Spawning with B
	if rl.IsKeyPressed(rl.KeyB) {
		simState.SpawnMovable = !simState.SpawnMovable
	}

	// Adjust Size with + and - (KeyEqual and KeyMinus)
	if rl.IsKeyDown(rl.KeyEqual) { // Plus
		simState.ParticleSize += 0.5
	}
	if rl.IsKeyDown(rl.KeyMinus) { // Minus
		if simState.ParticleSize > 1.0 {
			simState.ParticleSize -= 0.5
		}
	}

	mouseX := float64(rl.GetMouseX())
	mouseY := float64(rl.GetMouseY())

	// Handle Mouse Actions based on Mode
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if simState.MouseMode == state.MouseModeAdd {
			simState.IsDragging = true
			simState.DragStart = state.Vector2{X: mouseX, Y: mouseY}
		}
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if simState.MouseMode == state.MouseModeAdd && simState.IsDragging {
			simState.IsDragging = false
			dragEnd := state.Vector2{X: mouseX, Y: mouseY}

			if simState.SpawnType == state.SpawnTypeParticle {
				// Just spawn a particle at the click location (DragStart)
				// Or maybe at DragEnd? Let's stick to click location for particles for now, 
				// or maybe drag to set velocity? For now, just simple click spawn.
				// Actually, if we dragged, maybe we shouldn't spawn a simple particle?
				// Let's keep simple click for particles.
				// If distance is small, treat as click.
				dx := dragEnd.X - simState.DragStart.X
				dy := dragEnd.Y - simState.DragStart.Y
				if math.Sqrt(dx*dx+dy*dy) < 5.0 {
					// Click
					mass := math.Pi * simState.ParticleSize * simState.ParticleSize * 0.1
					if mass < 1.0 { mass = 1.0 }
					newParticle := particle.NewParticle(
						simState.DragStart.X, simState.DragStart.Y,
						0, 0, 0, 0, mass, simState.ParticleSize,
						particle.Color{R: 0.5, G: 0.7, B: 1, A: 1}, true,
					)
					*particles = append(*particles, newParticle)
				}
			} else if simState.SpawnType == state.SpawnTypeWall {
				// Drag to create wall
				x := math.Min(simState.DragStart.X, dragEnd.X)
				y := math.Min(simState.DragStart.Y, dragEnd.Y)
				width := math.Abs(dragEnd.X - simState.DragStart.X)
				height := math.Abs(dragEnd.Y - simState.DragStart.Y)

				if width > 5 && height > 5 {
					mass := width * height * 1.0
					if !simState.SpawnMovable {
						mass = 100000.0
					}
					newWall := particle.NewRectangleParticle(
						x, y, width, height, mass,
						particle.Color{R: 0.5, G: 0.5, B: 0.5, A: 1},
						simState.SpawnMovable,
					)
					*particles = append(*particles, newWall)
				}
			} else if simState.SpawnType == state.SpawnTypeBumper {
				// Drag to create bumper
				x := math.Min(simState.DragStart.X, dragEnd.X)
				y := math.Min(simState.DragStart.Y, dragEnd.Y)
				width := math.Abs(dragEnd.X - simState.DragStart.X)
				height := math.Abs(dragEnd.Y - simState.DragStart.Y)

				if width > 5 && height > 5 {
					newBumper := particle.NewBumperParticle(
						x, y, width, height,
						particle.Color{R: 0.8, G: 0.2, B: 0.2, A: 1}, // Reddish for danger/bounce
					)
					*particles = append(*particles, newBumper)
				}
			} else if simState.SpawnType == state.SpawnTypeSpring {
				// Connect two particles
				// Find particle at DragStart and particle at DragEnd
				p1 := findParticleAt(*particles, simState.DragStart.X, simState.DragStart.Y)
				p2 := findParticleAt(*particles, dragEnd.X, dragEnd.Y)

				// Auto-create particles if they don't exist
				if p1 == nil {
					mass := math.Pi * simState.ParticleSize * simState.ParticleSize * 0.1
					if mass < 1.0 { mass = 1.0 }
					p1 = particle.NewParticle(
						simState.DragStart.X, simState.DragStart.Y,
						0, 0, 0, 0, mass, simState.ParticleSize,
						particle.Color{R: 0.5, G: 0.7, B: 1, A: 1}, true,
					)
					*particles = append(*particles, p1)
				}
				if p2 == nil {
					mass := math.Pi * simState.ParticleSize * simState.ParticleSize * 0.1
					if mass < 1.0 { mass = 1.0 }
					p2 = particle.NewParticle(
						dragEnd.X, dragEnd.Y,
						0, 0, 0, 0, mass, simState.ParticleSize,
						particle.Color{R: 0.5, G: 0.7, B: 1, A: 1}, true,
					)
					*particles = append(*particles, p2)
				}

				if p1 != p2 {
					// Create spring
					dx := p1.X - p2.X
					dy := p1.Y - p2.Y
					dist := math.Sqrt(dx*dx + dy*dy)
					
					// If distance is too small, don't create spring (avoid zero length issues)
					if dist > 1.0 {
						spring := particle.NewSpring(p1, p2, 20.0, 0.5) // Default stiffness/damping
						spring.RestLength = dist
						*springs = append(*springs, spring)
					}
				}
			}
		}
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		switch simState.MouseMode {
		// Remove/Attract/Repel logic remains...
		case state.MouseModeRemove:
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) { 
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

func findParticleAt(particles []*particle.Particle, x, y float64) *particle.Particle {
	for _, p := range particles {
		if p.Shape == particle.ShapeCircle {
			dx := p.X - x
			dy := p.Y - y
			if dx*dx+dy*dy < p.Radius*p.Radius {
				return p
			}
		} else {
			if x >= p.X && x <= p.X+p.Width && y >= p.Y && y <= p.Y+p.Height {
				return p
			}
		}
	}
	return nil
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
		// Only affect Circle particles
		if p.Shape != particle.ShapeCircle {
			continue
		}

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

