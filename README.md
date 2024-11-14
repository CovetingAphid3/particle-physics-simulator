
# Particle Physics Simulator

This is a particle physics simulation that models the interaction of charged particles using Coulomb's Law for electrostatic forces. The simulation visualizes how particles with different charges move based on their interactions with each other and obstacles in a 3D space. The project also includes features like dynamic particle creation, obstacle handling, and allows users to experiment with different physics setups for a wide variety of interactions, such as slingshot effects and particle tunneling.

## Features

- **Electrostatic Forces**: Particles interact with each other via Coulomb's Law, which governs the attraction or repulsion between charged objects.
- **Particle Motion**: Particle positions, velocities, and accelerations are updated based on applied forces, including electrostatic and external forces.
- **Obstacle Interaction**: Obstacles can be placed in the simulation to influence particle motion, creating effects like collisions and deflections.
- **Charge-Based Behavior**: Particles with positive and negative charges attract or repel each other, depending on the sign of their charge.
- **Interactive Simulation**: Users can add and remove particles dynamically, making the simulation interactive and engaging.
- **Multiple Physics Options**: Easily toggle between different physics configurations, such as gravity, friction, and magnetic forces, by modifying a configuration file.
- **Dynamic Control**: Users can adjust charges, velocities, and obstacle placement to experiment with different simulation scenarios.

## Getting Started

To run this simulation, you’ll need to have Go installed. This project assumes basic knowledge of Go and physics simulation principles.

### Prerequisites

- Go version 1.18+ installed.
- Clone the repository to your local machine.

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/CovetingAphid/particle-physics-simulator.git
   cd particle-physics-simulator
   ```

2. Install the dependencies:

   ```bash
   go mod tidy
   ```

### Running the Simulation

To start the simulation, run the `main` function, which initializes the particles, applies forces, and starts the simulation loop.

**NOTE**: It may take some time to run the first time as it builds the dependencies.

```bash
go run cmd/main.go
```
**OR** Build and run

```bash
go build cmd/main.go
./main
```


This will launch the simulation with particles initialized at predefined positions, velocities, and charges. The particles will interact based on Coulomb’s Law, and you will see them move according to the forces applied.

### Customization

You can modify the following parameters in the `main.go` file to adjust the simulation:

- **Particle Charges**: Change the charge value of particles to see how they interact differently.
- **Initial Velocities**: Modify the initial velocities of particles to create various motion effects.
- **Obstacle Placement**: Add or modify obstacles in the simulation to influence particle movement.

Example:

```go
// Define particles with different initial positions and velocities
p1 := particle.NewParticle(100, 500, 0, 250, -150, 0.0, 0.0, 0.0, 0.0, 30.0, 25, color1, true)
//OR
// Define a new particle with custom charge and velocity
p1 := particle.NewCoulombParticle(100, 500, 0, 25, -15, 0.0, 0.0, 0.0, 0.0, 30.0, 25, color1, 0.010, true)
```

### Physics Options

The simulation allows you to toggle between different physics options by modifying the `physics.go` file:

- **Gravity**: Apply gravitational forces to simulate orbital mechanics and planetary interactions.
- **Magnetic Forces**: Introduce magnetic fields and interactions to model the behavior of charged particles in magnetic environments.
- **Friction**: Simulate resistance due to air or other forces acting against particle motion.

To switch between physics configurations, simply comment or uncomment the desired physics settings in the `physics.go` file.

## Future Enhancements

- **Relativistic Effects**: Implement relativistic calculations for particles moving at high velocities.
- **Graphical Visualization**: Enhance the visual representation of particles using a GUI or 3D visualization tool for better interaction.
- **Advanced Particle Interaction**: Explore other types of interactions, such as gravitational or quantum effects, for more complex simulations.

## Contributing

If you’d like to contribute to the project, feel free to fork the repository and submit a pull request. Contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

