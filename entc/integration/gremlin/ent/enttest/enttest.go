// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package enttest

import (
	"github.com/briancabbott/entgo/entc/integration/gremlin/ent"
	// required by schema hooks.
	_ "github.com/briancabbott/entgo/entc/integration/gremlin/ent/runtime"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...interface{})
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts []ent.Option
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...ent.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls ent.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *ent.Client {
	o := newOptions(opts)
	c, err := ent.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return c
}

// NewClient calls ent.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *ent.Client {
	o := newOptions(opts)
	c := ent.NewClient(o.opts...)
	return c
}
