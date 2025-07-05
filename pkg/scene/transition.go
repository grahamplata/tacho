package scene

// SceneTransition[T any] is an interface that defines the methods required for a scene transition.
type SceneTransition[T any] interface {
	ProtoScene[T]
	// Start initiates the transition from one scene to another.
	Start(from, to Scene[T], sc SceneController[T])
	// End finalizes the transition and cleans up any resources.
	End()
}
