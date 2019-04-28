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

	postsv1alpha1 "github.com/prksu/publr/pkg/api/posts/v1alpha1"
	"github.com/prksu/publr/pkg/storage/database"
)

// PostDatastore interface
type PostDatastore interface {
	List() ([]*postsv1alpha1.Post, error)
	Create(sitedomain string, post *postsv1alpha1.Post) error
	Get(sitedomain, username string) (*postsv1alpha1.Post, error)
	Update(sitedomain, username string, post *postsv1alpha1.Post) error
	Delete(sitedomain, username string) error
	Search(query string) ([]*postsv1alpha1.Post, error)
}

// datastore implement users service datastore
type datastore struct {
	DB *sql.DB
}

// NewDatastore create new users service datastore instance
func NewDatastore() PostDatastore {
	ds := new(datastore)
	ds.DB = database.NewDatabase().Connect()
	return ds
}

func (ds *datastore) List() ([]*postsv1alpha1.Post, error) { return nil, nil }

func (ds *datastore) Create(sitedomain string, post *postsv1alpha1.Post) error { return nil }

func (ds *datastore) Get(sitedomain, username string) (*postsv1alpha1.Post, error) { return nil, nil }

func (ds *datastore) Update(sitedomain, username string, post *postsv1alpha1.Post) error { return nil }

func (ds *datastore) Delete(sitedomain, username string) error { return nil }

func (ds *datastore) Search(query string) ([]*postsv1alpha1.Post, error) { return nil, nil }
