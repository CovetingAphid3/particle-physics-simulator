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

var (
	// UI Colors
	colBackground = rl.Color{R: 20, G: 20, B: 30, A: 240}
	colPanel      = rl.Color{R: 30, G: 30, B: 40, A: 200}
	colAccent     = rl.Color{R: 0, G: 200, B: 255, A: 255} // Cyan
	colText       = rl.Color{R: 220, G: 220, B: 220, A: 255}
	colActive     = rl.Color{R: 0, G: 255, B: 100, A: 255} // Green
	colDanger     = rl.Color{R: 255, G: 50, B: 50, A: 255} // Red
)

func DrawPanel(x, y, w, h int32, title string) {
	rl.DrawRectangle(x, y, w, h, colPanel)
	rl.DrawRectangleLines(x, y, w, h, colAccent)
	if title != "" {
		rl.DrawRectangle(x, y, w, 30, colAccent)
		rl.DrawText(title, x+10, y+5, 20, rl.Black)
	}
}

func DrawProgressBar(x, y, w, h int32, value, max float32, label string) {
	rl.DrawText(label, x, y-20, 15, colText)
	rl.DrawRectangle(x, y, w, h, rl.Black)
	
	fillW := int32((value / max) * float32(w))
	if fillW > w { fillW = w }
	if fillW < 0 { fillW = 0 }
	
	rl.DrawRectangle(x, y, fillW, h, colAccent)
	rl.DrawRectangleLines(x, y, w, h, rl.Gray)
}

func DrawButton(x, y, w, h int32, text string, active bool) {
	color := colPanel
	textColor := colText
	if active {
		color = colActive
		textColor = rl.Black
	}
	rl.DrawRectangle(x, y, w, h, color)
	rl.DrawRectangleLines(x, y, w, h, colAccent)
	rl.DrawText(text, x+10, y+5, 20, textColor)
}

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

	// Draw Dragging Visuals
	if simState.IsDragging {
		mouseX := int32(rl.GetMouseX())
		mouseY := int32(rl.GetMouseY())
		startX := int32(simState.DragStart.X)
		startY := int32(simState.DragStart.Y)

		if simState.SpawnType == state.SpawnTypeWall || simState.SpawnType == state.SpawnTypeBumper {
			x := min(startX, mouseX)
			y := min(startY, mouseY)
			width := abs(mouseX - startX)
			height := abs(mouseY - startY)
			color := rl.Green
			if simState.SpawnType == state.SpawnTypeBumper { color = rl.Red }
			rl.DrawRectangleLines(x, y, width, height, color)
		} else if simState.SpawnType == state.SpawnTypeSpring {
			rl.DrawLine(startX, startY, mouseX, mouseY, rl.Green)
		}
	}

	// --- HUD Layout ---
	
	// 1. Top Left Stats
	DrawPanel(10, 10, 200, 80, "Stats")
	rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), 20, 45, 20, colText)
	rl.DrawText(fmt.Sprintf("Particles: %d", len(particles)), 20, 65, 20, colText)

	// 2. Right Control Panel
	panelW := int32(300)
	panelX := int32(screenWidth) - panelW - 10
	panelY := int32(40) // Below window buttons
	
	DrawPanel(panelX, panelY, panelW, 600, "Control Center")
	
	currentY := panelY + 40
	padding := int32(10)
	
	// Section: Simulation
	rl.DrawText("Simulation", panelX+padding, currentY, 20, colAccent)
	currentY += 30
	DrawButton(panelX+padding, currentY, 130, 30, "Pause (Spc)", simState.Paused)
	DrawButton(panelX+150, currentY, 130, 30, "Reset (R)", false)
	currentY += 40

	// Section: Spawning
	rl.DrawText("Spawn Tool (T)", panelX+padding, currentY, 20, colAccent)
	currentY += 30
	
	// Spawn Type Buttons
	types := []string{"Particle", "Wall", "Spring", "Bumper"}
	for i, t := range types {
		isActive := simState.SpawnType.String() == t
		DrawButton(panelX+padding + int32(i%2)*140, currentY + int32(i/2)*40, 130, 30, t, isActive)
	}
	currentY += 90
	
	// Movable Toggle
	DrawButton(panelX+padding, currentY, 280, 30, fmt.Sprintf("Movable (B): %v", simState.SpawnMovable), simState.SpawnMovable)
	currentY += 50
	
	// Size Slider
	DrawProgressBar(panelX+padding, currentY, 280, 20, float32(simState.ParticleSize), 50.0, fmt.Sprintf("Size (+/-): %.1f", simState.ParticleSize))
	currentY += 50

	// Section: Physics
	rl.DrawText("Physics", panelX+padding, currentY, 20, colAccent)
	currentY += 30
	
	DrawButton(panelX+padding, currentY, 130, 30, "Gravity (G)", simState.GravityEnabled)
	DrawButton(panelX+150, currentY, 130, 30, "Electro (E)", simState.ElectrostaticsEnabled)
	currentY += 50
	
	DrawProgressBar(panelX+padding, currentY, 280, 20, float32(simState.GravityStrength), 2000.0, fmt.Sprintf("Gravity Str (Arrows): %.0f", simState.GravityStrength))
	currentY += 50

	// Section: Mouse Interaction
	rl.DrawText("Interaction", panelX+padding, currentY, 20, colAccent)
	currentY += 30
	DrawButton(panelX+padding, currentY, 280, 30, fmt.Sprintf("Mode (M): %s", simState.MouseMode.String()), true)
	currentY += 50
	DrawProgressBar(panelX+padding, currentY, 280, 20, float32(simState.AttractionStrength), 10000.0, fmt.Sprintf("Force Str ([/]): %.0f", simState.AttractionStrength))

	// Bottom Help
	rl.DrawText("Press keys to toggle options. Drag to create.", 10, int32(screenHeight)-30, 20, rl.Gray)
}

func DrawMenu(simState *state.SimulationState) {
	rl.ClearBackground(colBackground)
	
	centerX := int32(screenWidth) / 2
	centerY := int32(screenHeight) / 2
	
	// Title Panel
	DrawPanel(centerX-300, centerY-200, 600, 400, "")
	
	title := "PARTICLE SIMULATOR"
	titleW := rl.MeasureText(title, 50)
	rl.DrawText(title, centerX - titleW/2, centerY - 150, 50, colAccent)

	rl.DrawText("v2.0 - Physics Engine", centerX - 100, centerY - 90, 20, colText)

	// Particle Count Control
	rl.DrawText("Initial Particle Count", centerX - 100, centerY - 20, 20, colText)
	
	countStr := fmt.Sprintf("<  %d  >", simState.ParticleCount)
	countW := rl.MeasureText(countStr, 40)
	rl.DrawText(countStr, centerX - countW/2, centerY + 20, 40, colActive)
	
	rl.DrawText("Use Up/Down Arrows", centerX - 80, centerY + 70, 15, rl.Gray)

	// Start Prompt
	startText := "Press ENTER to Start"
	startW := rl.MeasureText(startText, 30)
	
	// Blinking effect
	if int(rl.GetTime()*2)%2 == 0 {
		rl.DrawText(startText, centerX - startW/2, centerY + 130, 30, colAccent)
	}
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
