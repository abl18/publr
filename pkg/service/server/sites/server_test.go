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
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"

	sitesv1alpha2 "github.com/prksu/publr/pkg/api/sites/v1alpha2"
	usersv1alpha2 "github.com/prksu/publr/pkg/api/users/v1alpha2"
	"github.com/prksu/publr/pkg/bindata/schema"
	"github.com/prksu/publr/pkg/bindata/testdata"
	"github.com/prksu/publr/pkg/service/mock/users"
	"github.com/prksu/publr/pkg/storage/database"
)

var (
	DSN = "root:@/publr_test?autocommit=true&parseTime=true"
)

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func init() {
	database := database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect()
	defer database.Close()

	database.Exec("DROP TABLE IF EXISTS sites")
	schema, err := schema.Asset("data/schema/sites.sql")
	if err != nil {
		log.Fatal(err)
	}

	sitetestdata, err := testdata.Asset("data/testdata/sites.sql")
	if err != nil {
		log.Fatal(err)
	}
	usertestdata, err := testdata.Asset("data/testdata/users.sql")
	if err != nil {
		log.Fatal(err)
	}

	database.Exec(string(schema))
	database.Exec(string(sitetestdata))
	database.Exec(string(usertestdata))
}

func TestServer_CreateSite(t *testing.T) {
	server := new(Server)
	server.Site = NewSiteDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserClient := users.NewMockUserServiceClient(ctrl)
	server.UserClient = mockUserClient

	type args struct {
		ctx context.Context
		req *sitesv1alpha2.CreateSiteRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *sitesv1alpha2.Site
		wantErr bool
	}{
		{
			name: "Test create site",
			args: args{
				context.Background(),
				&sitesv1alpha2.CreateSiteRequest{
					Site: &sitesv1alpha2.Site{
						Title:  "My Awesome Sites",
						Domain: "myawesome.site",
						Owner: &usersv1alpha2.User{
							Username: "ownertest",
							Email:    "ownertest@myawesome.site",
							Password: "secret",
						},
					},
				},
			},
			want: &sitesv1alpha2.Site{
				Name:   "sites/myawesome.site",
				Title:  "My Awesome Sites",
				Domain: "myawesome.site",
				Owner: &usersv1alpha2.User{
					Username: "ownertest",
					Email:    "ownertest@myawesome.site",
					Role:     3,
				},
			},
			wantErr: false,
		},
		{
			name: "Test create already existing site",
			args: args{
				context.Background(),
				&sitesv1alpha2.CreateSiteRequest{
					Site: &sitesv1alpha2.Site{
						Title:  "My Sites",
						Domain: "mysites.site",
						Owner: &usersv1alpha2.User{
							Username: "ownertest",
							Email:    "ownertest@myawesome.site",
							Password: "secret",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create site with null request",
			args: args{
				context.Background(),
				&sitesv1alpha2.CreateSiteRequest{},
			},
			wantErr: true,
		},
		{
			name: "Test create site with empty title",
			args: args{
				context.Background(),
				&sitesv1alpha2.CreateSiteRequest{
					Site: &sitesv1alpha2.Site{
						Domain: "myawesome.site",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create site with empty domain",
			args: args{
				context.Background(),
				&sitesv1alpha2.CreateSiteRequest{
					Site: &sitesv1alpha2.Site{
						Title: "My Sites",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create site without owner",
			args: args{
				context.Background(),
				&sitesv1alpha2.CreateSiteRequest{
					Site: &sitesv1alpha2.Site{
						Title:  "My Sites",
						Domain: "mysites.site",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != nil {
				mockUserClient.EXPECT().
					CreateUser(gomock.Any(), &rpcMsg{msg: &usersv1alpha2.CreateUserRequest{Parent: strings.Join([]string{"sites", tt.args.req.Site.Domain}, "/"), User: tt.args.req.Site.Owner}}).
					Return(tt.want.Owner, nil)
			}

			got, err := server.CreateSite(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				tt.want.CreateTime = got.CreateTime
				tt.want.UpdateTime = got.UpdateTime
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetSite(t *testing.T) {
	server := new(Server)
	server.Site = NewSiteDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())

	tests := []struct {
		name    string
		request *sitesv1alpha2.GetSiteRequest
		want    *sitesv1alpha2.Site
		wantErr bool
	}{
		{
			name: "Test get site",
			request: &sitesv1alpha2.GetSiteRequest{
				Name: "sites/mysites.site",
			},
			want: &sitesv1alpha2.Site{
				Name:   "sites/mysites.site",
				Title:  "My Sites",
				Domain: "mysites.site",
			},
			wantErr: false,
		},
		{
			name: "Test get not existing site",
			request: &sitesv1alpha2.GetSiteRequest{
				Name: "sites/notexist",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.GetSite(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				tt.want.CreateTime = got.CreateTime
				tt.want.UpdateTime = got.UpdateTime
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteSite(t *testing.T) {
	server := new(Server)
	server.Site = NewSiteDatastoreWithDB(database.NewDatabase().WithDriver("mysql").WithDSN(DSN).Connect())

	tests := []struct {
		name    string
		request *sitesv1alpha2.DeleteSiteRequest
		want    *empty.Empty
		wantErr bool
	}{
		{
			name: "Test delete site",
			request: &sitesv1alpha2.DeleteSiteRequest{
				Name: "sites/mysites.site",
			},
			want:    &empty.Empty{},
			wantErr: false,
		},
		{
			name: "Test delete not existing site",
			request: &sitesv1alpha2.DeleteSiteRequest{
				Name: "sites/mysites.site",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.DeleteSite(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
