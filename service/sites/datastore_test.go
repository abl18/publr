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
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/protobuf/ptypes"

	sitesv1alpha3 "github.com/prksu/publr/pkg/api/sites/v1alpha3"
)

func Test_datastore_Create(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	site := &sitesv1alpha3.Site{
		Title:  "My Sites",
		Domain: "mysites.site",
	}

	pre := mock.ExpectPrepare("INSERT INTO")

	type args struct {
		context context.Context
		site    *sitesv1alpha3.Site
	}
	tests := []struct {
		name         string
		args         args
		expectedExec *sqlmock.ExpectedExec
		wantErr      bool
	}{
		{
			name: "Test create sites",
			args: args{
				context: ctx,
				site:    site,
			},
			expectedExec: pre.ExpectExec().WithArgs(site.Title, site.Domain).WillReturnResult(sqlmock.NewResult(1, 1)),
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewSiteDatastoreWithDB(database)
			if err := datastore.Create(tt.args.context, tt.args.site); (err != nil) != tt.wantErr {
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
	site := &sitesv1alpha3.Site{
		Title:      "My Sites",
		Domain:     "mysites.site",
		CreateTime: protoTimestamp,
		UpdateTime: protoTimestamp,
	}

	rows := sqlmock.NewRows([]string{"title", "domain", "createtime", "updatetime"}).
		AddRow(site.Title, site.Domain, timestamp, timestamp)

	type args struct {
		ctx        context.Context
		sitedomain string
	}
	tests := []struct {
		name          string
		args          args
		expectedQuery *sqlmock.ExpectedQuery
		want          *sitesv1alpha3.Site
		wantErr       bool
	}{
		{
			name: "Test get site",
			args: args{
				ctx:        ctx,
				sitedomain: site.Domain,
			},
			expectedQuery: mock.ExpectQuery("SELECT title, domain, createtime, updatetime FROM sites").WithArgs(site.Domain).WillReturnRows(rows),
			want:          site,
			wantErr:       false,
		},
		{
			name: "Test get site not found",
			args: args{
				ctx:        ctx,
				sitedomain: "notfound",
			},
			expectedQuery: mock.ExpectQuery("SELECT title, domain, createtime, updatetime FROM sites").WithArgs("notfound").WillReturnError(sql.ErrNoRows),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewSiteDatastoreWithDB(database)
			got, err := datastore.Get(tt.args.ctx, tt.args.sitedomain)
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

func Test_datastore_Delete(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	sitedomain := "mysites.site"

	type args struct {
		context    context.Context
		sitedomain string
	}
	tests := []struct {
		name         string
		args         args
		expectedExec *sqlmock.ExpectedExec
		wantErr      bool
	}{
		{
			name: "Test delete site",
			args: args{
				context:    ctx,
				sitedomain: sitedomain,
			},
			expectedExec: mock.ExpectExec("DELETE FROM sites").WithArgs(sitedomain).WillReturnResult(sqlmock.NewResult(0, 1)),
		},
		{
			name: "Test delete site not found",
			args: args{
				context:    ctx,
				sitedomain: "notfound",
			},
			expectedExec: mock.ExpectExec("DELETE FROM sites").WithArgs("notfound").WillReturnResult(sqlmock.NewResult(0, 0)),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			datastore := NewSiteDatastoreWithDB(database)
			if err := datastore.Delete(tt.args.context, tt.args.sitedomain); (err != nil) != tt.wantErr {
				t.Errorf("datastore.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
