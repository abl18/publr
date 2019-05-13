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
	"log"
	"reflect"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	postsv1alpha1 "github.com/prksu/publr/pkg/api/posts/v1alpha1"
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

	database.Exec("DROP TABLE IF EXISTS post_sites")
	database.Exec("DROP TABLE IF EXISTS post_authors")
	database.Exec("DROP TABLE IF EXISTS posts")
	schema, err := schema.Asset("data/schema/posts.sql")
	if err != nil {
		log.Fatal(err)
	}

	testdata, err := testdata.Asset("data/testdata/posts.sql")
	if err != nil {
		log.Fatal(err)
	}

	database.Exec(string(schema))
	database.Exec(string(testdata))
}

func TestServer_ListPost(t *testing.T) {
	server := new(Server)
	server.Post = NewPostDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha1.ListPostRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *postsv1alpha1.PostList
		wantErr bool
	}{
		{
			name: "Test list site posts",
			args: args{
				context.Background(),
				&postsv1alpha1.ListPostRequest{
					Parent: "sites/mysites.site",
				},
			},
			want: &postsv1alpha1.PostList{
				Posts: []*postsv1alpha1.Post{
					{
						Name:      "sites/mysites.site/posts/my-first-posts",
						Title:     "My First Post",
						Slug:      "my-first-posts",
						Html:      "<p>My First Post</p>",
						Image:     "image.png",
						Published: true,
					},
					{
						Name:      "sites/mysites.site/posts/my-second-posts",
						Title:     "My Second Post",
						Slug:      "my-second-posts",
						Html:      "<p>My Second Post</p>",
						Image:     "image.png",
						Published: true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test list author posts",
			args: args{
				context.Background(),
				&postsv1alpha1.ListPostRequest{
					Parent: "sites/mysites.site/authors/authordemo",
				},
			},
			want: &postsv1alpha1.PostList{
				Posts: []*postsv1alpha1.Post{
					{
						Name:      "sites/mysites.site/authors/authordemo/posts/my-first-posts",
						Title:     "My First Post",
						Slug:      "my-first-posts",
						Html:      "<p>My First Post</p>",
						Image:     "image.png",
						Published: true,
					},
					{
						Name:      "sites/mysites.site/authors/authordemo/posts/my-second-posts",
						Title:     "My Second Post",
						Slug:      "my-second-posts",
						Html:      "<p>My Second Post</p>",
						Image:     "image.png",
						Published: true,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.ListPost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ListPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				for _, user := range tt.want.Posts {
					for _, g := range got.Posts {
						user.CreateTime = g.CreateTime
						user.PublishTime = g.PublishTime
						user.UpdateTime = g.UpdateTime
					}
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ListPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CreatePost(t *testing.T) {
	server := new(Server)
	server.Post = NewPostDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha1.CreatePostRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *postsv1alpha1.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	server := new(Server)
	server.Post = NewPostDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha1.GetPostRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *postsv1alpha1.Post
		wantErr bool
	}{
		{
			name: "Test get post",
			args: args{
				context.Background(),
				&postsv1alpha1.GetPostRequest{
					Name: "sites/mysites.site/posts/my-first-posts",
				},
			},
			want: &postsv1alpha1.Post{
				Name:      "sites/mysites.site/posts/my-first-posts",
				Title:     "My First Post",
				Slug:      "my-first-posts",
				Html:      "<p>My First Post</p>",
				Image:     "image.png",
				Published: true,
			},
			wantErr: false,
		},
		{
			name: "Test get author post",
			args: args{
				context.Background(),
				&postsv1alpha1.GetPostRequest{
					Name: "sites/mysites.site/authors/authordemo/posts/my-second-posts",
				},
			},
			want: &postsv1alpha1.Post{
				Name:      "sites/mysites.site/authors/authordemo/posts/my-second-posts",
				Title:     "My Second Post",
				Slug:      "my-second-posts",
				Html:      "<p>My Second Post</p>",
				Image:     "image.png",
				Published: true,
			},
			wantErr: false,
		},
		{
			name: "Test get not existing post",
			args: args{
				context.Background(),
				&postsv1alpha1.GetPostRequest{
					Name: "sites/mysites.site/posts/notexists",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.GetPost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				tt.want.CreateTime = got.CreateTime
				tt.want.PublishTime = got.PublishTime
				tt.want.UpdateTime = got.UpdateTime
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeletePost(t *testing.T) {
	server := new(Server)
	server.Post = NewPostDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	server.PageToken = util.NewPageToken()

	type args struct {
		ctx context.Context
		req *postsv1alpha1.DeletePostRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			name: "Test delete  post",
			args: args{
				context.Background(),
				&postsv1alpha1.DeletePostRequest{
					Name: "sites/mysites.site/authors/authordemo/posts/my-second-posts",
				},
			},
			want:    &empty.Empty{},
			wantErr: false,
		},
		{
			name: "Test delete not existing post",
			args: args{
				context.Background(),
				&postsv1alpha1.DeletePostRequest{
					Name: "sites/mysites.site/authors/authordemo/posts/my-second-posts",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
