package main

import (
	"context"
	"fmt"
)

type Mkdocs struct{}

// example usage: "dagger call grep-dir --directory-arg . --pattern GrepDir"
func (m *Mkdocs) Build(ctx context.Context, directoryArg *Directory) *Directory {
	return dag.Container().
		From("squidfunk/mkdocs-material:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"build"}).
		Directory("/mnt/site")
}

// This won't work as the directory is not updated in the container when you change it on the host.
// See https://discord.com/channels/707636530424053791/1120503349599543376/1133802731442806885
func (m *Mkdocs) Serve(ctx context.Context, directoryArg *Directory) *Service {
	return dag.Container().
		From("squidfunk/mkdocs-material:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExposedPort(8000).
		WithExec([]string{"serve", "--dev-addr", fmt.Sprintf("0.0.0.0:%d", 8000)}).
		AsService()
}
