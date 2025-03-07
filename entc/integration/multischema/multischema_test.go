// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package multischema

import (
	"context"
	"strings"
	"testing"

	"github.com/briancabbott/entgo/dialect"
	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/entc/integration/multischema/ent"
	"github.com/briancabbott/entgo/entc/integration/multischema/ent/group"
	"github.com/briancabbott/entgo/entc/integration/multischema/ent/migrate"
	"github.com/briancabbott/entgo/entc/integration/multischema/ent/pet"
	"github.com/briancabbott/entgo/entc/integration/multischema/ent/user"
	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/")
	require.NoError(t, err)
	defer db.Close()
	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS db1")
	require.NoError(t, err, "creating database")
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS db2")
	require.NoError(t, err, "creating database")
	defer db.ExecContext(ctx, "DROP DATABASE IF EXISTS db1")
	defer db.ExecContext(ctx, "DROP DATABASE IF EXISTS db2")
	setupSchema(t, db)

	client := ent.NewClient(ent.Driver(db), ent.AlternateSchema(ent.SchemaConfig{
		Pet:        "db1",
		User:       "db2",
		Group:      "db1",
		GroupUsers: "db2",
	}))
	pedro := client.Pet.Create().SetName("Pedro").SaveX(ctx)
	groups := client.Group.CreateBulk(
		client.Group.Create().SetName("GitHub"),
		client.Group.Create().SetName("GitLab"),
	).SaveX(ctx)
	usr := client.User.Create().SetName("a8m").AddPets(pedro).AddGroups(groups...).SaveX(ctx)

	// Custom modifier with schema config.
	var names []struct {
		User string `sql:"user_name"`
		Pet  string `sql:"pet_name"`
	}
	client.Pet.Query().
		Modify(func(s *sql.Selector) {
			// The below function is exported using a custom
			// template defined in ent/template/config.tmpl.
			cfg := ent.SchemaConfigFromContext(s.Context())
			t := sql.Table(user.Table).Schema(cfg.User)
			s.Join(t).On(s.C(pet.FieldOwnerID), t.C(user.FieldID))
			s.Select(
				sql.As(t.C(user.FieldName), "user_name"),
				sql.As(s.C(pet.FieldName), "pet_name"),
			)
		}).
		ScanX(ctx, &names)
	require.Len(t, names, 1)
	require.Equal(t, "a8m", names[0].User)
	require.Equal(t, "Pedro", names[0].Pet)

	id := client.Group.Query().
		Where(group.HasUsersWith(user.ID(usr.ID))).
		Limit(1).
		QueryUsers().
		QueryPets().
		OnlyIDX(ctx)
	require.Equal(t, pedro.ID, id)

	affected := client.Group.
		Update().
		ClearUsers().
		Where(
			group.And(
				group.Name(groups[0].Name),
				group.HasUsersWith(
					user.HasPetsWith(
						pet.Name(pedro.Name),
					),
				),
			),
		).
		SaveX(ctx)
	require.Equal(t, 1, affected)

	exist := groups[0].QueryUsers().ExistX(ctx)
	require.False(t, exist)
	exist = groups[1].QueryUsers().ExistX(ctx)
	require.True(t, exist)
	exist = pedro.QueryOwner().ExistX(ctx)
	require.True(t, exist)
	pedro = pedro.Update().ClearOwner().SaveX(ctx)
	exist = pedro.QueryOwner().ExistX(ctx)
	require.False(t, exist)

	require.Equal(t, client.User.Query().CountX(ctx), len(client.User.Query().AllX(ctx)))
	require.Equal(t, client.Pet.Query().CountX(ctx), len(client.Pet.Query().AllX(ctx)))
}

func setupSchema(t *testing.T, drv *sql.Driver) {
	client := ent.NewClient(ent.Driver(&rewriter{drv}))
	err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false))
	require.NoError(t, err)
}

type rewriter struct {
	dialect.Driver
}

func (r *rewriter) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := r.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txRewriter{tx}, nil
}

type txRewriter struct {
	dialect.Tx
}

func (r *txRewriter) Exec(ctx context.Context, query string, args, v interface{}) error {
	rp := strings.NewReplacer("`groups`", "`db1`.`groups`", "`pets`", "`db1`.`pets`", "`users`", "`db2`.`users`", "`group_users`", "`db2`.`group_users`")
	query = rp.Replace(query)
	return r.Tx.Exec(ctx, query, args, v)
}
