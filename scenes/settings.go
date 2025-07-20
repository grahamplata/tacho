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

// SettingsScene is a simple implementation of a Scene that represents a loading screen.
type SettingsScene struct {
	id         string
	face       *text.GoTextFace
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
		id:   "Settings",
		face: assets.SharedTitleFace,
	}
}

// Layout returns the size of the settings screen scene.
func (l *SettingsScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Draw renders the loading screen scene to the provided image.
func (l *SettingsScene) Draw(screen *ebiten.Image) {
	message := "Settings"
	centerX := float64(screen.Bounds().Dx()) / 2
	centerY := float64(screen.Bounds().Dy()) / 2

	settingsWidth, _ := text.Measure(message, l.face, 0)

	settingsOptions := &text.DrawOptions{}
	settingsOptions.GeoM.Translate(centerX-settingsWidth/2, centerY-25)

	text.Draw(screen, message, l.face, settingsOptions)

	if l.gameState != nil && l.gameState.Enabled {
		helpers.DrawDebugInfo(screen, l.gameState)
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
