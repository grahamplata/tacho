package scene

// A Scene is a part of the game that can be displayed, such as a loading screen, start screen, or settings screen.

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ProtoScene[T any] defines the base interface for a scene in the game.
type ProtoScene[T any] interface {
	ebiten.Game
}

// SceneController[T any] defines the interface for a controller that manages scene transitions.
type SceneController[T any] interface {
	ReturnFromTransition(scene, origin Scene[T])
	SwitchTo(scene Scene[T])
}

// Scene[T] is an interface that defines the methods required for a scene in the game.
type Scene[T any] interface {
	ProtoScene[T]
	// OnEnter is called when the scene becomes active.
	// controller: the scene controller managing transitions.
	OnEnter(T, SceneController[T])
	// OnExit is called when the scene is deactivated.
	// Returns the current state/context to be passed to the next scene.
	OnExit() T
}

// TransitionAwareScene[T any] extends Scene[T] to include methods for handling transitions between scenes.
type TransitionAwareScene[T any] interface {
	Scene[T]
	// PreTransition is called before the scene transitions to another scene.
	PreTransition(Scene[T]) T
	// PostTransition is called after the scene has transitioned to another scene.
	PostTransition(T, Scene[T])
}
