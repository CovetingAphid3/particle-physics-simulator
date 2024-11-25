package force

import (
	"math"
	"particle-physics-simulator/internal/constants"
	"particle-physics-simulator/internal/electrostatics"
	"particle-physics-simulator/internal/particle"
	"runtime"
	"sync"
)

const (
	CutoffDistanceSquared = 1e4         
)

// Force struct represents a force with its magnitude and components (X, Y directions).
type Force struct {
	Value      float64
	XComponent float64
	YComponent float64
}

// NewForce creates a new Force object
func NewForce(value, xComponent, yComponent float64) *Force {
	return &Force{
		Value:      value,
		XComponent: xComponent,
		YComponent: yComponent,
	}
}

// Preallocated memory pool for force arrays.
var forcePool = sync.Pool{
	New: func() interface{} {
		return make([]float64, 0, 1024) // Default capacity of 1024
	},
}

// Utility functions to manage pooled arrays.
func getForceArray(size int) []float64 {
	arr := forcePool.Get().([]float64)
	if cap(arr) < size {
		return make([]float64, size) 
	}
	return arr[:size]
}

func releaseForceArray(arr []float64) {
	forcePool.Put(arr[:0]) // Reset slice for reuse
}

// ApplyForcesParallel calculates electrostatic and gravitational forces concurrently.
func ApplyForcesParallel(particles []*particle.Particle) {
	n := len(particles)
	if n == 0 {
		return
	}

	// Preallocate force arrays
	forceX := getForceArray(n)
	forceY := getForceArray(n)
	defer releaseForceArray(forceX)
	defer releaseForceArray(forceY)

	// Number of goroutines to use
	numGoroutines := runtime.NumCPU()
	chunkSize := (n + numGoroutines - 1) / numGoroutines
	var wg sync.WaitGroup

	// Parallelize force calculation
	for g := 0; g < numGoroutines; g++ {
		start := g * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				p1 := particles[i]
				if !p1.Movable {
					continue
				}

				for j := i + 1; j < n; j++ {
					p2 := particles[j]
					if !p2.Movable {
						continue
					}

					// Distance calculation
					dx := p2.X - p1.X
					dy := p2.Y - p1.Y
					distSq := dx*dx + dy*dy

					// Skip calculations if distance is negligible or exceeds cutoff
					if distSq < 1e-10 || distSq > CutoffDistanceSquared {
						continue
					}

					// Inverse distance and distance squared
					invDist := 1.0 / math.Sqrt(distSq)
					invDistSq := invDist * invDist

					// Calculate electrostatic force
					electroForce := electrostatics.CalculateElectrostaticForce(p1, p2)

					// Calculate gravitational force
					gravForce := constants.GravitationalConstant * p1.Mass * p2.Mass * invDistSq

					// Total force magnitude
					totalForce := electroForce + gravForce

					// Force components
					fx := totalForce * dx * invDist
					fy := totalForce * dy * invDist

					// Apply Newton's third law
					forceX[i] += fx
					forceY[i] += fy
					forceX[j] -= fx
					forceY[j] -= fy
				}
			}
		}(start, end)
	}

	wg.Wait()

	// Apply forces to update particle velocities
	for i, p := range particles {
		if p.Movable {
			invMass := 1.0 / p.Mass
			p.Vx += forceX[i] * invMass
			p.Vy += forceY[i] * invMass
		}
	}
}

// ApplyGravitationalForces calculates gravitational forces between all particles.
func ApplyGravitationalForces(particles []*particle.Particle) {
	for i := 0; i < len(particles); i++ {
		for j := i + 1; j < len(particles); j++ {
			p1 := particles[i]
			p2 := particles[j]

			// Distance calculation
			dx := p2.X - p1.X
			dy := p2.Y - p1.Y
			distSq := dx*dx + dy*dy

			if distSq < 1e-10 || distSq > CutoffDistanceSquared {
				continue
			}

			// Calculate gravitational force
			invDist := 1.0 / math.Sqrt(distSq)
			forceMagnitude := constants.GravitationalConstant * p1.Mass * p2.Mass * invDist * invDist

			// Force components
			fx := forceMagnitude * dx * invDist
			fy := forceMagnitude * dy * invDist

			// Apply forces (action-reaction)
			p1.Ax += fx / p1.Mass
			p1.Ay += fy / p1.Mass
			p2.Ax -= fx / p2.Mass
			p2.Ay -= fy / p2.Mass
		}
	}
}

