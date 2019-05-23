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

package datastore

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/protobuf/ptypes"

	usersv1alpha2 "github.com/prksu/publr/pkg/api/users/v1alpha2"
)

func Test_datastore_List(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	timestamp := time.Now()
	protoTimestamp, err := ptypes.TimestampProto(timestamp)
	if err != nil {
		t.Fatal(err)
	}

	sitedomain := "mysites.site"
	users := []*usersv1alpha2.User{
		{
			Email:      "testuser@mysites.site",
			Username:   "testuser",
			Fullname:   "Test User",
			Role:       0,
			CreateTime: protoTimestamp,
			UpdateTime: protoTimestamp,
		},
		{
			Email:      "testauthor@mysites.site",
			Username:   "testauthor",
			Fullname:   "Test Author",
			Role:       1,
			CreateTime: protoTimestamp,
			UpdateTime: protoTimestamp,
		},
		{
			Email:      "testadmin@mysites.site",
			Username:   "testadmin",
			Fullname:   "Test Admin",
			Role:       2,
			CreateTime: protoTimestamp,
			UpdateTime: protoTimestamp,
		},
		{
			Email:      "testowner@mysites.site",
			Username:   "testowner",
			Fullname:   "Test Owner",
			Role:       3,
			CreateTime: protoTimestamp,
			UpdateTime: protoTimestamp,
		},
	}

	rows := new(sqlmock.Rows)
	columns := sqlmock.NewRows([]string{"email", "username", "fullname", "role", "createtime", "updatetime"})
	for _, u := range users {
		rows = columns.AddRow(u.Email, u.Username, u.Fullname, u.Role, timestamp, timestamp)
	}

	sqlrows := `
		SELECT u.email, u.username, u.fullname, su.role, u.createtime, u.updatetime
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
	`

	sqlcount := `
		SELECT COUNT(.+)
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
	`

	type args struct {
		sitedomain string
		start      int
		limit      int
	}
	tests := []struct {
		name               string
		args               args
		expectedQuery      *sqlmock.ExpectedQuery
		expectedCountQuery *sqlmock.ExpectedQuery
		want               []*usersv1alpha2.User
		want1              int
		wantErr            bool
	}{
		{
			name: "Test list users",
			args: args{
				sitedomain: sitedomain,
				start:      0,
				limit:      10,
			},
			expectedQuery:      mock.ExpectQuery(sqlrows).WithArgs(sitedomain, 0, 10).WillReturnRows(rows),
			expectedCountQuery: mock.ExpectQuery(sqlcount).WithArgs(sitedomain).WillReturnRows(sqlmock.NewRows([]string{"found_rows"}).AddRow(len(users))),
			want:               users,
			want1:              len(users),
		},
		// TODO: Add more test cases. eg: with start and limit condition.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewUserDatastoreWithDB(database)
			got, got1, err := datastore.List(tt.args.sitedomain, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("datastore.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("datastore.List() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("datastore.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_datastore_Create(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	sitedomain := "mysites.site"
	user := &usersv1alpha2.User{
		Email:    "testuser@mysites.site",
		Username: "testuser",
		Fullname: "Test User",
		Role:     0,
	}

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		sitedomain string
		user       *usersv1alpha2.User
	}
	tests := []struct {
		name             string
		args             args
		expectedBegin    *sqlmock.ExpectedBegin
		expectedRollback *sqlmock.ExpectedRollback
		expectedExec1    *sqlmock.ExpectedExec
		expectedExec2    *sqlmock.ExpectedExec
		expectedCommit   *sqlmock.ExpectedCommit
		wantErr          bool
	}{
		{
			name: "Test create user",
			args: args{
				sitedomain,
				user,
			},
			expectedBegin:  mock.ExpectBegin(),
			expectedExec1:  mock.ExpectPrepare("INSERT INTO users").ExpectExec().WithArgs(user.Email, user.Username, user.Fullname).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedExec2:  mock.ExpectPrepare("INSERT INTO site_users").ExpectExec().WithArgs(user.Username, sitedomain, user.Role).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedCommit: mock.ExpectCommit(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewUserDatastoreWithDB(database)
			if err := datastore.Create(tt.args.sitedomain, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("datastore.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_datastore_Get(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	timestamp := time.Now()
	protoTimestamp, err := ptypes.TimestampProto(timestamp)
	if err != nil {
		t.Fatal(err)
	}

	var sqlquery = `
		SELECT u.email, u.username, u.fullname, su.role, u.createtime, u.updatetime
		FROM users AS u 
		LEFT JOIN site_users AS su on u.username=su.user_username
	`

	sitedomain := "mysites.site"
	user := &usersv1alpha2.User{
		Email:      "testuser@mysites.site",
		Username:   "testuser",
		Fullname:   "Test User",
		Role:       0,
		CreateTime: protoTimestamp,
		UpdateTime: protoTimestamp,
	}

	rows := sqlmock.NewRows([]string{"email", "username", "fullname", "role", "createtime", "updatetime"}).
		AddRow(user.Email, user.Username, user.Fullname, user.Role, timestamp, timestamp)

	type args struct {
		sitedomain string
		username   string
	}
	tests := []struct {
		name          string
		args          args
		expectedQuery *sqlmock.ExpectedQuery
		want          *usersv1alpha2.User
		wantErr       bool
	}{
		{
			name: "Test get user",
			args: args{
				sitedomain: sitedomain,
				username:   user.Username,
			},
			expectedQuery: mock.ExpectQuery(sqlquery).WithArgs(user.Username, sitedomain).WillReturnRows(rows),
			want:          user,
		},
		{
			name: "Test get user not found",
			args: args{
				sitedomain: sitedomain,
				username:   "notfound",
			},
			expectedQuery: mock.ExpectQuery(sqlquery).WithArgs("notfound", sitedomain).WillReturnError(sql.ErrNoRows),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewUserDatastoreWithDB(database)
			got, err := datastore.Get(tt.args.sitedomain, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("datastore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("datastore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_datastore_Update(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	timestamp := time.Now()
	protoTimestamp, err := ptypes.TimestampProto(timestamp)
	if err != nil {
		t.Fatal(err)
	}

	sitedomain := "mysites.site"
	user := &usersv1alpha2.User{
		Email:      "testuser@mysites.site",
		Username:   "testuser",
		Fullname:   "Test User",
		Role:       0,
		CreateTime: protoTimestamp,
		UpdateTime: protoTimestamp,
	}
	updateuser := &usersv1alpha2.User{
		Email:      "testupdateuser@mysites.site",
		Username:   "testupdateuser",
		Fullname:   "Test Update User",
		Role:       0,
		CreateTime: user.CreateTime,
		UpdateTime: protoTimestamp,
	}

	var sqlquery = `
		UPDATE users AS u LEFT JOIN site_users AS su on u.username=su.user_username
	`

	type args struct {
		sitedomain string
		username   string
		user       *usersv1alpha2.User
	}
	tests := []struct {
		name         string
		args         args
		expectedExec *sqlmock.ExpectedExec
		wantErr      bool
	}{
		{
			name: "Test update users",
			args: args{
				sitedomain: sitedomain,
				username:   user.Username,
				user:       updateuser,
			},
			expectedExec: mock.ExpectPrepare(sqlquery).ExpectExec().WithArgs(updateuser.Username, updateuser.Fullname, updateuser.Role, user.Username, sitedomain).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewUserDatastoreWithDB(database)
			if err := datastore.Update(tt.args.sitedomain, tt.args.username, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("datastore.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_datastore_Delete(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	sitedomain := "mysites.site"
	username := "testuser"

	var sqlquery = `
		DELETE users FROM users
		LEFT JOIN site_users on users.username=site_users.user_username
	`

	type args struct {
		sitedomain string
		username   string
	}
	tests := []struct {
		name         string
		args         args
		expectedExec *sqlmock.ExpectedExec
		wantErr      bool
	}{
		{
			name: "Test delete user",
			args: args{
				sitedomain: sitedomain,
				username:   username,
			},
			expectedExec: mock.ExpectExec(sqlquery).WithArgs(sitedomain, username).WillReturnResult(sqlmock.NewResult(0, 1)),
		},
		{
			name: "Test delete user not found",
			args: args{
				sitedomain: sitedomain,
				username:   "notfound",
			},
			expectedExec: mock.ExpectExec(sqlquery).WithArgs(sitedomain, "notfound").WillReturnResult(sqlmock.NewResult(0, 0)),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewUserDatastoreWithDB(database)
			if err := datastore.Delete(tt.args.sitedomain, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("datastore.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
