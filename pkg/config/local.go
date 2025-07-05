package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// LocalConfig is a struct that implements the Config interface for local configuration settings.
type LocalConfig struct {
	// Path to the configuration file or storage location.
	Path string
	// Format of the configuration file (e.g., JSON, YAML).
	Format string
}

// NewLocalConfig creates a new LocalConfig instance with the specified path and format.
func NewLocalConfig(path, format string) *LocalConfig {
	return &LocalConfig{
		Path:   path,
		Format: format,
	}
}

// Load loads the configuration settings from a local file or storage into dest.
func (c *LocalConfig) Load(dest any) error {
	file, err := os.Open(c.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	switch c.Format {
	case "json":
		if err := json.Unmarshal(bytes, dest); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported config format: %s", c.Format)
	}
	return nil
}

// Save saves the configuration settings from src to a local file or storage.
func (c *LocalConfig) Save(src any) error {
	var data []byte
	var err error
	switch c.Format {
	case "json":
		data, err = json.MarshalIndent(src, "", "  ")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported config format: %s", c.Format)
	}

	tempFile, err := os.CreateTemp("", "config-*.tmp")
	if err != nil {
		return err
	}
	tempPath := tempFile.Name()
	defer func() {
		tempFile.Close()
		os.Remove(tempPath)
	}()

	if _, err := tempFile.Write(data); err != nil {
		return err
	}
	if err := tempFile.Sync(); err != nil {
		return err
	}
	if err := tempFile.Close(); err != nil {
		return err
	}

	if err := os.Rename(tempPath, c.Path); err != nil {
		return err
	}
	return nil
}
