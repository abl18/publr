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
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sitesv1alpha3 "github.com/prksu/publr/pkg/api/sites/v1alpha3"
	usersv1alpha3 "github.com/prksu/publr/pkg/api/users/v1alpha3"
	mock_sites "github.com/prksu/publr/test/mock/sites"
	mock_users "github.com/prksu/publr/test/mock/users"
)

func TestServer_CreateSite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := mock_sites.NewMockSiteDatastore(ctrl)
	mockUserClient := mock_users.NewMockUserServiceClient(ctrl)

	testsite := &sitesv1alpha3.Site{
		Title:  "My Sites",
		Domain: "mysites.site",
		Owner: &usersv1alpha3.User{
			Username: "testowner",
			Email:    "testowner@mysites.site",
			Password: "secret",
			Fullname: "Test Owner",
		},
	}

	type args struct {
		ctx context.Context
		req *sitesv1alpha3.CreateSiteRequest
	}
	tests := []struct {
		name                string
		args                args
		expectedCreateSite  *gomock.Call
		expectedCreateOwner *gomock.Call
		want                *sitesv1alpha3.Site
		wantErr             bool
	}{
		{
			name: "Test create site",
			args: args{
				context.Background(),
				&sitesv1alpha3.CreateSiteRequest{
					Site: testsite,
				},
			},
			expectedCreateSite:  mockDatastore.EXPECT().Create(context.Background(), testsite).Return(nil),
			expectedCreateOwner: mockUserClient.EXPECT().CreateUser(gomock.Any(), &usersv1alpha3.CreateUserRequest{Parent: strings.Join([]string{"sites", testsite.Domain}, "/"), User: testsite.Owner}).Return(testsite.Owner, nil),
			want:                testsite,
		},
		{
			name: "Test create site with null request",
			args: args{
				context.Background(),
				&sitesv1alpha3.CreateSiteRequest{},
			},
			wantErr: true,
		},
		{
			name: "Test create site with empty title",
			args: args{
				context.Background(),
				&sitesv1alpha3.CreateSiteRequest{
					Site: &sitesv1alpha3.Site{
						Domain: "mysites.site",
						Owner: &usersv1alpha3.User{
							Username: "testowner",
							Email:    "testowner@mysites.site",
							Password: "secret",
							Fullname: "Test Owner",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create site with empty domain",
			args: args{
				context.Background(),
				&sitesv1alpha3.CreateSiteRequest{
					Site: &sitesv1alpha3.Site{
						Title: "My Sites",
						Owner: &usersv1alpha3.User{
							Username: "testowner",
							Email:    "testowner@mysites.site",
							Password: "secret",
							Fullname: "Test Owner",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Test create site with empty owner",
			args: args{
				context.Background(),
				&sitesv1alpha3.CreateSiteRequest{
					Site: &sitesv1alpha3.Site{
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
			server := newServiceServer(mockDatastore, mockUserClient)
			got, err := server.CreateSite(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetSite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := mock_sites.NewMockSiteDatastore(ctrl)
	mockUserClient := mock_users.NewMockUserServiceClient(ctrl)

	testsite := &sitesv1alpha3.Site{
		Name:   strings.Join([]string{"sites", "mysites.site"}, "/"),
		Title:  "My Sites",
		Domain: "mysites.site",
	}

	type args struct {
		ctx context.Context
		req *sitesv1alpha3.GetSiteRequest
	}
	tests := []struct {
		name            string
		args            args
		expectedGetSite *gomock.Call
		want            *sitesv1alpha3.Site
		wantErr         bool
	}{
		{
			name: "Test get site",
			args: args{
				context.Background(),
				&sitesv1alpha3.GetSiteRequest{
					Name: testsite.Name,
				},
			},
			expectedGetSite: mockDatastore.EXPECT().Get(context.Background(), testsite.Domain).Return(testsite, nil),
			want:            testsite,
		},
		{
			name: "Test get site not found",
			args: args{
				context.Background(),
				&sitesv1alpha3.GetSiteRequest{
					Name: strings.Join([]string{"sites", "notfound"}, "/"),
				},
			},
			expectedGetSite: mockDatastore.EXPECT().Get(context.Background(), "notfound").Return(nil, status.Error(codes.NotFound, "site not found")),
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, mockUserClient)
			got, err := server.GetSite(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteSite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDatastore := mock_sites.NewMockSiteDatastore(ctrl)
	mockUserClient := mock_users.NewMockUserServiceClient(ctrl)

	testsite := &sitesv1alpha3.Site{
		Name:   strings.Join([]string{"sites", "mysites.site"}, "/"),
		Title:  "My Sites",
		Domain: "mysites.site",
	}

	type args struct {
		ctx context.Context
		req *sitesv1alpha3.DeleteSiteRequest
	}
	tests := []struct {
		name               string
		args               args
		expectedDeleteSite *gomock.Call
		want               *empty.Empty
		wantErr            bool
	}{
		{
			name: "Test delete site",
			args: args{
				context.Background(),
				&sitesv1alpha3.DeleteSiteRequest{
					Name: testsite.Name,
				},
			},
			expectedDeleteSite: mockDatastore.EXPECT().Delete(context.Background(), testsite.Domain).Return(nil),
			want:               &empty.Empty{},
		},
		{
			name: "Test delete site not found",
			args: args{
				context.Background(),
				&sitesv1alpha3.DeleteSiteRequest{
					Name: strings.Join([]string{"sites", "notfound"}, "/"),
				},
			},
			expectedDeleteSite: mockDatastore.EXPECT().Delete(context.Background(), "notfound").Return(status.Error(codes.NotFound, "site not found")),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := newServiceServer(mockDatastore, mockUserClient)
			got, err := server.DeleteSite(tt.args.ctx, tt.args.req)
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
