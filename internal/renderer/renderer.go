package renderer

import (
	"fmt"
	"math"
	"particle-physics-simulator/internal/particle"
	"particle-physics-simulator/internal/physics"
	"particle-physics-simulator/internal/state"

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
	color := rl.Color{
		R: uint8(p.Color.R * 255),
		G: uint8(p.Color.G * 255),
		B: uint8(p.Color.B * 255),
		A: uint8(p.Color.A * 255),
	}

	if p.Shape == particle.ShapeRectangle {
		rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), color)
	} else if p.Shape == particle.ShapeBumper {
		// Draw background
		rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), color)
		// Draw coil pattern
		rl.DrawRectangleLines(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), rl.White)
		
		// Draw sine wave or zigzag inside
		startX := int32(p.X)
		startY := int32(p.Y) + int32(p.Height)/2
		endX := int32(p.X + p.Width)
		
		prevX := startX
		prevY := startY
		
		step := 5
		for x := startX; x <= endX; x += int32(step) {
			t := float64(x - startX) * 0.2
			yOffset := math.Sin(t) * (p.Height * 0.4)
			currY := startY + int32(yOffset)
			
			rl.DrawLine(prevX, prevY, x, currY, rl.White)
			prevX = x
			prevY = currY
		}

	} else {
		rl.DrawCircle(int32(p.X), int32(p.Y), float32(p.Radius), color)
	}
	
	physics.ApplyBoundaryConditions(p, screenWidth, screenHeight)
}

// drawParticleCircle is deprecated/merged into DrawParticle
// func drawParticleCircle(p *particle.Particle) { ... }

