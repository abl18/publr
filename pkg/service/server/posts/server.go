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
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	postsv1alpha1 "github.com/prksu/publr/pkg/api/posts/v1alpha1"
	"github.com/prksu/publr/pkg/service/util"
)

// Posts service
var (
	ServiceName    = "posts"
	ServiceAddress = "0.0.0.0:9000"
	ServiceVersion = "v1alpha1"
)

// Server implement postsv1alpha1.PostServiceServer.
type Server struct {
	Post      PostDatastore
	PageToken util.PageToken
}

// NewServiceServer create new users service server.
// returns postsv1alpha1.PostServiceServer
func NewServiceServer() postsv1alpha1.PostServiceServer {
	server := new(Server)
	server.Post = NewPostDatastore()
	server.PageToken = util.NewPageToken()
	return server
}

// ListPost handler method
func (s *Server) ListPost(ctx context.Context, req *postsv1alpha1.ListPostRequest) (*postsv1alpha1.PostList, error) {
	parent := req.Parent
	sparent := strings.Split(parent, "/")

	start, err := s.PageToken.Parse(req.PageToken)
	if err != nil {
		return nil, err
	}

	limit := int(req.PageSize)
	if limit == 0 {
		limit = 10
	}

	var sitedomain string
	var author string

	sitedomain = sparent[1]
	if len(sparent) > 3 {
		author = sparent[3]
	}

	posts, totalSize, err := s.Post.List(sitedomain, author, start, limit)
	if err != nil {
		return nil, err
	}

	for _, i := range posts {
		i.Name = strings.Join([]string{parent, "posts", i.Slug}, "/")
	}

	var nextPageToken string
	if (start + limit) < totalSize {
		nextPageToken = s.PageToken.Generate(start + limit)
	}

	res := new(postsv1alpha1.PostList)
	res.Posts = posts
	res.NextPageToken = nextPageToken
	return res, nil
}

// CreatePost handler method
func (s *Server) CreatePost(ctx context.Context, req *postsv1alpha1.CreatePostRequest) (*postsv1alpha1.Post, error) {
	parent := req.Parent
	post := req.Post

	if post == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if post.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}

	if post.Slug == "" {
		return nil, status.Error(codes.InvalidArgument, "slug is required")
	}

	if post.Html == "" {
		return nil, status.Error(codes.InvalidArgument, "html is required")
	}

	sitedomain := strings.Split(parent, "/")[1]
	author := strings.Split(parent, "/")[3]

	slug := strings.ToLower(strings.Replace(post.Slug, " ", "-", -1))
	if err := s.Post.Create(sitedomain, author, post); err != nil {
		return nil, err
	}

	res, err := s.Post.Get(sitedomain, author, slug)
	if err != nil {
		return nil, err
	}

	res.Name = strings.Join([]string{"sites", sitedomain, "authors", author, "posts", slug}, "/")
	return res, nil
}

// GetPost handler method
func (s *Server) GetPost(ctx context.Context, req *postsv1alpha1.GetPostRequest) (*postsv1alpha1.Post, error) {
	name := req.Name
	sname := strings.Split(name, "/")

	var sitedomain string
	var author string
	var slug string

	sitedomain = sname[1]
	switch len(sname) {
	case 4:
		slug = sname[3]
	case 6:
		author = sname[3]
		slug = sname[5]
	}

	res, err := s.Post.Get(sitedomain, author, slug)
	if err != nil {
		return nil, err
	}

	res.Name = name
	return res, nil
}

// UpdatePost handler method
func (s *Server) UpdatePost(ctx context.Context, req *postsv1alpha1.UpdatePostRequest) (*postsv1alpha1.Post, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// DeletePost handler method
func (s *Server) DeletePost(ctx context.Context, req *postsv1alpha1.DeletePostRequest) (*empty.Empty, error) {
	name := req.Name
	sname := strings.Split(name, "/")

	sitedomain := sname[1]
	author := sname[3]
	slug := sname[5]

	// just check if the posts is exist
	if _, err := s.Post.Get(sitedomain, author, slug); err != nil {
		return nil, err
	}

	if err := s.Post.Delete(sitedomain, author, slug); err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

// SearchPost handler method
func (s *Server) SearchPost(ctx context.Context, req *postsv1alpha1.SearchPostRequest) (*postsv1alpha1.PostList, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}
