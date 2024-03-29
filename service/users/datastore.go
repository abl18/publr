// Copyright 2019 Publr Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersv1alpha3 "github.com/prksu/publr/pkg/api/users/v1alpha3"
	"github.com/prksu/publr/pkg/storage/database"
)


// UserDatastore interface
type UserDatastore interface {
	List(ctx context.Context, sitedomain string, start, offset int) ([]*usersv1alpha3.User, int, error)
	Create(ctx context.Context, sitedomain string, user *usersv1alpha3.User) error
	Get(ctx context.Context, sitedomain, username string) (*usersv1alpha3.User, error)
	Update(ctx context.Context, sitedomain, username string, user *usersv1alpha3.User) error
	Delete(ctx context.Context, sitedomain, username string) error
	Search(ctx context.Context, query string) ([]*usersv1alpha3.User, error)
}

// datastore implement users service datastore
type datastore struct {
	DB *sql.DB
}

// NewUserDatastore create new users service datastore instance with configured database connection
func NewUserDatastore() UserDatastore {
	return NewUserDatastoreWithDB(database.NewDatabase().Connect())
}

// NewUserDatastoreWithDB create new users service datastore instance with sql.DB params.
func NewUserDatastoreWithDB(database *sql.DB) UserDatastore {
	ds := new(datastore)
	ds.DB = database
	return ds
}

func (ds *datastore) List(ctx context.Context, sitedomain string, start, limit int) ([]*usersv1alpha3.User, int, error) {
	var users []*usersv1alpha3.User
	var foundRows int

	sqlrows := `
		SELECT u.email, u.username, u.fullname, su.role, u.createtime, u.updatetime
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
		WHERE su.site_domain=?
		LIMIT ?, ?
	`

	sqlcount := `
		SELECT COUNT(*)
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
		WHERE su.site_domain=?
	`

	rows, err := ds.DB.QueryContext(ctx, sqlrows, sitedomain, start, limit)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	for rows.Next() {
		var user usersv1alpha3.User
		var createTime mysql.NullTime
		var updateTime mysql.NullTime
		if err := rows.Scan(&user.Email, &user.Username, &user.Fullname, &user.Role, &createTime, &updateTime); err != nil {
			return nil, 0, err
		}
		user.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
		user.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)
		users = append(users, &user)
	}

	if err := ds.DB.QueryRowContext(ctx, sqlcount, sitedomain).Scan(&foundRows); err != nil {
		return nil, 0, err
	}

	return users, foundRows, nil
}

func (ds *datastore) Create(ctx context.Context, sitedomain string, user *usersv1alpha3.User) error {
	tx, err := ds.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	{
		stmt, err := tx.PrepareContext(ctx, "INSERT INTO users (email, username, fullname) VALUES (?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.ExecContext(ctx, user.Email, user.Username, user.Fullname); err != nil {
			tx.Rollback()
			return err
		}
	}
	{
		stmt, err := tx.PrepareContext(ctx, "INSERT INTO site_users (user_username, site_domain, role) VALUES (?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.ExecContext(ctx, user.Username, sitedomain, user.Role); err != nil {
			tx.Rollback()
			if err.(*mysql.MySQLError).Number == 1062 {
				return status.Error(codes.AlreadyExists, "user already exists")
			}
			return err
		}
	}

	return tx.Commit()
}

func (ds *datastore) Get(ctx context.Context, sitedomain, username string) (*usersv1alpha3.User, error) {
	user := new(usersv1alpha3.User)
	var createTime mysql.NullTime
	var updateTime mysql.NullTime
	var sqlquery = `
		SELECT u.email, u.username, u.fullname, su.role, u.createtime, u.updatetime
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
		WHERE username=? AND su.site_domain=?
	`
	if err := ds.DB.QueryRowContext(ctx, sqlquery, username, sitedomain).
		Scan(&user.Email, &user.Username, &user.Fullname, &user.Role, &createTime, &updateTime); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, err
	}

	user.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
	user.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)
	return user, nil
}

func (ds *datastore) Update(ctx context.Context, sitedomain, username string, user *usersv1alpha3.User) error {
	var sqlquery = `
		UPDATE users AS u LEFT JOIN site_users AS su on u.username=su.user_username
		SET u.username=?, u.fullname=?, su.role=?
		WHERE u.username=? AND su.site_domain=?
	`
	stmt, err := ds.DB.PrepareContext(ctx, sqlquery)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, user.Username, user.Fullname, user.Role, username, sitedomain); err != nil {
		return err
	}

	return nil
}

func (ds *datastore) Delete(ctx context.Context, sitedomain, username string) error {
	var sqlquery = `
		DELETE users FROM users
		LEFT JOIN site_users on users.username=site_users.user_username
		WHERE site_users.site_domain=? AND users.username=?
	`

	result, err := ds.DB.ExecContext(ctx, sqlquery, sitedomain, username)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return status.Error(codes.NotFound, "user not found")
	}

	return nil
}

func (ds *datastore) Search(ctx context.Context, query string) ([]*usersv1alpha3.User, error) {
	return nil, nil
}
