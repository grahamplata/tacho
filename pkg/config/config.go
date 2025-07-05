package config

// Config is an interface that defines methods for loading and saving configuration settings.
type Config interface {
	// Load loads the configuration settings from a persistent storage.
	Load() error
	// Save saves the configuration settings to a persistent storage.
	Save() error
}

// Option is a functional option type for configuring the Config.
type Option func(Config)

// WithOption applies a functional option to the Config.
func WithOption(c Config, opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}
