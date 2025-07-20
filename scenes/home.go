package scenes

import (
	"github.com/grahamplata/tacho/assets"
	"github.com/grahamplata/tacho/helpers"
	"github.com/grahamplata/tacho/pkg/scene"
	"github.com/grahamplata/tacho/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// HomeScreenScene is a simple implementation of a Scene that represents the start screen.
type HomeScreenScene struct {
	id         string
	face       *text.GoTextFace
	gameState  *state.GameState
	controller scene.SceneController[*state.GameState]
}

// NewHomeScreenScene creates a new HomeScreenScene instance.
func NewHomeScreenScene() *HomeScreenScene {
	return &HomeScreenScene{
		id:   "Home",
		face: assets.SharedTitleFace,
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
	message := "Welcome"
	centerX := float64(screen.Bounds().Dx()) / 2
	centerY := float64(screen.Bounds().Dy()) / 2

	welcomeWidth, _ := text.Measure(message, l.face, 0)

	welcomeOptions := &text.DrawOptions{}
	welcomeOptions.GeoM.Translate(centerX-welcomeWidth/2, centerY-25)

	// Draw the title and subtitle on the screen
	text.Draw(screen, message, l.face, welcomeOptions)

	if l.gameState != nil && l.gameState.Enabled {
		helpers.DrawDebugInfo(screen, l.gameState)
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
