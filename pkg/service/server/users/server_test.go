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
	"log"
	"reflect"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	usersv1alpha1 "github.com/prksu/publr/pkg/api/users/v1alpha1"
	"github.com/prksu/publr/pkg/bindata/schema"
	"github.com/prksu/publr/pkg/bindata/testdata"
	"github.com/prksu/publr/pkg/service/util"
	"github.com/prksu/publr/pkg/storage/database"
)

var (
	DSN = "root:@/publr_test?autocommit=true&parseTime=true&multiStatements=true"
)

func init() {
	database := database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect()
	defer database.Close()

	database.Exec("DROP TABLE IF EXISTS site_users")
	database.Exec("DROP TABLE IF EXISTS users")

	schema, err := schema.Asset("data/schema/users.sql")
	if err != nil {
		log.Fatal(err)
	}

	testdata, err := testdata.Asset("data/testdata/users.sql")
	if err != nil {
		log.Fatal(err)
	}

	database.Exec(string(schema))
	database.Exec(string(testdata))

}

func TestServer_ListUser(t *testing.T) {
	server := new(Server)
	server.User = NewUserDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha1.ListUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *usersv1alpha1.UserList
		wantErr bool
	}{
		{
			name: "Test list user",
			args: args{
				context.Background(),
				&usersv1alpha1.ListUserRequest{
					Parent: "sites/mysites.site",
				},
			},
			want: &usersv1alpha1.UserList{
				Users: []*usersv1alpha1.User{
					{
						Name:     "sites/mysites.site/users/userdemo",
						Email:    "userdemo@mysites.site",
						Username: "userdemo",
						Fullname: "User Demo",
						Role:     0,
					},
					{
						Name:     "sites/mysites.site/users/authordemo",
						Email:    "authordemo@mysites.site",
						Username: "authordemo",
						Fullname: "Author Demo",
						Role:     1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := server.ListUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				for _, user := range tt.want.Users {
					for _, g := range got.Users {
						user.CreateTime = g.CreateTime
						user.UpdateTime = g.UpdateTime
					}
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CreateUser(t *testing.T) {
	server := new(Server)
	server.User = NewUserDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha1.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *usersv1alpha1.User
		wantErr bool
	}{
		{
			name: "Test create user",
			args: args{
				context.Background(),
				&usersv1alpha1.CreateUserRequest{
					Parent: "sites/mysites.site",
					User: &usersv1alpha1.User{
						Email:    "testuser@mysites.site",
						Password: "secret",
						Username: "testuser",
						Fullname: "Test User",
					},
				},
			},
			want: &usersv1alpha1.User{
				Name:     "sites/mysites.site/users/testuser",
				Email:    "testuser@mysites.site",
				Username: "testuser",
				Fullname: "Test User",
				Role:     0,
			},
			wantErr: false,
		},
		{
			name: "Test create existing user",
			args: args{
				context.Background(),
				&usersv1alpha1.CreateUserRequest{
					Parent: "sites/mysites.site",
					User: &usersv1alpha1.User{
						Email:    "testuser@mysites.site",
						Password: "secret",
						Username: "testuser",
						Fullname: "Test User",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create user with null request",
			args: args{
				context.Background(),
				&usersv1alpha1.CreateUserRequest{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				tt.want.CreateTime = got.CreateTime
				tt.want.UpdateTime = got.UpdateTime
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUser(t *testing.T) {
	server := new(Server)
	server.User = NewUserDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha1.GetUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *usersv1alpha1.User
		wantErr bool
	}{
		{
			name: "Test get user",
			args: args{
				context.Background(),
				&usersv1alpha1.GetUserRequest{
					Name: "sites/mysites.site/users/userdemo",
				},
			},
			want: &usersv1alpha1.User{
				Name:     "sites/mysites.site/users/userdemo",
				Email:    "userdemo@mysites.site",
				Username: "userdemo",
				Fullname: "User Demo",
				Role:     0,
			},
			wantErr: false,
		},
		{
			name: "Test get not existing user",
			args: args{
				context.Background(),
				&usersv1alpha1.GetUserRequest{
					Name: "sites/mysites.site/users/notexists",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				tt.want.CreateTime = got.CreateTime
				tt.want.UpdateTime = got.UpdateTime
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateUser(t *testing.T) {
	server := new(Server)
	server.User = NewUserDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha1.UpdateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *usersv1alpha1.User
		wantErr bool
	}{
		{
			name: "Test update user",
			args: args{
				context.Background(),
				&usersv1alpha1.UpdateUserRequest{
					Name: "sites/mysites.site/users/testuser",
					User: &usersv1alpha1.User{
						Email:    "updatetestuser@mysites.site",
						Username: "updatetestuser",
						Fullname: "Update Test User",
					},
				},
			},
			want: &usersv1alpha1.User{
				Name:     "sites/mysites.site/users/updatetestuser",
				Email:    "updatetestuser@mysites.site",
				Username: "updatetestuser",
				Fullname: "Update Test User",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				tt.want.CreateTime = got.CreateTime
				tt.want.UpdateTime = got.UpdateTime
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteUser(t *testing.T) {
	server := new(Server)
	server.User = NewUserDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha1.DeleteUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			name: "Test delete user",
			args: args{
				context.Background(),
				&usersv1alpha1.DeleteUserRequest{
					Name: "sites/mysites.site/users/userdemo",
				},
			},
			want:    &empty.Empty{},
			wantErr: false,
		},
		{
			name: "Test delete not existing user",
			args: args{
				context.Background(),
				&usersv1alpha1.DeleteUserRequest{
					Name: "sites/mysites.site/users/userdemo",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.DeleteUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
