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

package posts

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	postsv1alpha1 "github.com/prksu/publr/pkg/api/posts/v1alpha1"
)

// Posts service
var (
	ServiceName    = "posts"
	ServiceAddress = "0.0.0.0:9000"
	ServiceVersion = "v1alpha1"
)

// Server implement postsv1alpha1.PostServiceServer.
type Server struct {
	DS PostDatastore
}

// NewServer create new users service server.
// returns postsv1alpha1.PostServiceServer
func NewServer() postsv1alpha1.PostServiceServer {
	server := new(Server)
	server.DS = NewDatastore()
	return server
}

// ListPost handler method
func (s *Server) ListPost(ctx context.Context, req *postsv1alpha1.ListPostRequest) (*postsv1alpha1.PostList, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// CreatePost handler method
func (s *Server) CreatePost(ctx context.Context, req *postsv1alpha1.CreatePostRequest) (*postsv1alpha1.Post, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// GetPost handler method
func (s *Server) GetPost(ctx context.Context, req *postsv1alpha1.GetPostRequest) (*postsv1alpha1.Post, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// UpdatePost handler method
func (s *Server) UpdatePost(ctx context.Context, req *postsv1alpha1.UpdatePostRequest) (*postsv1alpha1.Post, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// DeletePost handler method
func (s *Server) DeletePost(ctx context.Context, req *postsv1alpha1.DeletePostRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// SearchPost handler method
func (s *Server) SearchPost(ctx context.Context, req *postsv1alpha1.SearchPostRequest) (*postsv1alpha1.PostList, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}
