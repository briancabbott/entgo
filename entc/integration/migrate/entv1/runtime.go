// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"github.com/briancabbott/entgo/entc/integration/migrate/entv1/schema"
	"github.com/briancabbott/entgo/entc/integration/migrate/entv1/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescBlob is the schema descriptor for blob field.
	userDescBlob := userFields[7].Descriptor()
	// user.BlobValidator is a validator for the "blob" field. It is called by the builders before save.
	user.BlobValidator = userDescBlob.Validators[0].(func([]byte) error)
	// userDescWorkplace is the schema descriptor for workplace field.
	userDescWorkplace := userFields[10].Descriptor()
	// user.WorkplaceValidator is a validator for the "workplace" field. It is called by the builders before save.
	user.WorkplaceValidator = userDescWorkplace.Validators[0].(func(string) error)
}
