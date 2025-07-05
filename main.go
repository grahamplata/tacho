package main

import (
	"flag"
	"log"

	"github.com/grahamplata/wow-kewl/pkg/config"
	"github.com/grahamplata/wow-kewl/pkg/scene"
	"github.com/grahamplata/wow-kewl/pkg/state"
	"github.com/grahamplata/wow-kewl/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Allow config path override via flag or environment variable
	configPath := flag.String("config", "", "Path to config file")
	configType := flag.String("config-type", "json", "Type of config file (json, yaml, etc.)")
	flag.Parse()

	path := "config.json"
	if *configPath != "" {
		path = *configPath
	}

	// Always initialize with defaults, then override with config if available
	gameState := state.NewGameState(
		state.WithTitle("Wow Kewl"),
		state.WithScreenSize(1280, 720),
		state.WithVsync(true),
		state.WithTargetFPS(60),
		state.WithDebugMode(false),
	)

	// Load configuration from the specified path
	err := config.NewLocalConfig(path, *configType).Load(gameState)
	if err != nil {
		log.Printf("Warning: could not load %s, using defaults: %v", path, err)
	}

	// Create the initial scene controller with a loading screen
	initialScene := scenes.NewLoadingScreenScene()

	// Initialize the scene controller with the game state
	controller := scene.NewController(initialScene, gameState)

	// Set Ebiten window and game options
	ebiten.SetWindowTitle(gameState.Title)
	ebiten.SetWindowSize(gameState.ScreenWidth, gameState.ScreenHeight)
	ebiten.SetVsyncEnabled(gameState.Vsync)
	ebiten.SetTPS(gameState.TargetFPS)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	defer func() {
		// Save the game state to the config file on exit
		if err := config.NewLocalConfig(path, *configType).Save(gameState); err != nil {
			log.Printf("Error saving config: %v", err)
		}
	}()

	// Run the game
	if err := ebiten.RunGame(controller); err != nil {
		log.Fatal(err)
	}
}
