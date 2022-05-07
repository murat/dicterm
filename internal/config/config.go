package config

import (
	"fmt"
	"io"
	"os"
)

// FileName is the default configuration file name
const FileName = ".dicterm"

var _ io.ReadWriteCloser = &Config{}

// Config is configuration file
type Config struct {
	File *os.File
}

// var _ io.ReadWriter = &Config{}

// New returns configuration file
func New(path string) (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get user home, %w", err)
	}
	if path == "" {
		path = fmt.Sprintf("%s/%s", home, FileName)
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not open config file, %w", err)
	}

	return &Config{File: file}, nil
}

// Read reads key from config file
func (cfg *Config) Read(p []byte) (int, error) {
	return cfg.File.Read(p)
}

// Write writes key to config file
func (cfg *Config) Write(p []byte) (int, error) {
	if err := cfg.File.Truncate(0); err != nil {
		return 0, fmt.Errorf("could not truncate file, %w", err)
	}
	if _, err := cfg.File.Seek(0, 0); err != nil {
		return 0, fmt.Errorf("could not seek file, %w", err)
	}

	n, err := cfg.File.WriteString(string(p))
	if err != nil {
		return 0, fmt.Errorf("could not write file, %w", err)
	}

	return n, nil
}

func (cfg *Config) Close() error {
	return cfg.File.Close()
}
