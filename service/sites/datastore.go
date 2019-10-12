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
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sitesv1alpha3 "github.com/prksu/publr/pkg/api/sites/v1alpha3"
	"github.com/prksu/publr/pkg/storage/database"
)

// SiteDatastore interface
type SiteDatastore interface {
	Create(ctx context.Context, site *sitesv1alpha3.Site) error
	Get(ctx context.Context, sitedomain string) (*sitesv1alpha3.Site, error)
	Delete(ctx context.Context, sitedomain string) error
}

// datastore implement sites service datastore
type datastore struct {
	DB *sql.DB
}

// NewSiteDatastore create new sites service datastore instance with configured database connection.
func NewSiteDatastore() SiteDatastore {
	return NewSiteDatastoreWithDB(database.NewDatabase().Connect())
}

// NewSiteDatastoreWithDB create new sites service datastore instance with sql.DB params.
func NewSiteDatastoreWithDB(database *sql.DB) SiteDatastore {
	ds := new(datastore)
	ds.DB = database
	return ds
}

func (ds *datastore) Create(ctx context.Context, site *sitesv1alpha3.Site) error {
	stmt, err := ds.DB.PrepareContext(ctx, "INSERT INTO sites (title, domain) VALUES (?, ?)")
	if err != nil {
		return err
	}

	if _, err = stmt.ExecContext(ctx, site.Title, site.Domain); err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			return status.Error(codes.AlreadyExists, "site already exists")
		}
		return err
	}

	return nil
}

func (ds *datastore) Get(ctx context.Context, sitedomain string) (*sitesv1alpha3.Site, error) {
	site := new(sitesv1alpha3.Site)
	var createTime mysql.NullTime
	var updateTime mysql.NullTime
	if err := ds.DB.QueryRowContext(ctx, "SELECT title, domain, createtime, updatetime FROM sites WHERE domain=?", sitedomain).
		Scan(&site.Title, &site.Domain, &createTime, &updateTime); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "site not found")
		}
		return nil, err
	}

	site.CreateTime, _ = ptypes.TimestampProto(createTime.Time)
	site.UpdateTime, _ = ptypes.TimestampProto(updateTime.Time)
	return site, nil
}

func (ds *datastore) Delete(ctx context.Context, sitedomain string) error {
	result, err := ds.DB.ExecContext(ctx, "DELETE FROM sites WHERE domain=?", sitedomain)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return status.Error(codes.NotFound, "site not found")
	}

	return nil
}
