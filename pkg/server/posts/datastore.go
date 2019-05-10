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
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	postsv1alpha1 "github.com/prksu/publr/pkg/api/posts/v1alpha1"
	"github.com/prksu/publr/pkg/storage/database"
)

// PostDatastore interface
type PostDatastore interface {
	List(sitedomain, author string, start, offset int) ([]*postsv1alpha1.Post, error)
	Create(sitedomain, author string, post *postsv1alpha1.Post) error
	Get(sitedomain, author, slug string) (*postsv1alpha1.Post, error)
	Update(sitedomain, author, slug string, post *postsv1alpha1.Post) error
	Delete(sitedomain, author, slug string) error
	Search(query string) ([]*postsv1alpha1.Post, error)
}

// datastore implement users service datastore
type datastore struct {
	DB *sql.DB
}

// NewPostDatastore create new users service datastore instance
func NewPostDatastore() PostDatastore {
	ds := new(datastore)
	ds.DB = database.NewDatabase().Connect()
	return ds
}

func (ds *datastore) List(sitedomain, author string, start, offset int) ([]*postsv1alpha1.Post, error) {
	var posts []*postsv1alpha1.Post
	var sqlrows *sql.Rows
	var err error

	if author == "" {
		sqlquery := `
			SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
			FROM posts AS p
			LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
			WHERE ps.site_domain=?
			LIMIT ?, ?
		`
		sqlrows, err = ds.DB.Query(sqlquery, sitedomain, start, offset)
		if err != nil {
			return nil, err
		}

	} else {
		sqlquery := `
			SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
			FROM posts AS p
			LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
			LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
			WHERE ps.site_domain=? AND pa.author_username=?
			LIMIT ?, ?
		`
		sqlrows, err = ds.DB.Query(sqlquery, sitedomain, author, start, offset)
		if err != nil {
			return nil, err
		}

	}

	defer sqlrows.Close()
	for sqlrows.Next() {
		var post postsv1alpha1.Post
		var createTime mysql.NullTime
		var publishTime mysql.NullTime
		var updateTime mysql.NullTime
		if err := sqlrows.Scan(&post.Title, &post.Slug, &post.Html, &post.Image, &post.Published, &createTime, &publishTime, &updateTime); err != nil {
			return nil, err
		}

		post.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
		post.PublishTime, _ = ptypes.TimestampProto(publishTime.Time)
		post.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)
		posts = append(posts, &post)
	}

	return posts, nil
}

func (ds *datastore) Create(sitedomain, author string, post *postsv1alpha1.Post) error {
	tx, err := ds.DB.Begin()
	if err != nil {
		return err
	}
	{
		stmt, err := tx.Prepare("INSERT INTO posts (title, slug, html, image, published) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(post.Title, post.Slug, post.Html, post.Image, post.Published); err != nil {
			tx.Rollback()
			return err
		}
	}
	{
		stmt, err := tx.Prepare("INSERT INTO post_sites (post_slug, site_domain) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(post.Slug, sitedomain); err != nil {
			tx.Rollback()
			if err.(*mysql.MySQLError).Number == 1062 {
				return status.Error(codes.AlreadyExists, "post slug already taken")
			}
			return err
		}
	}
	{
		stmt, err := tx.Prepare("INSERT INTO post_authors (post_slug, author_username) VALUES (?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(post.Slug, author); err != nil {
			tx.Rollback()
			if err.(*mysql.MySQLError).Number == 1062 {
				return status.Error(codes.AlreadyExists, "cannot add author in this post")
			}
			return err
		}
	}

	return tx.Commit()
}

func (ds *datastore) Get(sitedomain, author, slug string) (*postsv1alpha1.Post, error) {
	post := new(postsv1alpha1.Post)
	var sqlrow *sql.Row

	if author == "" {
		sqlquery := `
			SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
			FROM posts AS p
			LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
			WHERE ps.site_domain=? AND p.slug=?
		`
		sqlrow = ds.DB.QueryRow(sqlquery, sitedomain, slug)

	} else {
		sqlquery := `
			SELECT p.title, p.slug, p.html, p.image, p.published, p.createtime, p.publishtime, p.updatetime
			FROM posts AS p
			LEFT JOIN post_sites AS ps on p.slug=ps.post_slug
			LEFT JOIN post_authors AS pa on p.slug=pa.post_slug
			WHERE ps.site_domain=? AND pa.author_username=? AND p.slug=?
		`
		sqlrow = ds.DB.QueryRow(sqlquery, sitedomain, author, slug)
	}

	var createTime mysql.NullTime
	var publishTime mysql.NullTime
	var updateTime mysql.NullTime
	if err := sqlrow.Scan(&post.Title, &post.Slug, &post.Html, &post.Image, &post.Published, &createTime, &publishTime, &updateTime); err != nil {
		return nil, err
	}

	post.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
	post.PublishTime, _ = ptypes.TimestampProto(publishTime.Time)
	post.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)

	return post, nil
}

func (ds *datastore) Update(sitedomain, author, slug string, post *postsv1alpha1.Post) error {
	return nil
}

func (ds *datastore) Delete(sitedomain, author, slug string) error {
	var sqlquery = `
		DELETE posts FROM posts
		LEFT JOIN post_sites on posts.slug=post_sites.post_slug
		LEFT JOIN post_authors on posts.slug=post_authors.post_slug
		WHERE post_sites.site_domain=? AND post_authors.author_username=? AND posts.slug=?
	`

	if _, err := ds.DB.Exec(sqlquery, sitedomain, author, slug); err != nil {
		return err
	}
	return nil
}

func (ds *datastore) Search(query string) ([]*postsv1alpha1.Post, error) { return nil, nil }
