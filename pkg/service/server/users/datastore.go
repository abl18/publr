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
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersv1alpha1 "github.com/prksu/publr/pkg/api/users/v1alpha1"
	"github.com/prksu/publr/pkg/storage/database"
)

// UserDatastore interface
type UserDatastore interface {
	List(sitedomain string, start, offset int) ([]*usersv1alpha1.User, error)
	Create(sitedomain string, user *usersv1alpha1.User) error
	Get(sitedomain, username string) (*usersv1alpha1.User, error)
	Update(sitedomain, username string, user *usersv1alpha1.User) error
	Delete(sitedomain, username string) error
	Search(query string) ([]*usersv1alpha1.User, error)
}

// datastore implement users service datastore
type datastore struct {
	DB *sql.DB
}

// NewUserDatastore create new users service datastore instance
func NewUserDatastore() UserDatastore {
	ds := new(datastore)
	ds.DB = database.NewDatabase().Connect()
	return ds
}

func (ds *datastore) List(sitedomain string, start, offset int) ([]*usersv1alpha1.User, error) {
	var users []*usersv1alpha1.User

	sqlquery := `
		SELECT u.email, u.username, u.fullname, su.role, u.createtime, u.updatetime
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
		WHERE su.site_domain=?
		LIMIT ?, ?
	`

	rows, err := ds.DB.Query(sqlquery, sitedomain, start, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user usersv1alpha1.User
		var createTime mysql.NullTime
		var updateTime mysql.NullTime
		if err := rows.Scan(&user.Email, &user.Username, &user.Fullname, &user.Role, &createTime, &updateTime); err != nil {
			return nil, err
		}
		user.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
		user.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)
		users = append(users, &user)
	}

	return users, nil
}

func (ds *datastore) Create(sitedomain string, user *usersv1alpha1.User) error {
	tx, err := ds.DB.Begin()
	if err != nil {
		return err
	}
	{
		stmt, err := tx.Prepare("INSERT INTO users (email, username, fullname) VALUES (?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(user.Email, user.Username, user.Fullname); err != nil {
			tx.Rollback()
			return err
		}
	}
	{
		stmt, err := tx.Prepare("INSERT INTO site_users (user_username, site_domain, role) VALUES (?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(user.Username, sitedomain, user.Role); err != nil {
			tx.Rollback()
			if err.(*mysql.MySQLError).Number == 1062 {
				return status.Error(codes.AlreadyExists, "user already exists")
			}
			return err
		}
	}

	return tx.Commit()
}

func (ds *datastore) Get(sitedomain, username string) (*usersv1alpha1.User, error) {
	user := new(usersv1alpha1.User)
	var createTime mysql.NullTime
	var updateTime mysql.NullTime
	var sqlquery = `
		SELECT u.email, u.username, u.fullname, su.role, u.createtime, u.updatetime
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
		WHERE username=? AND su.site_domain=?
	`
	if err := ds.DB.QueryRow(sqlquery, username, sitedomain).
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

func (ds *datastore) Update(sitedomain, username string, user *usersv1alpha1.User) error {
	var sqlquery = `
		UPDATE users AS u LEFT JOIN site_users AS su on u.username=su.user_username
		SET u.username=?, u.fullname=?, su.role=?
		WHERE u.username=? AND su.site_domain=?
	`
	stmt, err := ds.DB.Prepare(sqlquery)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(user.Username, user.Fullname, user.Role, username, sitedomain); err != nil {
		return err
	}

	return nil
}

func (ds *datastore) Delete(sitedomain, username string) error {
	var sqlquery = `
		DELETE users FROM users
		LEFT JOIN site_users on users.username=site_users.user_username
		WHERE site_users.site_domain=? AND users.username=?
	`

	if _, err := ds.DB.Exec(sqlquery, sitedomain, username); err != nil {
		return err
	}

	return nil
}

func (ds *datastore) Search(query string) ([]*usersv1alpha1.User, error) { return nil, nil }
