// A generated module for Renovate functions

package main

import (
	"context"

	"dagger.io/dagger/dag"
)

type Renovate struct{}

// Runs Renovate
func (m *Renovate) Run(ctx context.Context) (string, error) {
	return dag.Container().
		From("renovate/renovate:latest").
		WithExec([]string{"renovate", "--help"}).
		Stdout(ctx)
}
