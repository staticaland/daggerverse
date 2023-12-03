package main

import (
	"context"
)

type Mkdocs struct{}

// example usage: "dagger call grep-dir --directory-arg . --pattern GrepDir"
func (m *Mkdocs) Build(ctx context.Context, directoryArg *Directory) (string, error) {
	return dag.Container().
		From("squidfunk/mkdocs-material:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"build"}).
		Stdout(ctx)
}
