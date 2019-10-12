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
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/protobuf/ptypes"

	postsv1alpha3 "github.com/prksu/publr/pkg/api/posts/v1alpha3"
)

func Test_datastore_List(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	timestamp := time.Now()
	protoTimestamp, err := ptypes.TimestampProto(timestamp)
	if err != nil {
		t.Fatal(err)
	}

	sitedomain := "mysites.site"
	author := "testauthor"
	posts := []*postsv1alpha3.Post{
		{
			Title:       "First Post",
			Slug:        "first-post",
			Html:        "<p>First Post</p>",
			Image:       "image.png",
			Published:   true,
			CreateTime:  protoTimestamp,
			PublishTime: protoTimestamp,
			UpdateTime:  protoTimestamp,
		},
		{
			Title:       "Second Post",
			Slug:        "second-post",
			Html:        "<p>Second Post</p>",
			Published:   true,
			CreateTime:  protoTimestamp,
			PublishTime: protoTimestamp,
			UpdateTime:  protoTimestamp,
		},
		{
			Title:       "Third Post",
			Slug:        "third-post",
			Html:        "<p>Third Post</p>",
			Published:   true,
			CreateTime:  protoTimestamp,
			PublishTime: protoTimestamp,
			UpdateTime:  protoTimestamp,
		},
		{
			Title:       "Fourth Post",
			Slug:        "fourth-post",
			Html:        "<p>Fourth Post</p>",
			Published:   false,
			CreateTime:  protoTimestamp,
			PublishTime: protoTimestamp,
			UpdateTime:  protoTimestamp,
		},
	}

	rows := new(sqlmock.Rows)
	columns := sqlmock.NewRows([]string{"title", "slug", "html", "imgage", "published", "createtime", "publishtime", "updatetime"})
	for _, p := range posts {
		rows = columns.AddRow(p.Title, p.Slug, p.Html, p.Image, p.Published, timestamp, timestamp, timestamp)
	}

	sqlrows := `
		SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
		FROM posts AS p
		LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
		LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
	`

	sqlcount := `
		SELECT COUNT(.+)
		FROM posts AS p
		LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
		LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
	`

	type args struct {
		context    context.Context
		sitedomain string
		author     string
		start      int
		limit      int
	}
	tests := []struct {
		name               string
		args               args
		expectedQuery      *sqlmock.ExpectedQuery
		expectedCountQuery *sqlmock.ExpectedQuery
		want               []*postsv1alpha3.Post
		want1              int
		wantErr            bool
	}{
		{
			name: "Test list post",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
				author:     author,
				start:      0,
				limit:      10,
			},
			expectedQuery:      mock.ExpectQuery(sqlrows).WithArgs(sitedomain, author, author, 0, 10).WillReturnRows(rows),
			expectedCountQuery: mock.ExpectQuery(sqlcount).WithArgs(sitedomain, author, author).WillReturnRows(sqlmock.NewRows([]string{"found_rows"}).AddRow(len(posts))),
			want:               posts,
			want1:              len(posts),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewPostDatastoreWithDB(database)
			got, got1, err := datastore.List(tt.args.context, tt.args.sitedomain, tt.args.author, tt.args.start, tt.args.limit)
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

	ctx := context.Background()
	sitedomain := "mysites.site"
	author := "testauthor"

	post := &postsv1alpha3.Post{
		Title:     "First Post",
		Slug:      "first-post",
		Html:      "<p>First Post</p>",
		Image:     "image.png",
		Published: true,
	}

	type args struct {
		context    context.Context
		sitedomain string
		author     string
		post       *postsv1alpha3.Post
	}
	tests := []struct {
		name             string
		args             args
		expectedBegin    *sqlmock.ExpectedBegin
		expectedExec1    *sqlmock.ExpectedExec
		expectedExec2    *sqlmock.ExpectedExec
		expectedExec3    *sqlmock.ExpectedExec
		expectedRollback *sqlmock.ExpectedRollback
		expectedCommit   *sqlmock.ExpectedCommit
		wantErr          bool
	}{
		{
			name: "Test create post",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
				author:     author,
				post:       post,
			},
			expectedBegin:  mock.ExpectBegin(),
			expectedExec1:  mock.ExpectPrepare("INSERT INTO posts").ExpectExec().WithArgs(post.Title, post.Slug, post.Html, post.Image, post.Published).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedExec2:  mock.ExpectPrepare("INSERT INTO post_sites").ExpectExec().WithArgs(post.Slug, sitedomain).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedExec3:  mock.ExpectPrepare("INSERT INTO post_authors").ExpectExec().WithArgs(post.Slug, author).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedCommit: mock.ExpectCommit(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewPostDatastoreWithDB(database)
			if err := datastore.Create(tt.args.context, tt.args.sitedomain, tt.args.author, tt.args.post); (err != nil) != tt.wantErr {
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

	ctx := context.Background()
	timestamp := time.Now()
	protoTimestamp, err := ptypes.TimestampProto(timestamp)
	if err != nil {
		t.Fatal(err)
	}

	post := &postsv1alpha3.Post{
		Title:       "First Post",
		Slug:        "first-post",
		Html:        "<p>First Post</p>",
		Image:       "image.png",
		Published:   true,
		CreateTime:  protoTimestamp,
		PublishTime: protoTimestamp,
		UpdateTime:  protoTimestamp,
	}

	sitedomain := "mysites.site"
	author := "testauthor"
	slug := "first-post"

	sqlquery := `
		SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
		FROM posts AS p
		LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
		LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
	`

	rows := sqlmock.NewRows([]string{"title", "slug", "html", "imgage", "published", "createtime", "publishtime", "updatetime"}).
		AddRow(post.Title, post.Slug, post.Html, post.Image, post.Published, timestamp, timestamp, timestamp)

	type args struct {
		context    context.Context
		sitedomain string
		author     string
		slug       string
	}
	tests := []struct {
		name          string
		args          args
		expectedQuery *sqlmock.ExpectedQuery
		want          *postsv1alpha3.Post
		wantErr       bool
	}{
		{
			name: "Test get post",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
				author:     author,
				slug:       slug,
			},
			expectedQuery: mock.ExpectQuery(sqlquery).WithArgs(sitedomain, slug, author, author).WillReturnRows(rows),
			want:          post,
		},
		{
			name: "Test get post not found",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
				author:     author,
				slug:       "notfound",
			},
			expectedQuery: mock.ExpectQuery(sqlquery).WithArgs(sitedomain, "notfound", author, author).WillReturnError(sql.ErrNoRows),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewPostDatastoreWithDB(database)
			got, err := datastore.Get(tt.args.context, tt.args.sitedomain, tt.args.author, tt.args.slug)
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
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		context    context.Context
		sitedomain string
		author     string
		slug       string
		post       *postsv1alpha3.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &datastore{
				DB: tt.fields.DB,
			}
			if err := ds.Update(tt.args.context, tt.args.sitedomain, tt.args.author, tt.args.slug, tt.args.post); (err != nil) != tt.wantErr {
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

	ctx := context.Background()
	sitedomain := "mysites.site"
	author := "testauthor"
	slug := "first-post"

	var sqlquery = `
		DELETE posts FROM posts
		LEFT JOIN post_sites on posts.slug=post_sites.post_slug
		LEFT JOIN post_authors on posts.slug=post_authors.post_slug
	`

	type args struct {
		context    context.Context
		sitedomain string
		author     string
		slug       string
	}
	tests := []struct {
		name         string
		args         args
		expectedExec *sqlmock.ExpectedExec
		wantErr      bool
	}{
		{
			name: "Test delete post",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
				author:     author,
				slug:       slug,
			},
			expectedExec: mock.ExpectExec(sqlquery).WithArgs(sitedomain, author, slug).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name: "Test delete post not found",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
				author:     author,
				slug:       "notfound",
			},
			expectedExec: mock.ExpectExec(sqlquery).WithArgs(sitedomain, author, "notfound").WillReturnResult(sqlmock.NewResult(0, 0)),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewPostDatastoreWithDB(database)
			if err := datastore.Delete(tt.args.context, tt.args.sitedomain, tt.args.author, tt.args.slug); (err != nil) != tt.wantErr {
				t.Errorf("datastore.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_datastore_Search(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		context context.Context
		query   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*postsv1alpha3.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &datastore{
				DB: tt.fields.DB,
			}
			got, err := ds.Search(tt.args.context, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("datastore.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("datastore.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
