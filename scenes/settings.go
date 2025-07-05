package scenes

import (
	"github.com/grahamplata/tacho/pkg/scene"
	"github.com/grahamplata/tacho/pkg/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// SettingsScene is a simple implementation of a Scene that represents a loading screen.
type SettingsScene struct {
	id         string
	gameState  *state.GameState
	controller scene.SceneController[*state.GameState]
}

// Update updates the settings screen scene. Press 'Escape' to return to Home.
func (l *SettingsScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		l.controller.SwitchTo(NewHomeScreenScene())
	}
	return nil
}

// NewSettingsScene creates a new SettingsScene instance.
func NewSettingsScene() *SettingsScene {
	return &SettingsScene{
		id: "Settings",
	}
}

// Layout returns the size of the settings screen scene.
func (l *SettingsScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Draw renders the loading screen scene to the provided image.
func (l *SettingsScene) Draw(screen *ebiten.Image) {
	if l.gameState != nil && l.gameState.DebugMode {
		scene.DrawDebugInfo(screen, l.gameState.DebugMode, "Settings")
	}
}

// OnEnter is called when the loading screen scene becomes active.
func (l *SettingsScene) OnEnter(s *state.GameState, controller scene.SceneController[*state.GameState]) {
	l.gameState = s
	l.controller = controller
}

// OnExit is called when the loading screen scene is deactivated.
func (l *SettingsScene) OnExit() *state.GameState {
	return l.gameState
}
