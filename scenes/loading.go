package scenes

import (
	"github.com/grahamplata/tacho/assets"
	"github.com/grahamplata/tacho/pkg/scene"
	"github.com/grahamplata/tacho/pkg/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// LoadingScreenScene represents a loading screen scene.
type LoadingScreenScene struct {
	id         string
	title      string
	subtitle   string
	face       *text.GoTextFace
	state      *state.GameState
	controller scene.SceneController[*state.GameState]
	// Cached measured widths for efficiency
	titleWidth    float64
	subtitleWidth float64
	// Dirty flags to track if text has changed
	titleDirty    bool
	subtitleDirty bool
}

// NewLoadingScreenScene creates a new LoadingScreenScene with the given title and subtitle.
func NewLoadingScreenScene() *LoadingScreenScene {
	title, subtitle := "Loading", "Please wait..."
	// Use the shared face for both title and subtitle (do not mutate global state)
	return &LoadingScreenScene{
		id:            "Loading",
		title:         title,
		subtitle:      subtitle,
		face:          assets.SharedTitleFace,
		titleDirty:    true,
		subtitleDirty: true,
	}
}

// Update updates the loading screen scene.
func (l *LoadingScreenScene) Update() error {
	// Show prompt to continue
	if l.subtitle != "Press Space to continue" {
		l.subtitle = "Press Space to continue"
		l.subtitleDirty = true
	}
	// Switch to home scene on space press
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if l.title != "Loading Complete" {
			l.title = "Loading Complete"
			l.titleDirty = true
		}
		l.controller.SwitchTo(NewHomeScreenScene())
	}
	return nil
}

// Layout returns the size of the loading screen scene. It returns a fixed size.
func (l *LoadingScreenScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Draw renders the loading screen scene to the provided image.
func (l *LoadingScreenScene) Draw(screen *ebiten.Image) {
	centerX := float64(screen.Bounds().Dx()) / 2
	centerY := float64(screen.Bounds().Dy()) / 2

	// Only measure text widths if dirty
	if l.titleDirty {
		l.titleWidth, _ = text.Measure(l.title, l.face, 0)
		l.titleDirty = false
	}
	if l.subtitleDirty {
		l.subtitleWidth, _ = text.Measure(l.subtitle, l.face, 0)
		l.subtitleDirty = false
	}

	// Title options for drawing text - center horizontally and position above center
	titleOptions := &text.DrawOptions{}
	titleOptions.GeoM.Translate(centerX-l.titleWidth/2, centerY-25)

	// Subtitle options for drawing text - center horizontally and position below center
	subtitleOptions := &text.DrawOptions{}
	subtitleOptions.GeoM.Translate(centerX-l.subtitleWidth/2, centerY+25)

	// Draw the title and subtitle on the screen
	text.Draw(screen, l.title, l.face, titleOptions)
	text.Draw(screen, l.subtitle, l.face, subtitleOptions)

	if l.state != nil && l.state.DebugMode {
		scene.DrawDebugInfo(screen, l.state.DebugMode, "Loading")
	}
}

// OnEnter is called when the loading screen scene becomes active.
func (l *LoadingScreenScene) OnEnter(s *state.GameState, controller scene.SceneController[*state.GameState]) {
	l.state = s
	l.controller = controller
}

// OnExit is called when the loading screen scene becomes deactivated.
func (l *LoadingScreenScene) OnExit() *state.GameState {
	return l.state
}
