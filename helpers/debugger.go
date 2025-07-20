package helpers

import (
	"fmt"
	"runtime"

	"github.com/grahamplata/tacho/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw renders the loading screen scene to the provided image.
func DrawDebugInfo(screen *ebiten.Image, state *state.GameState) {
	debugInfo := ""
	debugInfo += fmt.Sprintf("TPS: %.2f\n", ebiten.ActualTPS())
	debugInfo += fmt.Sprintf("FPS: %.2f\n", ebiten.ActualFPS())
	debugInfo += fmt.Sprintf("Scene: %s\n", state.Title)
	debugInfo += fmt.Sprintf("Goroutines: %d\n", runtime.NumGoroutine())

	ebitenutil.DebugPrint(screen, debugInfo)
}
