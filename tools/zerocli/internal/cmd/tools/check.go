// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

package tools

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

type requiredTool struct {
	name        string
	description string
	exec        string
}

var requiredTools = []*requiredTool{
	{"Rust compiler", "Needed for compilation of code inside rust/ folder", "rustc"},
	{"Cargo", "Package management inside rust/", "cargo"},
	{"CMake", "Package management for C/C++", "cmake"},
	{"Xorriso", "Needed for .ISO image generation", "xorriso"},
}

func Check() *cobra.Command {
	check := &cobra.Command{
		Use:   "check",
		Short: "Check if required commands are installed",
	}
	check.AddCommand(CheckPath())
	check.AddCommand(CheckList())
	return check
}

func CheckPath() *cobra.Command {
	path := &cobra.Command{
		Use:   "path",
		Short: "Check the installed commands in PATH environment variable",
		Run: func(cmd *cobra.Command, args []string) {
			failedCount := 0
			ok := true
			for _, req := range requiredTools {
				fmt.Printf("-- %s (as executable: %s) ... ", req.name, req.exec)
				if !req.inPath() {
					ok = false
					failedCount++
					fmt.Println("NOT FOUND")
				} else {
					fmt.Println("DONE")
				}
			}

			if !ok {
				fmt.Printf("Some tools are missing (count: %d)", failedCount)
			} else {
				fmt.Println("All tools are found.")
			}
		},
	}
	return path
}

func CheckList() *cobra.Command {
	list := &cobra.Command{
		Use:   "list",
		Short: "List required tools, what need to be installed",
		Run: func(cmd *cobra.Command, args []string) {
			for _, req := range requiredTools {
				fmt.Printf("-- %s (as executable: %s): %s\n", req.name, req.exec, req.description)
			}
		},
	}
	return list
}

func (r *requiredTool) inPath() bool {
	_, err := exec.LookPath(r.exec)
	return err == nil
}
