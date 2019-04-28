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

package sites

import (
	"database/sql"

	sitesv1alpha1 "github.com/prksu/publr/pkg/api/sites/v1alpha1"
	"github.com/prksu/publr/pkg/storage/database"
)

// SiteDatastore interface
type SiteDatastore interface {
	Create(site *sitesv1alpha1.Site) error
	Get(sitedomain string) (*sitesv1alpha1.Site, error)
	Delete(sitedomain string) error
}

// datastore implement sites service datastore
type datastore struct {
	DB *sql.DB
}

// NewDatastore create new sites service datastore instance
func NewDatastore() SiteDatastore {
	ds := new(datastore)
	ds.DB = database.NewDatabase().Connect()
	return ds
}

func (ds *datastore) Create(site *sitesv1alpha1.Site) error { return nil }

func (ds *datastore) Get(sitedomain string) (*sitesv1alpha1.Site, error) { return nil, nil }

func (ds *datastore) Delete(sitedomain string) error { return nil }
