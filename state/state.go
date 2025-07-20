package state

// GameState holds shared game state that scenes might need access to.
type GameState struct {
	// Title defines the title of the game.
	Title string
	// ScreenWidth and ScreenHeight define the dimensions of the game window.
	ScreenWidth int
	// ScreenHeight defines the height of the game window.
	ScreenHeight int
	// GameSettings contains the settings for the game.
	GameSettings
	// DebugSettings contains the settings for debugging.
	DebugSettings
}

// Option is a functional option type for configuring the GameState.
type Option func(*GameState)

// NewGameState creates a new GameState with the provided options.
func NewGameState(opts ...Option) *GameState {
	state := &GameState{
		Title:        "Wow Kewl Game",
		ScreenWidth:  800,
		ScreenHeight: 600,
		GameSettings: GameSettings{
			TargetFPS:  60,
			Vsync:      true,
			Fullscreen: false,
		},
		DebugSettings: DebugSettings{
			ShowTPS:        false,
			ShowFPS:        false,
			ShowSceneName:  false,
			ShowGoroutines: false,
		},
	}

	for _, opt := range opts {
		opt(state)
	}

	return state
}
