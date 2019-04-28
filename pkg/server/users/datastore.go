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
	"database/sql"

	usersv1alpha1 "github.com/prksu/publr/pkg/api/users/v1alpha1"
	"github.com/prksu/publr/pkg/storage/database"
)

// UserDatastore interface
type UserDatastore interface {
	List() ([]*usersv1alpha1.User, error)
	Create(sitedomain string, user *usersv1alpha1.User) error
	Get(sitedomain, username string) (*usersv1alpha1.User, error)
	Update(sitedomain, username string, user *usersv1alpha1.User) error
	Delete(sitedomain, username string) error
	Search(query string) ([]*usersv1alpha1.User, error)
}

// datastore implement users service datastore
type datastore struct {
	DB *sql.DB
}

// NewDatastore create new users service datastore instance
func NewDatastore() UserDatastore {
	ds := new(datastore)
	ds.DB = database.NewDatabase().Connect()
	return ds
}

func (ds *datastore) List() ([]*usersv1alpha1.User, error) { return nil, nil }

func (ds *datastore) Create(sitedomain string, user *usersv1alpha1.User) error { return nil }

func (ds *datastore) Get(sitedomain, username string) (*usersv1alpha1.User, error) { return nil, nil }

func (ds *datastore) Update(sitedomain, username string, user *usersv1alpha1.User) error { return nil }

func (ds *datastore) Delete(sitedomain, username string) error { return nil }

func (ds *datastore) Search(query string) ([]*usersv1alpha1.User, error) { return nil, nil }
