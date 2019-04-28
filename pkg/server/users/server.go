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

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersv1alpha1 "github.com/prksu/publr/pkg/api/users/v1alpha1"
)

// Users service
var (
	ServiceName    = "users"
	ServiceAddress = "0.0.0.0:9000"
	ServiceVersion = "v1alpha1"
)

// Server implement usersv1alpha1.UserServiceServer.
type Server struct {
	DS UserDatastore
}

// NewServer create new users service server.
// returns usersv1alpha1.SiteServiceServer
func NewServer() usersv1alpha1.UserServiceServer {
	server := new(Server)
	server.DS = NewDatastore()
	return server
}

// ListUser handler method
func (s *Server) ListUser(ctx context.Context, req *usersv1alpha1.ListUserRequest) (*usersv1alpha1.UserList, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// CreateUser handler method
func (s *Server) CreateUser(ctx context.Context, req *usersv1alpha1.CreateUserRequest) (*usersv1alpha1.User, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// GetUser handler method
func (s *Server) GetUser(ctx context.Context, req *usersv1alpha1.GetUserRequest) (*usersv1alpha1.User, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// UpdateUser handler method
func (s *Server) UpdateUser(ctx context.Context, req *usersv1alpha1.UpdateUserRequest) (*usersv1alpha1.User, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// DeleteUser handler method
func (s *Server) DeleteUser(ctx context.Context, req *usersv1alpha1.DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// SearchUser handler method
func (s *Server) SearchUser(ctx context.Context, req *usersv1alpha1.SearchUserRequest) (*usersv1alpha1.UserList, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}
