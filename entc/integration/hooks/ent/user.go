// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/entc/integration/hooks/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Worth holds the value of the "worth" field.
	Worth uint `json:"worth,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges            UserEdges `json:"edges"`
	user_best_friend *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// BestFriend holds the value of the best_friend edge.
	BestFriend *User `json:"best_friend,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// BestFriendOrErr returns the BestFriend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) BestFriendOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.BestFriend == nil {
			// The edge best_friend was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.BestFriend, nil
	}
	return nil, &NotLoadedError{edge: "best_friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldVersion, user.FieldWorth:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // user_best_friend
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				u.Version = int(value.Int64)
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldWorth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field worth", values[i])
			} else if value.Valid {
				u.Worth = uint(value.Int64)
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_best_friend", value)
			} else if value.Valid {
				u.user_best_friend = new(int)
				*u.user_best_friend = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the User entity.
func (u *User) QueryCards() *CardQuery {
	return (&UserClient{config: u.config}).QueryCards(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{config: u.config}).QueryFriends(u)
}

// QueryBestFriend queries the "best_friend" edge of the User entity.
func (u *User) QueryBestFriend() *UserQuery {
	return (&UserClient{config: u.config}).QueryBestFriend(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", version=")
	builder.WriteString(fmt.Sprintf("%v", u.Version))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", worth=")
	builder.WriteString(fmt.Sprintf("%v", u.Worth))
	builder.WriteString(", password=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
