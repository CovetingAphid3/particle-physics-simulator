package state

type MouseMode int

const (
	MouseModeAdd MouseMode = iota
	MouseModeRemove
	MouseModeAttract
	MouseModeRepel
)

func (m MouseMode) String() string {
	switch m {
	case MouseModeAdd:
		return "Add"
	case MouseModeRemove:
		return "Remove"
	case MouseModeAttract:
		return "Attract"
	case MouseModeRepel:
		return "Repel"
	default:
		return "Unknown"
	}
}

type AppState int

const (
	AppStateMenu AppState = iota
	AppStateRunning
)

type SpawnType int

const (
	SpawnTypeParticle SpawnType = iota
	SpawnTypeWall
	SpawnTypeSpring
	SpawnTypeBumper
)

func (s SpawnType) String() string {
	switch s {
	case SpawnTypeParticle:
		return "Particle"
	case SpawnTypeWall:
		return "Wall"
	case SpawnTypeSpring:
		return "Spring"
	case SpawnTypeBumper:
		return "Bumper"
	default:
		return "Unknown"
	}
}

type SimulationState struct {
	AppState             AppState
	Paused               bool
	GravityEnabled       bool
	ElectrostaticsEnabled bool
	GravityStrength      float64
	MouseMode            MouseMode
	AttractionStrength   float64
	ParticleCount        int
	SpawnType            SpawnType
	ParticleSize         float64
	SpawnMovable         bool
	IsDragging           bool
	DragStart            Vector2
}

type Vector2 struct {
	X, Y float64
}
