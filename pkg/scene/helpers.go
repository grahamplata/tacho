package scene

import (
	"fmt"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// DrawDebugInfo draws debug information on the screen if the debug mode is enabled.
func DrawDebugInfo(screen *ebiten.Image, debugMode bool, id string) {
	// Optionally, you can draw additional UI elements or debug information here
	if debugMode {
		// Draw debug information, such as FPS or other game state
		tps := ebiten.ActualTPS()

		// Get memory stats
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		// Convert memory usage to MB
		memMB := float64(m.Alloc) / 1024 / 1024

		// Get number of goroutines (as a proxy for CPU usage)
		numGoroutine := runtime.NumGoroutine()

		// Display the debug information on the screen
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f\nMem: %.2f MB\nGoroutines: %d\n\nScene: %s", tps, memMB, numGoroutine, id))
	}
}
