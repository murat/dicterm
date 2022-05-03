package config

import (
	"fmt"
	"io"
	"os"
)

// IConfig is interface for config file
type IConfig interface {
	Read() (*string, error)
	Write(string) error
}

// Config is configuration file
type Config struct {
	File *os.File
}

// FileName is the default configuration file name
const FileName = ".dicterm"

// New returns configuration file
func New(path string) (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get user home, %w", err)
	}
	if path == "" {
		path = fmt.Sprintf("%s/%s", home, FileName)
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not open config file, %w", err)
	}

	return &Config{file}, nil
}

// Read reads key from config file
func (cfg *Config) Read() (*string, error) {
	buf := make([]byte, 1024)
	n, err := cfg.File.Read(buf)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("could not read file, %w", err)
	}

	key := string(buf[:n])
	if key == "" {
		return nil, ErrEmptyFile
	}

	return &key, nil
}

// Write writes key to config file
func (cfg *Config) Write(key string) error {
	_, err := cfg.File.Write([]byte(key))
	if err != nil {
		return fmt.Errorf("could not write to the config file, %w", err)
	}

	return nil
}
