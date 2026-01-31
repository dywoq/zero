// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/dywoq/zero/tools/zerocli/internal/cmd"
	"github.com/dywoq/zero/tools/zerocli/internal/root"
)

// die Checks if err is not nil.
// If it's true, it prints a formatted message and exits with code 1.
func die(err error, format string, v ...any) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", fmt.Sprintf(format, v...))
		os.Exit(1)
	}
}

func main() {
	// Search for root configuration first.
	rootFile, err := os.Open("zerocli.toml")
	die(err, "Failed to open zerocli.toml configuration file: %v", err)

	rootParser := root.ConfigParser{}
	config, err := rootParser.Parse(rootFile)
	die(err, "Failed to parse zerocli.toml configuration file: %v", err)

	cmdParser := cmd.ConfigParser{}

	// Search for command configurations
	configs := []*cmd.Config{}
	err = filepath.WalkDir(config.CommandsDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(d.Name())
		if ext == ".toml" {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			config, err := cmdParser.Parse(f)
			if err != nil {
				return err
			}
			configs = append(configs, config)
		}

		return nil
	})
	die(err, "Failed to walk through \"%s\" directory: %s", config.CommandsDir, err)
}
