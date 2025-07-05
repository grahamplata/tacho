package scenes

import (
	"github.com/grahamplata/wow-kewl/pkg/scene"
	"github.com/grahamplata/wow-kewl/pkg/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// HomeScreenScene is a simple implementation of a Scene that represents the start screen.
type HomeScreenScene struct {
	id         string
	gameState  *state.GameState
	controller scene.SceneController[*state.GameState]
}

// NewHomeScreenScene creates a new HomeScreenScene instance.
func NewHomeScreenScene() *HomeScreenScene {
	return &HomeScreenScene{
		id: "Home",
	}
}

// Update updates the home screen scene. Press 'S' to go to settings.
func (l *HomeScreenScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		l.controller.SwitchTo(NewSettingsScene())
	}
	return nil
}

// Layout returns the size of the home screen scene.
func (l *HomeScreenScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Draw renders the home screen scene to the provided image.
func (l *HomeScreenScene) Draw(screen *ebiten.Image) {
	if l.gameState != nil && l.gameState.DebugMode {
		scene.DrawDebugInfo(screen, l.gameState.DebugMode, l.id)
	}
}

// OnEnter is called when the home screen scene becomes active.
func (l *HomeScreenScene) OnEnter(s *state.GameState, controller scene.SceneController[*state.GameState]) {
	l.gameState = s
	l.controller = controller
}

// OnExit is called when the home screen scene is deactivated.
func (l *HomeScreenScene) OnExit() *state.GameState {
	return l.gameState
}
