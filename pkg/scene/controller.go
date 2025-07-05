package scene

import "github.com/hajimehoshi/ebiten/v2"

// Controller is a struct that manages the current scene in the game.
type Controller[T any] struct {
	current ProtoScene[T]
}

// SwitchTo switches to a new scene without a transition
func (c *Controller[T]) SwitchTo(scene Scene[T]) {
	if prevTransition, ok := c.current.(SceneTransition[T]); ok {
		prevTransition.End()
	}
	if cur, ok := c.current.(Scene[T]); ok {
		scene.OnEnter(cur.OnExit(), c)
		c.current = scene
	}
}

// NewController creates a new Controller instance with the given scene and initial state
func NewController[T any](scene Scene[T], state T) *Controller[T] {
	c := &Controller[T]{current: scene}
	scene.OnEnter(state, c)
	return c
}

// SwitchWithTransition switches to a new scene with a transition
func (c *Controller[T]) SwitchWithTransition(scene Scene[T], transition SceneTransition[T]) {
	if prevTransition, ok := c.current.(SceneTransition[T]); ok {
		prevTransition.End()
	}
	sc := c.current.(Scene[T])
	transition.Start(sc, scene, c)
	if cur, ok := sc.(TransitionAwareScene[T]); ok {
		scene.OnEnter(cur.PreTransition(scene), c)
	} else {
		scene.OnEnter(sc.OnExit(), c)
	}
	c.current = transition
}

// ReturnFromTransition is called to return from a transition scene to the original scene.
func (c *Controller[T]) ReturnFromTransition(scene, origin Scene[T]) {
	if cur, ok := scene.(TransitionAwareScene[T]); ok {
		cur.PostTransition(origin.OnExit(), origin)
	} else {
		scene.OnEnter(origin.OnExit(), c)
	}
	c.current = scene
}

// Update updates the current scene.
func (c *Controller[T]) Update() error {
	return c.current.Update()
}

// Draw renders the current scene to the provided screen image.
func (c *Controller[T]) Draw(screen *ebiten.Image) {
	c.current.Draw(screen)
}

// Layout returns the size of the current scene.
func (c *Controller[T]) Layout(w, h int) (int, int) {
	return c.current.Layout(w, h)
}
