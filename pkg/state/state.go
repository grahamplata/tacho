package state

// GameState holds shared game state that scenes might need access to.
type GameState struct {
	// ID string // Unique identifier for the game state, can be generated or set externally
	ID string
	// ScreenWidth and ScreenHeight define the dimensions of the game window.
	ScreenWidth int
	// ScreenHeight defines the height of the game window.
	ScreenHeight int
	// TargetFPS defines the target frames per second for the game.
	TargetFPS int
	// Vsync enables or disables vertical synchronization.
	Vsync bool
	// Title defines the title of the game.
	Title string
	// Fullscreen
	Fullscreen bool
	// DebugMode enables or disables debug information display.
	DebugMode bool
}

// Option is a functional option type for configuring the GameState.
type Option func(*GameState)

// DisplayOptions returns a slice of functional options that can be used to configure the GameState.
func WithScreenSize(width, height int) Option {
	return func(s *GameState) {
		s.ScreenWidth = width
		s.ScreenHeight = height
	}
}

// WithTargetFPS sets the target frames per second for the game.
func WithTargetFPS(fps int) Option {
	return func(s *GameState) {
		s.TargetFPS = fps
	}
}

// WithVsync sets whether vertical synchronization is enabled.
func WithVsync(vsync bool) Option {
	return func(s *GameState) {
		s.Vsync = vsync
	}
}

// WithTitle sets the title of the game.
func WithTitle(title string) Option {
	return func(s *GameState) {
		s.Title = title
	}
}

// WithFullscreen sets whether the game should run in fullscreen mode.
func WithFullscreen(fullscreen bool) Option {
	return func(s *GameState) {
		s.Fullscreen = fullscreen
	}
}

// WithDebugMode sets whether debug mode is enabled.
func WithDebugMode(debugMode bool) Option {
	return func(s *GameState) {
		s.DebugMode = debugMode
	}
}

// NewGameState creates a new GameState with the provided options.
func NewGameState(opts ...Option) *GameState {
	// Default game state values
	// These can be overridden by the provided options.
	state := &GameState{
		ID:           "default",
		ScreenWidth:  800,
		ScreenHeight: 600,
		TargetFPS:    60,
		Vsync:        true,
		Title:        "Wow Kewl Game",
		Fullscreen:   false,
		DebugMode:    false,
	}

	for _, opt := range opts {
		opt(state)
	}

	return state
}
