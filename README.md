## Particle Physics Simulator

Interactive 2D particle simulation written in Go using Raylib. It models basic mechanics, collisions, and optional electromagnetic interactions between charged particles, rendered in real time at 120 FPS.

Use this README as your presentation guide: it explains what the project does, how to run it, how to use it live, and how the code is structured.

## Highlights

- **Real‑time simulation**: 120 FPS loop with deterministic time step.
- **Interactive controls**: Pause/resume, add/remove particles with the mouse.
- **Elastic collisions**: Pairwise collision prediction and response.
- **Boundaries & ground**: Damping, friction, and ground contact handling.
- **Electrostatics & magnetism (optional)**: Coulomb forces and uniform/non‑uniform B‑fields are implemented and can be enabled.

## Demo controls (for the live presentation)

- Space: Pause/Resume
- Left click: Add a particle at the cursor
- Right click: Remove the nearest particle to the cursor
- Top‑right buttons: Close, maximize/restore, minimize (custom window controls)

When paused, particles freeze; you can still add/remove particles.

## Requirements

- Go 1.23+
- OS: Linux/macOS/Windows
- Raylib Go bindings are fetched via `go mod`; first build may take longer.

Tip (Linux): ensure common build tools are installed. For some distros you may need X11/ALSA dev packages for Raylib.

## Setup

```bash
git clone https://github.com/CovetingAphid/particle-physics-simulator.git
cd particle-physics-simulator
go mod tidy
```

## Run

```bash
go run cmd/main.go
```

Or build a binary:

```bash
go build -o simulator cmd/main.go
./simulator
```

You should see 200 colorful particles and one immovable obstacle; particles bounce within the window and collide.

## How it works (architecture)

- `cmd/main.go`: Entry point. Creates initial particles (including a fixed obstacle) and starts the simulation.
- `internal/simulation/`: The main loop and input handling.
  - `simulation.go`: Frame loop, parallel velocity/position updates, collision checks, rendering orchestration.
  - `controls.go`: Space/LMB/RMB handling, particle creation/removal.
- `internal/physics/`: Core kinematics and global forces.
  - Gravity, air drag, friction, boundary constraints, helpers to apply electrostatic/magnetic forces.
- `internal/collisions/`: Collision prediction (`WillCollide`) and elastic response (`HandleCollision`).
- `internal/electrostatics/`: Coulomb force magnitude and vector calculations; batch helpers.
- `internal/force/`: Force utilities, including gravitational/electrostatic parallelism and magnetic field models.
- `internal/particle/`: `Particle` and `Color` types and constructors.
- `internal/renderer/`: Raylib window, drawing of particles/UI, and custom window buttons.
- `internal/constants/`: Tunable physics constants (gravity, damping, friction, etc.).

## Configuration knobs (what to show in a demo)

- Toggle forces in `internal/simulation/simulation.go` by uncommenting the electrostatic/magnetic calls inside the update block.
- Change gravity, drag, friction, damping in `internal/constants/constants.go`.
- Adjust timestep and FPS in `internal/simulation/simulation.go` and `internal/renderer/renderer.go`.
- Window size and UI text are in `internal/renderer/renderer.go`.
- Initial particle count, distribution, colors, and the fixed obstacle are in `cmd/main.go`.

## Extending ideas

- Enable full electrostatics each frame via `physics.ApplyElectrostaticForces(particles)` and integrate `Fx,Fy` into acceleration.
- Use `force.ApplyForcesParallel` for scalable pairwise force computation.
- Experiment with `force.MagneticField` or `MagneticField2D` for Lenz/Lorentz‑style motion.

## Testing

Unit tests exist for collisions, electrostatics, forces, particles, and physics. Run:

```bash
go test ./...
```

## License

MIT — see `LICENSE`.

## Collision prediction and response

We separate collision prediction from response to keep the frame update simple and robust.

1) Predict potential collisions (cheap, no square roots)

- Overlap test (`CheckCollision`): compares squared center distance with squared radii sum to avoid `sqrt`.
- Next‑step prediction (`WillCollide`): projects each particle forward one time step using its current velocity, then compares squared distance at t + dt against squared radii sum. This prevents tunneling in fast‑moving cases and reduces unnecessary response calls.

2) Resolve collisions (impulse, elastic)

- Compute normalized collision normal n from current positions.
- Relative velocity v_rel along n determines whether particles are separating; if `dot(v_rel, n) >= 0` we skip.
- Impulse magnitude for an elastic collision: `j = 2 * dot(v_rel, n) / (m1 + m2)`.
- Velocities are updated along n with ±j scaled by the other mass.
- Special case: if one particle is immovable (`Movable == false`), only the movable particle’s velocity is updated.

3) Positional stability and boundaries

- Boundary constraints with damping are applied during rendering (`renderer.DrawParticle` calls `physics.ApplyBoundaryConditions`). Small velocities near the ground are zeroed using a threshold to avoid jitter.

File references

- Detection/response: `internal/collisions/collisions.go`
- Boundary/ground handling: `internal/renderer/renderer.go`, `internal/physics/physics.go`

