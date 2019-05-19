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
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersv1alpha2 "github.com/prksu/publr/pkg/api/users/v1alpha2"
	"github.com/prksu/publr/pkg/service/util"
)

func TestServer_ListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := NewMockUserDatastore(ctrl)
	pageToken := util.NewPageToken()

	testuserlist := &usersv1alpha2.UserList{
		Users: []*usersv1alpha2.User{
			{
				Name:     strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
				Email:    "testuser@mysites.site",
				Username: "testuser",
				Fullname: "Test User",
				Role:     0,
			},
			{
				Name:     strings.Join([]string{"sites", "mysites.site", "users", "testauthor"}, "/"),
				Email:    "testauthor@mysites.site",
				Username: "testauthor",
				Fullname: "Test Author",
				Role:     1,
			},
			{
				Name:     strings.Join([]string{"sites", "mysites.site", "users", "testadmin"}, "/"),
				Email:    "testadmin@mysites.site",
				Username: "testadmin",
				Fullname: "Test Admin",
				Role:     2,
			},
			{
				Name:     strings.Join([]string{"sites", "mysites.site", "users", "testowner"}, "/"),
				Email:    "testowner@mysites.site",
				Username: "testowner",
				Fullname: "Test Owner",
				Role:     3,
			},
		},
	}

	type args struct {
		ctx context.Context
		req *usersv1alpha2.ListUserRequest
	}
	tests := []struct {
		name              string
		args              args
		expectedListUsers *gomock.Call
		want              *usersv1alpha2.UserList
		wantErr           bool
	}{
		{
			name: "Test list users",
			args: args{
				context.Background(),
				&usersv1alpha2.ListUserRequest{
					Parent: strings.Join([]string{"sites", "mysites.site"}, "/"),
				},
			},
			expectedListUsers: mockDatastore.EXPECT().List("mysites.site", 0, 10).Return(testuserlist.Users, len(testuserlist.Users), nil),
			want:              testuserlist,
			wantErr:           false,
		},
		{
			name: "Test list users with page_size",
			args: args{
				context.Background(),
				&usersv1alpha2.ListUserRequest{
					Parent:   strings.Join([]string{"sites", "mysites.site"}, "/"),
					PageSize: 2,
				},
			},
			expectedListUsers: mockDatastore.EXPECT().List("mysites.site", 0, 2).Return(testuserlist.Users[0:2], len(testuserlist.Users), nil),
			want: &usersv1alpha2.UserList{
				Users:         testuserlist.Users[0:2],
				NextPageToken: pageToken.Generate(2),
			},
			wantErr: false,
		},
		{
			name: "Test list users with page_size and page_token",
			args: args{
				context.Background(),
				&usersv1alpha2.ListUserRequest{
					Parent:    strings.Join([]string{"sites", "mysites.site"}, "/"),
					PageSize:  2,
					PageToken: pageToken.Generate(2),
				},
			},
			expectedListUsers: mockDatastore.EXPECT().List("mysites.site", 2, 2).Return(testuserlist.Users[2:4], len(testuserlist.Users), nil),
			want: &usersv1alpha2.UserList{
				Users: testuserlist.Users[2:4],
			},
			wantErr: false,
		},
		{
			name: "Test list users with invalid page_token",
			args: args{
				context.Background(),
				&usersv1alpha2.ListUserRequest{
					Parent:    strings.Join([]string{"sites", "mysites.site"}, "/"),
					PageToken: "invalid",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.ListUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := NewMockUserDatastore(ctrl)
	pageToken := util.NewPageToken()

	testuser := &usersv1alpha2.User{
		Name:     strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
		Email:    "testuser@mysites.site",
		Password: "secret",
		Username: "testuser",
		Fullname: "Test User",
	}

	type args struct {
		ctx context.Context
		req *usersv1alpha2.CreateUserRequest
	}
	tests := []struct {
		name               string
		args               args
		expectedCreateUser *gomock.Call
		expectedGetUser    *gomock.Call
		want               *usersv1alpha2.User
		wantErr            bool
	}{
		{
			name: "Test create user",
			args: args{
				context.Background(),
				&usersv1alpha2.CreateUserRequest{
					Parent: strings.Join([]string{"sites", "mysites.site"}, "/"),
					User:   testuser,
				},
			},
			expectedCreateUser: mockDatastore.EXPECT().Create("mysites.site", testuser).Return(nil),
			expectedGetUser:    mockDatastore.EXPECT().Get("mysites.site", testuser.Username).Return(testuser, nil),
			want:               testuser,
		},
		{
			name: "Test create user with nil request",
			args: args{
				context.Background(),
				&usersv1alpha2.CreateUserRequest{},
			},
			wantErr: true,
		},
		{
			name: "Test create user with empty username",
			args: args{
				context.Background(),
				&usersv1alpha2.CreateUserRequest{
					Parent: strings.Join([]string{"sites", "mysites.site"}, "/"),
					User: &usersv1alpha2.User{
						Email:    "testuser@mysites.site",
						Password: "secret",
						Fullname: "Test User",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create user with empty username",
			args: args{
				context.Background(),
				&usersv1alpha2.CreateUserRequest{
					Parent: strings.Join([]string{"sites", "mysites.site"}, "/"),
					User: &usersv1alpha2.User{
						Username: "testuser",
						Password: "secret",
						Fullname: "Test User",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create user with empty password",
			args: args{
				context.Background(),
				&usersv1alpha2.CreateUserRequest{
					Parent: strings.Join([]string{"sites", "mysites.site"}, "/"),
					User: &usersv1alpha2.User{
						Email:    "testuser@mysites.site",
						Username: "testuser",
						Fullname: "Test User",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := NewMockUserDatastore(ctrl)
	pageToken := util.NewPageToken()

	testuser := &usersv1alpha2.User{
		Name:     strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
		Email:    "testuser@mysites.site",
		Password: "secret",
		Username: "testuser",
		Fullname: "Test User",
	}

	type args struct {
		ctx context.Context
		req *usersv1alpha2.GetUserRequest
	}
	tests := []struct {
		name            string
		args            args
		expectedGetUser *gomock.Call
		want            *usersv1alpha2.User
		wantErr         bool
	}{
		{
			name: "Test get user",
			args: args{
				context.Background(),
				&usersv1alpha2.GetUserRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
				},
			},
			expectedGetUser: mockDatastore.EXPECT().Get("mysites.site", testuser.Username).Return(testuser, nil),
			want:            testuser,
		},
		{
			name: "Test get user not found",
			args: args{
				context.Background(),
				&usersv1alpha2.GetUserRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "users", "notfound"}, "/"),
				},
			},
			expectedGetUser: mockDatastore.EXPECT().Get("mysites.site", "notfound").Return(nil, status.Error(codes.NotFound, "user not found")),
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := NewMockUserDatastore(ctrl)
	pageToken := util.NewPageToken()

	testuser := &usersv1alpha2.User{
		Name:     strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
		Email:    "testuser@mysites.site",
		Password: "secret",
		Username: "testuser",
		Fullname: "Test User",
	}

	testupdateuser := &usersv1alpha2.User{
		Name:     strings.Join([]string{"sites", "mysites.site", "users", "testupdateuser"}, "/"),
		Email:    "testupdateuser@mysites.site",
		Password: "secret",
		Username: "testupdateuser",
		Fullname: "Test User",
	}

	type args struct {
		ctx context.Context
		req *usersv1alpha2.UpdateUserRequest
	}
	tests := []struct {
		name               string
		args               args
		expectedGetUser    *gomock.Call
		expectedUpdateUser *gomock.Call
		want               *usersv1alpha2.User
		wantErr            bool
	}{
		{
			name: "Test update user",
			args: args{
				context.Background(),
				&usersv1alpha2.UpdateUserRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
					User: testupdateuser,
				},
			},
			expectedGetUser:    mockDatastore.EXPECT().Get("mysites.site", testuser.Username).Return(testuser, nil),
			expectedUpdateUser: mockDatastore.EXPECT().Update("mysites.site", "testuser", testupdateuser).Return(nil),
			want:               testupdateuser,
		},
		{
			name: "Test update user not found",
			args: args{
				context.Background(),
				&usersv1alpha2.UpdateUserRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "users", "notfound"}, "/"),
					User: testupdateuser,
				},
			},
			expectedGetUser: mockDatastore.EXPECT().Get("mysites.site", "notfound").Return(nil, status.Error(codes.NotFound, "user not found")),
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := NewMockUserDatastore(ctrl)
	pageToken := util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha2.DeleteUserRequest
	}
	tests := []struct {
		name               string
		args               args
		want               *empty.Empty
		expectedDeleteUser *gomock.Call
		wantErr            bool
	}{
		{
			name: "Test delete user",
			args: args{
				context.Background(),
				&usersv1alpha2.DeleteUserRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "users", "testuser"}, "/"),
				},
			},
			expectedDeleteUser: mockDatastore.EXPECT().Delete("mysites.site", "testuser").Return(nil),
			want:               &empty.Empty{},
		},
		{
			name: "Test delete user not found",
			args: args{
				context.Background(),
				&usersv1alpha2.DeleteUserRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "users", "notfound"}, "/"),
				},
			},
			expectedDeleteUser: mockDatastore.EXPECT().Delete("mysites.site", "notfound").Return(status.Error(codes.NotFound, "user not found")),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
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

func TestServer_SearchUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := NewMockUserDatastore(ctrl)
	pageToken := util.NewPageToken()

	type args struct {
		ctx context.Context
		req *usersv1alpha2.SearchUserRequest
	}
	tests := []struct {
		name               string
		args               args
		expectedSearchUser *gomock.Call
		want               *usersv1alpha2.UserList
		wantErr            bool
	}{
		{
			name: "Test list users",
			args: args{
				context.Background(),
				&usersv1alpha2.SearchUserRequest{
					Parent: strings.Join([]string{"sites", "mysites.site"}, "/"),
					Query:  "whocares",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.SearchUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.SearchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
