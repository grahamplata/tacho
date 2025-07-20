package state

// GameSettings are settings that can be configured for the game via the settings menu.
type GameSettings struct {
	// TargetFPS defines the target frames per second for the game.
	TargetFPS int
	// Vsync enables or disables vertical synchronization.
	Vsync bool
	// Fullscreen enables or disables fullscreen mode.
	Fullscreen bool
}

// DebugSettings are settings that can be configured for the game.
type DebugSettings struct {
	// Enabled enables or disables the display of debug information.
	Enabled bool
	// ShowTPS enables or disables the display of ticks per second.
	ShowTPS bool
	// ShowFPS enables or disables the display of frames per second.
	ShowFPS bool
	// ShowSceneName enables or disables the display of the current scene name.
	ShowSceneName bool
	// Show Goroutines enables or disables the display of the number of goroutines.
	ShowGoroutines bool
}

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
		s.GameSettings.TargetFPS = fps
	}
}

// WithVsync sets whether vertical synchronization is enabled.
func WithVsync(vsync bool) Option {
	return func(s *GameState) {
		s.GameSettings.Vsync = vsync
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
		s.GameSettings.Fullscreen = fullscreen
	}
}

// WithDebugMode sets whether debug mode is enabled.
func WithDebugMode(debugMode bool) Option {
	return func(s *GameState) {
		s.DebugSettings.ShowTPS = debugMode
	}
}