// DrawParticleInfo shows particle info (e.g., mass, velocity) when the mouse hovers over a particle.
func DrawParticleInfo(particles []*particle.Particle) {
	mouseX := float64(rl.GetMouseX())
	mouseY := float64(rl.GetMouseY())
	var nearestParticle *particle.Particle
	var minDistance float64 = 1000000 // Set a high initial value

	// Find the nearest particle
	for _, p := range particles {
		var distance float64
		if p.Shape == particle.ShapeCircle {
			dx, dy := p.X-mouseX, p.Y-mouseY
			distance = math.Sqrt(dx*dx + dy*dy)
			if distance < p.Radius && distance < minDistance {
				minDistance = distance
				nearestParticle = p
			}
		} else {
			// Simple rect hover check
			if mouseX >= p.X && mouseX <= p.X+p.Width && mouseY >= p.Y && mouseY <= p.Y+p.Height {
				distance = 0 // Inside
				minDistance = 0
				nearestParticle = p
			}
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


func DrawSprings(springs []*particle.Spring) {
	for _, s := range springs {
		start := rl.Vector2{X: float32(s.P1.X), Y: float32(s.P1.Y)}
		end := rl.Vector2{X: float32(s.P2.X), Y: float32(s.P2.Y)}
		
		// Calculate direction and length
		diff := rl.Vector2Subtract(end, start)
		length := rl.Vector2Length(diff)
		
		if length == 0 {
			continue
		}

		// Normalize direction
		dir := rl.Vector2Scale(diff, 1.0/length)
		// Perpendicular vector for zigzag
		perp := rl.Vector2{X: -dir.Y, Y: dir.X}

		// Zigzag parameters
		segments := int(length / 10.0) // One segment every 10 pixels
		if segments < 2 {
			segments = 2
		}
		width := float32(5.0) // Width of the zigzag

		// Draw zigzag
		prevPoint := start
		for i := 1; i <= segments; i++ {
			t := float32(i) / float32(segments)
			currentBase := rl.Vector2Add(start, rl.Vector2Scale(diff, t))
			
			offset := float32(0.0)
			if i < segments { // Don't offset the last point (it should be exactly at 'end')
				if i%2 == 0 {
					offset = width
				} else {
					offset = -width
				}
			}
			
			currentPoint := rl.Vector2Add(currentBase, rl.Vector2Scale(perp, offset))
			
			// If it's the last segment, connect to end exactly
			if i == segments {
				currentPoint = end
			}

			rl.DrawLineEx(prevPoint, currentPoint, 2.0, rl.White)
			prevPoint = currentPoint
		}
	}
}

func DrawUI(particles []*particle.Particle, simState *state.SimulationState) {
	if simState.AppState == state.AppStateMenu {
		DrawMenu(simState)
		return
	}

	// Draw Dragging Rectangle/Line
	if simState.IsDragging {
		mouseX := int32(rl.GetMouseX())
		mouseY := int32(rl.GetMouseY())
		startX := int32(simState.DragStart.X)
		startY := int32(simState.DragStart.Y)

		if simState.SpawnType == state.SpawnTypeWall || simState.SpawnType == state.SpawnTypeBumper {
			// Draw Rectangle
			x := min(startX, mouseX)
			y := min(startY, mouseY)
			width := abs(mouseX - startX)
			height := abs(mouseY - startY)
			color := rl.Green
			if simState.SpawnType == state.SpawnTypeBumper {
				color = rl.Red
			}
			rl.DrawRectangleLines(x, y, width, height, color)
		} else if simState.SpawnType == state.SpawnTypeSpring {
			// Draw Line
			rl.DrawLine(startX, startY, mouseX, mouseY, rl.Green)
		}
	}

	fps := rl.GetFPS()
	particleCount := len(particles)
	pauseStatus := "Running"
	if simState.Paused {
		pauseStatus = "Paused"
	}

	// Display FPS, particle count, and status
	rl.DrawText(fmt.Sprintf("FPS: %d", fps), 10, 10, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Particles: %d", particleCount), 10, 30, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Status: %s", pauseStatus), 10, 50, 20, rl.RayWhite)

	// Display Simulation State
	gravityStatus := "OFF"
	if simState.GravityEnabled {
		gravityStatus = fmt.Sprintf("ON (%.0f)", simState.GravityStrength)
	}
	rl.DrawText(fmt.Sprintf("Gravity (G): %s", gravityStatus), 10, 80, 20, rl.RayWhite)

	electroStatus := "OFF"
	if simState.ElectrostaticsEnabled {
		electroStatus = "ON"
	}
	rl.DrawText(fmt.Sprintf("Electrostatics (E): %s", electroStatus), 10, 100, 20, rl.RayWhite)

	rl.DrawText(fmt.Sprintf("Mouse Mode (M): %s", simState.MouseMode.String()), 10, 120, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Attraction Strength ([/]): %.0f", simState.AttractionStrength), 10, 140, 20, rl.RayWhite)
	
	rl.DrawText(fmt.Sprintf("Spawn Type (T): %s", simState.SpawnType.String()), 10, 160, 20, rl.RayWhite)
	movableStatus := "Static"
	if simState.SpawnMovable {
		movableStatus = "Movable"
	}
	rl.DrawText(fmt.Sprintf("Spawn Mode (B): %s", movableStatus), 10, 180, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Size (+/-): %.1f", simState.ParticleSize), 10, 200, 20, rl.RayWhite)

	// Display instructions for controls
	instructions := "Controls: [Space] Pause | [G] Gravity | [E] Electrostatics | [M] Mouse Mode | [T] Type | [B] Movable | [+/-] Size"
	rl.DrawText(instructions, 10, int32(screenHeight)-30, 20, rl.Gray)
}

func DrawMenu(simState *state.SimulationState) {
	rl.ClearBackground(rl.Black)
	
	title := "Particle Physics Simulator"
	rl.DrawText(title, int32(screenWidth)/2 - 200, int32(screenHeight)/3, 40, rl.RayWhite)

	instructions := "Use Up/Down arrows to adjust particle count"
	rl.DrawText(instructions, int32(screenWidth)/2 - 200, int32(screenHeight)/2, 20, rl.Gray)

	countText := fmt.Sprintf("Particle Count: %d", simState.ParticleCount)
	rl.DrawText(countText, int32(screenWidth)/2 - 100, int32(screenHeight)/2 + 40, 30, rl.Yellow)

	startText := "Press ENTER to Start"
	rl.DrawText(startText, int32(screenWidth)/2 - 150, int32(screenHeight)/2 + 100, 30, rl.Green)
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

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func abs(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}
