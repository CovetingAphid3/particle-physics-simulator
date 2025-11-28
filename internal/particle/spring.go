package particle

type Spring struct {
	P1, P2     *Particle
	RestLength float64
	Stiffness  float64
	Damping    float64
}

func NewSpring(p1, p2 *Particle, stiffness, damping float64) *Spring {

	// restLength := math.Sqrt(dx*dx + dy*dy) // Calculate rest length based on initial distance
    // For now let's assume we want them to stay at current distance
    // We need math import if we calculate it here.
    // Let's just pass restLength or calculate it in the caller if needed, 
    // or just use 0 if we want them to pull together completely (rubber band).
    // Usually springs have a rest length.
    
    return &Spring{
        P1: p1,
        P2: p2,
        RestLength: 0, // Placeholder, should be set by caller usually
        Stiffness: stiffness,
        Damping: damping,
    }
}
