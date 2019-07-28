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
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	postsv1alpha2 "github.com/prksu/publr/pkg/api/posts/v1alpha2"
	"github.com/prksu/publr/pkg/storage/database"
)

// PostDatastore interface
type PostDatastore interface {
	List(ctx context.Context, sitedomain, author string, start, limit int) ([]*postsv1alpha2.Post, int, error)
	Create(ctx context.Context, sitedomain, author string, post *postsv1alpha2.Post) error
	Get(ctx context.Context, sitedomain, author, slug string) (*postsv1alpha2.Post, error)
	Update(ctx context.Context, sitedomain, author, slug string, post *postsv1alpha2.Post) error
	Delete(ctx context.Context, sitedomain, author, slug string) error
	Search(ctx context.Context, query string) ([]*postsv1alpha2.Post, error)
}

// datastore implement users service datastore
type datastore struct {
	DB *sql.DB
}

// NewPostDatastore create new posts service datastore instance with configured database connection.
func NewPostDatastore() PostDatastore {
	return NewPostDatastoreWithDB(database.NewDatabase().Connect())
}

// NewPostDatastoreWithDB create new posts service datastore instance with sql.DB params.
func NewPostDatastoreWithDB(database *sql.DB) PostDatastore {
	ds := new(datastore)
	ds.DB = database
	return ds
}

func (ds *datastore) List(ctx context.Context, sitedomain, author string, start, limit int) ([]*postsv1alpha2.Post, int, error) {
	var posts []*postsv1alpha2.Post
	var foundRows int

	sqlrows := `
		SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
		FROM posts AS p
		LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
		LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
		WHERE ps.site_domain = ? AND ( NULLIF(?, '') IS NULL OR pa.author_username = ? )
		LIMIT ?, ?
	`

	sqlcount := `
		SELECT COUNT(*)
		FROM posts AS p
		LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
		LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
		WHERE ps.site_domain = ? AND ( NULLIF(?, '') IS NULL OR pa.author_username = ? )
	`

	rows, err := ds.DB.QueryContext(ctx, sqlrows, sitedomain, author, author, start, limit)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	for rows.Next() {
		var post postsv1alpha2.Post
		var createTime mysql.NullTime
		var publishTime mysql.NullTime
		var updateTime mysql.NullTime
		if err := rows.Scan(&post.Title, &post.Slug, &post.Html, &post.Image, &post.Published, &createTime, &publishTime, &updateTime); err != nil {
			return nil, 0, err
		}

		post.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
		post.PublishTime, _ = ptypes.TimestampProto(publishTime.Time)
		post.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)
		posts = append(posts, &post)
	}

	if err := ds.DB.QueryRow(sqlcount, sitedomain, author, author).Scan(&foundRows); err != nil {
		return nil, 0, err
	}

	return posts, foundRows, nil
}

func (ds *datastore) Create(ctx context.Context, sitedomain, author string, post *postsv1alpha2.Post) error {
	tx, err := ds.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	{
		stmt, err := tx.PrepareContext(ctx, "INSERT INTO posts (title, slug, html, image, published) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.ExecContext(ctx, post.Title, post.Slug, post.Html, post.Image, post.Published); err != nil {
			tx.Rollback()
			return err
		}
	}
	{
		stmt, err := tx.PrepareContext(ctx, "INSERT INTO post_sites (post_slug, site_domain) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.ExecContext(ctx, post.Slug, sitedomain); err != nil {
			tx.Rollback()
			if err.(*mysql.MySQLError).Number == 1062 {
				return status.Error(codes.AlreadyExists, "post slug already taken")
			}
			return err
		}
	}
	{
		stmt, err := tx.PrepareContext(ctx, "INSERT INTO post_authors (post_slug, author_username) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.ExecContext(ctx, post.Slug, author); err != nil {
			tx.Rollback()
			if err.(*mysql.MySQLError).Number == 1062 {
				return status.Error(codes.AlreadyExists, "cannot add author in this post")
			}
			return err
		}
	}

	return tx.Commit()
}

func (ds *datastore) Get(ctx context.Context, sitedomain, author, slug string) (*postsv1alpha2.Post, error) {
	post := new(postsv1alpha2.Post)
	var sqlrow *sql.Row

	sqlquery := `
		SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
		FROM posts AS p
		LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
		LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
		WHERE ps.site_domain=? AND p.slug=? AND ( NULLIF(?, '') IS NULL OR pa.author_username = ? )
	`

	sqlrow = ds.DB.QueryRowContext(ctx, sqlquery, sitedomain, slug, author, author)

	var createTime mysql.NullTime
	var publishTime mysql.NullTime
	var updateTime mysql.NullTime
	if err := sqlrow.Scan(&post.Title, &post.Slug, &post.Html, &post.Image, &post.Published, &createTime, &publishTime, &updateTime); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "post not found")
		}
		return nil, err
	}

	post.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
	post.PublishTime, _ = ptypes.TimestampProto(publishTime.Time)
	post.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)

	return post, nil
}

func (ds *datastore) Update(ctx context.Context, sitedomain, author, slug string, post *postsv1alpha2.Post) error {
	return nil
}

func (ds *datastore) Delete(ctx context.Context, sitedomain, author, slug string) error {
	var sqlquery = `
		DELETE posts FROM posts
		LEFT JOIN post_sites on posts.slug=post_sites.post_slug
		LEFT JOIN post_authors on posts.slug=post_authors.post_slug
		WHERE post_sites.site_domain=? AND post_authors.author_username=? AND posts.slug=?
	`

	result, err := ds.DB.ExecContext(ctx, sqlquery, sitedomain, author, slug)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return status.Error(codes.NotFound, "post not found")
	}

	return nil
}

func (ds *datastore) Search(ctx context.Context, query string) ([]*postsv1alpha2.Post, error) {
	return nil, nil
}
