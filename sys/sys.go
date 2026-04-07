// Package sys contains common application functions
package sys

import (
	"fmt"
)

type Env = string

const (
	EnvDev    Env    = "dev"
	EnvProd   Env    = "prod"
	okMessage string = "OK"
)

type Initializer[A any] struct {
	Fn   func(*A) error
	Name string
}

// IsValidEnv checks if env is valid ('dev' or 'prod')
func IsValidEnv(env Env) error {
	if env != EnvDev && env != EnvProd {
		return fmt.Errorf("invalid app env")
	}
	return nil
}

// IsProdEnv checks if env is 'prod'
func IsProdEnv(env Env) bool {
	return env == EnvProd
}

// RunInitializers runs all given initializers
func RunInitializers[A any](initializers []Initializer[A], app *A) error {
	for _, initializer := range initializers {
		err := initializer.Fn(app)
		if err != nil {
			return fmt.Errorf("%s: failed to initialize: %w", initializer.Name, err)
		}
	}
	return nil
}
