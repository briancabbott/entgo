// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	ent "github.com/briancabbott/entgo"
	"github.com/briancabbott/entgo/dialect"
	"github.com/briancabbott/entgo/entc/integration/multischema/ent/internal"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks

	schemaConfig SchemaConfig
}

// hooks per client, for fast access.
type hooks struct {
	Group []ent.Hook
	Pet   []ent.Hook
	User  []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// SchemaConfigFromContext exports the internal.SchemaConfigFromContext
// for external usage (inside custom predicates or modifiers).
func SchemaConfigFromContext(ctx context.Context) SchemaConfig {
	return internal.SchemaConfigFromContext(ctx)
}

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig = internal.SchemaConfig

// AlternateSchemas allows alternate schema names to be
// passed into ent operations.
func AlternateSchema(schemaConfig SchemaConfig) Option {
	return func(c *config) {
		c.schemaConfig = schemaConfig
	}
}
