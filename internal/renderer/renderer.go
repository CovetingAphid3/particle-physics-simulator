package renderer

import (
	"fmt"
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth   int     = 1800
	screenHeight  int     = 950
	groundLevel   float64 = 580
	dampingFactor float64 = 0.9
	buttonSize    int     = 20 
)

func InitWindow() {
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Particle Physics Simulator")
	rl.SetTargetFPS(120) 
}

func DrawParticle(p *particle.Particle) {
	drawParticleCircle(p)
	physics.ApplyBoundaryConditions(p, screenWidth, screenHeight)
}

// drawParticleCircle draws a particle as a circle on the screen using the particle's color and position.
func drawParticleCircle(p *particle.Particle) {
	color := rl.Color{
		R: uint8(p.Color.R * 255),
		G: uint8(p.Color.G * 255),
		B: uint8(p.Color.B * 255),
		A: uint8(p.Color.A * 255),
	}
	rl.DrawCircle(int32(p.X), int32(p.Y), float32(p.Radius), color)
}

// DrawParticleInfo shows particle info (e.g., mass, velocity) when the mouse hovers over a particle.
func DrawParticleInfo(particles []*particle.Particle) {
	mouseX := float64(rl.GetMouseX())
	mouseY := float64(rl.GetMouseY())
	var nearestParticle *particle.Particle
	var minDistance float64 = 1000000 // Set a high initial value

	// Find the nearest particle
	for _, p := range particles {
		dx, dy := p.X-mouseX, p.Y-mouseY
		distance := math.Sqrt(dx*dx + dy*dy)

		if distance < p.Radius && distance < minDistance {
			minDistance = distance
			nearestParticle = p
		}
	}

	if nearestParticle != nil {
		info := fmt.Sprintf("Mass: %.2f, Velocity: (%.2f, %.2f)", nearestParticle.Mass, nearestParticle.Vx, nearestParticle.Vy)
		textHeight := 10
		xPos := int32(nearestParticle.X) + 10
		yPos := int32(nearestParticle.Y) - int32(textHeight) - int32(5)
		rl.DrawText(info, xPos, yPos, 10, rl.Yellow)
	}
}


func DrawUI(particles []*particle.Particle, paused bool) {
	fps := rl.GetFPS()
	particleCount := len(particles)
	pauseStatus := "Running"
	if paused {
		pauseStatus = "Paused"
	}

	// Display FPS, particle count, and status
	rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 10, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Particles: %d", particleCount), 10, 30, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Status: %s", pauseStatus), 10, 50, 20, rl.RayWhite)

	// Display instructions for controls
	instructions := "Controls: [Space] Pause/Resume | [Left Click] Add Particle | [Right Click] Remove Particle"
	rl.DrawText(instructions, 10, int32(screenHeight)-870, 15, rl.Gray)
}

func CloseWindow() {
	rl.CloseWindow()
}

// DrawWindowButtons draws and handles actions for window control buttons like close, minimize, and maximize.
func DrawWindowButtons() {
	closeButtonColor := rl.Red
	maximizeButtonColor := rl.Gray
	minimizeButtonColor := rl.Yellow

	// Button positions (top-right corner of the window)
	closeButtonPos := rl.Rectangle{X: float32(screenWidth - 3*buttonSize), Y: 0, Width: float32(buttonSize), Height: float32(buttonSize)}
	maximizeButtonPos := rl.Rectangle{X: float32(screenWidth - 2*buttonSize), Y: 0, Width: float32(buttonSize), Height: float32(buttonSize)}
	minimizeButtonPos := rl.Rectangle{X: float32(screenWidth - buttonSize), Y: 0, Width: float32(buttonSize), Height: float32(buttonSize)}

	// Draw buttons
	rl.DrawRectangleRec(closeButtonPos, closeButtonColor)
	rl.DrawRectangleRec(maximizeButtonPos, maximizeButtonColor)
	rl.DrawRectangleRec(minimizeButtonPos, minimizeButtonColor)

	// Handle button clicks
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mouseX := float32(rl.GetMouseX())
		mouseY := float32(rl.GetMouseY())

		// Close button action
		if rl.CheckCollisionPointRec(rl.Vector2{X: mouseX, Y: mouseY}, closeButtonPos) {
			rl.CloseWindow() // Close the window
		}
		// Maximize button action
		if rl.CheckCollisionPointRec(rl.Vector2{X: mouseX, Y: mouseY}, maximizeButtonPos) {
			isMaximized := rl.IsWindowMaximized()
			if isMaximized {
				rl.RestoreWindow()
			} else {
				rl.MaximizeWindow()
			}
		}
		// Minimize button action
		if rl.CheckCollisionPointRec(rl.Vector2{X: mouseX, Y: mouseY}, minimizeButtonPos) {
			rl.MinimizeWindow()
		}
	}
}

