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
	"encoding/json"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersv1alpha1 "github.com/prksu/publr/pkg/api/users/v1alpha1"
	"github.com/prksu/publr/pkg/util"
)

// Users service
var (
	ServiceName    = "users"
	ServiceAddress = "0.0.0.0:9000"
	ServiceVersion = "v1alpha1"
)

// Server implement usersv1alpha1.UserServiceServer.
type Server struct {
	User      UserDatastore
	PageToken util.PageToken
}

// NewServer create new users service server.
// returns usersv1alpha1.SiteServiceServer
func NewServer() usersv1alpha1.UserServiceServer {
	server := new(Server)
	server.User = NewUserDatastore()
	server.PageToken = util.NewPageToken()
	return server
}

// ListUser handler method
func (s *Server) ListUser(ctx context.Context, req *usersv1alpha1.ListUserRequest) (*usersv1alpha1.UserList, error) {
	parent := req.Parent
	start, err := s.PageToken.Parse(req.PageToken)
	if err != nil {
		return nil, err
	}
	offset := int(req.PageSize)
	if offset == 0 {
		offset = 10
	}

	sitedomain := strings.Split(parent, "/")[1]
	users, err := s.User.List(sitedomain, start, offset)
	if err != nil {
		return nil, err
	}

	for _, i := range users {
		i.Name = strings.Join([]string{parent, "users", i.Username}, "/")
	}

	var nextPageToken string
	if len(users) == offset {
		nextPageToken = s.PageToken.Generate(start + offset)
	}

	res := new(usersv1alpha1.UserList)
	res.Users = users
	res.NextPageToken = nextPageToken
	return res, nil
}

// CreateUser handler method
func (s *Server) CreateUser(ctx context.Context, req *usersv1alpha1.CreateUserRequest) (*usersv1alpha1.User, error) {
	parent := req.Parent
	user := req.User

	if user.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if user.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if user.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	sitedomain := strings.Split(parent, "/")[1]
	username := user.Username

	// TODO: Check if the site is really exists.

	if err := s.User.Create(sitedomain, user); err != nil {
		return nil, err
	}

	res, err := s.User.Get(sitedomain, username)
	if err != nil {
		return nil, err
	}

	res.Name = strings.Join([]string{parent, "users", username}, "/")
	return res, nil
}

// GetUser handler method
func (s *Server) GetUser(ctx context.Context, req *usersv1alpha1.GetUserRequest) (*usersv1alpha1.User, error) {
	name := req.Name
	sitedomain := strings.Split(name, "/")[1]
	username := strings.Split(name, "/")[3]
	res, err := s.User.Get(sitedomain, username)
	if err != nil {
		return nil, err
	}

	res.Name = name
	return res, nil
}

// UpdateUser handler method
func (s *Server) UpdateUser(ctx context.Context, req *usersv1alpha1.UpdateUserRequest) (*usersv1alpha1.User, error) {
	name := req.Name
	sitedomain := strings.Split(name, "/")[1]
	username := strings.Split(name, "/")[3]

	res, err := s.User.Get(sitedomain, username)
	if err != nil {
		return nil, err
	}

	if req.User.Role == 0 && res.Role != 3 {
		res.Role = 0
	} else if res.Role == 3 && req.User.Role != 3 {
		return nil, status.Error(codes.InvalidArgument, "owner role cannot be update")
	}

	b, err := json.Marshal(req.User)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, res); err != nil {
		return nil, err
	}

	if err := s.User.Update(sitedomain, username, res); err != nil {
		return nil, err
	}

	res.Name = strings.Join([]string{"sites", sitedomain, "users", res.Username}, "/")
	return res, nil
}

// DeleteUser handler method
func (s *Server) DeleteUser(ctx context.Context, req *usersv1alpha1.DeleteUserRequest) (*empty.Empty, error) {
	name := req.Name
	sitedomain := strings.Split(name, "/")[1]
	username := strings.Split(name, "/")[3]

	if err := s.User.Delete(sitedomain, username); err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

// SearchUser handler method
func (s *Server) SearchUser(ctx context.Context, req *usersv1alpha1.SearchUserRequest) (*usersv1alpha1.UserList, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}
