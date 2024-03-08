// A generated module for Cloc functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
)

type Cloc struct{}

// Count lines of code in a directory
func (m *Cloc) ClocDir(
	ctx context.Context,
	directoryArg *Directory,
	// json output
	// +optional
	jsonOutput bool,
	// yaml output
	// +optional
	yamlOutput bool,
) (string, error) {

	command := []string{"cloc", "--quiet", "--hide-rate"}

	if jsonOutput {
		command = append(command, "--json")
	}

	if yamlOutput {
		command = append(command, "--yaml")
	}

	command = append(command, ".")

	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "cloc"}).
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec(command).
		Stdout(ctx)
}
