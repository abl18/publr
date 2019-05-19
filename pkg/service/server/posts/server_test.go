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
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	postsv1alpha2 "github.com/prksu/publr/pkg/api/posts/v1alpha2"
	"github.com/prksu/publr/pkg/service/util"
)

func TestServer_ListPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatastore := NewMockPostDatastore(ctrl)
	pageToken := util.NewPageToken()

	testpostlist := &postsv1alpha2.PostList{
		Posts: []*postsv1alpha2.Post{
			{
				Name:      strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "first-post"}, "/"),
				Title:     "First Post",
				Slug:      "first-post",
				Html:      "<p>First Post</p>",
				Image:     "image.png",
				Published: true,
			},
			{
				Name:      strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "second-post"}, "/"),
				Title:     "Second Post",
				Slug:      "second-post",
				Html:      "<p>Second Post</p>",
				Published: true,
			},
			{
				Name:      strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "third-post"}, "/"),
				Title:     "Third Post",
				Slug:      "third-post",
				Html:      "<p>Third Post</p>",
				Published: true,
			},
			{
				Name:      strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "fourth-post"}, "/"),
				Title:     "Fourth Post",
				Slug:      "fourth-post",
				Html:      "<p>Fourth Post</p>",
				Published: false,
			},
		},
	}

	type args struct {
		ctx context.Context
		req *postsv1alpha2.ListPostRequest
	}
	tests := []struct {
		name              string
		args              args
		expectedListPosts *gomock.Call
		want              *postsv1alpha2.PostList
		wantErr           bool
	}{
		{
			name: "Test list posts",
			args: args{
				context.Background(),
				&postsv1alpha2.ListPostRequest{
					Parent: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
				},
			},
			expectedListPosts: mockDatastore.EXPECT().List("mysites.site", "testauthor", 0, 10).Return(testpostlist.Posts, len(testpostlist.Posts), nil),
			want:              testpostlist,
		},
		{
			name: "Test list posts with page_size",
			args: args{
				context.Background(),
				&postsv1alpha2.ListPostRequest{
					Parent:   strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
					PageSize: 2,
				},
			},
			expectedListPosts: mockDatastore.EXPECT().List("mysites.site", "testauthor", 0, 2).Return(testpostlist.Posts[0:2], len(testpostlist.Posts), nil),
			want: &postsv1alpha2.PostList{
				Posts:         testpostlist.Posts[0:2],
				NextPageToken: pageToken.Generate(2),
			},
		},
		{
			name: "Test list posts with page_size and page_token",
			args: args{
				context.Background(),
				&postsv1alpha2.ListPostRequest{
					Parent:    strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
					PageSize:  2,
					PageToken: pageToken.Generate(2),
				},
			},
			expectedListPosts: mockDatastore.EXPECT().List("mysites.site", "testauthor", 2, 2).Return(testpostlist.Posts[2:4], len(testpostlist.Posts), nil),
			want: &postsv1alpha2.PostList{
				Posts: testpostlist.Posts[2:4],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.ListPost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ListPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ListPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatastore := NewMockPostDatastore(ctrl)
	pageToken := util.NewPageToken()

	testpost := &postsv1alpha2.Post{
		Name:      strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "first-post"}, "/"),
		Title:     "First Post",
		Slug:      "first-post",
		Html:      "<p>First Post</p>",
		Image:     "image.png",
		Published: true,
	}

	type args struct {
		ctx context.Context
		req *postsv1alpha2.CreatePostRequest
	}
	tests := []struct {
		name               string
		args               args
		expectedCreatePost *gomock.Call
		expectedGetPost    *gomock.Call
		want               *postsv1alpha2.Post
		wantErr            bool
	}{
		{
			name: "Test create post",
			args: args{
				context.Background(),
				&postsv1alpha2.CreatePostRequest{
					Parent: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
					Post:   testpost,
				},
			},
			expectedCreatePost: mockDatastore.EXPECT().Create("mysites.site", "testauthor", testpost).Return(nil),
			expectedGetPost:    mockDatastore.EXPECT().Get("mysites.site", "testauthor", testpost.Slug).Return(testpost, nil),
			want:               testpost,
		},
		{
			name: "Test create post with nil request",
			args: args{
				context.Background(),
				&postsv1alpha2.CreatePostRequest{},
			},
			wantErr: true,
		},
		{
			name: "Test create post with empty title",
			args: args{
				context.Background(),
				&postsv1alpha2.CreatePostRequest{
					Parent: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
					Post: &postsv1alpha2.Post{
						Slug:      "first-post",
						Html:      "<p>First Post</p>",
						Image:     "image.png",
						Published: true,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create post with empty slug",
			args: args{
				context.Background(),
				&postsv1alpha2.CreatePostRequest{
					Parent: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
					Post: &postsv1alpha2.Post{
						Title:     "First Post",
						Html:      "<p>First Post</p>",
						Image:     "image.png",
						Published: true,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create post with empty html body",
			args: args{
				context.Background(),
				&postsv1alpha2.CreatePostRequest{
					Parent: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor"}, "/"),
					Post: &postsv1alpha2.Post{
						Title:     "First Post",
						Slug:      "first-post",
						Image:     "image.png",
						Published: true,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.CreatePost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatastore := NewMockPostDatastore(ctrl)
	pageToken := util.NewPageToken()

	testpost := &postsv1alpha2.Post{
		Name:      strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "first-post"}, "/"),
		Title:     "First Post",
		Slug:      "first-post",
		Html:      "<p>First Post</p>",
		Image:     "image.png",
		Published: true,
	}

	type args struct {
		ctx context.Context
		req *postsv1alpha2.GetPostRequest
	}
	tests := []struct {
		name            string
		args            args
		expectedGetPost *gomock.Call
		want            *postsv1alpha2.Post
		wantErr         bool
	}{
		{
			name: "Test get post",
			args: args{
				context.Background(),
				&postsv1alpha2.GetPostRequest{
					Name: testpost.Name,
				},
			},
			expectedGetPost: mockDatastore.EXPECT().Get("mysites.site", "testauthor", testpost.Slug).Return(testpost, nil),
			want:            testpost,
		},
		{
			name: "Test get post not found",
			args: args{
				context.Background(),
				&postsv1alpha2.GetPostRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "notfound"}, "/"),
				},
			},
			expectedGetPost: mockDatastore.EXPECT().Get("mysites.site", "testauthor", "notfound").Return(nil, status.Error(codes.NotFound, "post not found")),
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.GetPost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatastore := NewMockPostDatastore(ctrl)
	pageToken := util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha2.UpdatePostRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *postsv1alpha2.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.UpdatePost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatastore := NewMockPostDatastore(ctrl)
	pageToken := util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha2.DeletePostRequest
	}
	tests := []struct {
		name               string
		args               args
		expectedDeletePost *gomock.Call
		want               *empty.Empty
		wantErr            bool
	}{
		{
			name: "Test delete post",
			args: args{
				context.Background(),
				&postsv1alpha2.DeletePostRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "first-post"}, "/"),
				},
			},
			expectedDeletePost: mockDatastore.EXPECT().Delete("mysites.site", "testauthor", "first-post").Return(nil),
			want:               &empty.Empty{},
		},
		{
			name: "Test delete post not found",
			args: args{
				context.Background(),
				&postsv1alpha2.DeletePostRequest{
					Name: strings.Join([]string{"sites", "mysites.site", "authors", "testauthor", "posts", "notfound"}, "/"),
				},
			},
			expectedDeletePost: mockDatastore.EXPECT().Delete("mysites.site", "testauthor", "notfound").Return(status.Error(codes.NotFound, "post not found")),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.DeletePost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeletePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_SearchPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDatastore := NewMockPostDatastore(ctrl)
	pageToken := util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha2.SearchPostRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *postsv1alpha2.PostList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, pageToken)
			got, err := server.SearchPost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.SearchPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.SearchPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
