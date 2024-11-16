package constants

// Physics Constants

// Gravitational constant (G) in m³/(kg·s²)
const GravitationalConstant = 6.67430e-11 // m³/(kg·s²)

// Coulomb's constant (k_e) in N·m²/C²
const CoulombsConstant = 8.9875517923e9 // N·m²/C²

// Permittivity of free space (ε₀) in C²/(N·m²)
const PermittivityOfFreeSpace = 8.854187817e-12 // C²/(N·m²)

// Elementary charge (e) in Coulombs
const ElectronCharge = 1.602176634e-19 // C

// SI unit conversion constants
const MeterToPixelConversion = 100.0 // Conversion from meters to pixels for rendering
const SecondsPerFrame = 1.0 / 120.0 // Simulation frame rate (120 FPS)

// Particle Constants

// Default mass and radius for particles
const DefaultMass = 1.0 // kg
const DefaultRadius = 10.0 // pixels

// Friction Constants

// Coefficient of kinetic friction for ground surfaces
const GroundFrictionCoefficient = 0.0001

// Air resistance coefficient (drag coefficient, dimensionless)
const AirDragCoefficient = 0.00009

// Increased gravity in pixels/second²
const Gravity = 980.0

// Damping factor for bounces
const DampingFactor = 0.7

// Velocity threshold for stopping
const VelocityThreshold = 20.0

const CoefficientOfRestitution = 0.8 // Typical value for elastic collisions

