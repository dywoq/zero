// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

package cmd

import (
	"io"

	"github.com/BurntSushi/toml"
)

// Flag represents the flag of the command.
type Flag struct {
	Name        string `toml:"name"`
	Short       string `toml:"short"`
	Description string `toml:"description"`
}

// Config represents the configuration of the command.
type Config struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	Flags       []Flag `toml:"flags"`
	Script      string `toml:"script"`
}

// ConfigParser represents the parser of the command configuration files.
type ConfigParser struct{}

// Parse parses .TOML config from r, decoding.
// Returns an error if it fails.
func (p *ConfigParser) Parse(r io.Reader) (*Config, error) {
	var config Config
	_, err := toml.NewDecoder(r).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
