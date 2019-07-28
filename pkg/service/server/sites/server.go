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
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sitesv1alpha2 "github.com/prksu/publr/pkg/api/sites/v1alpha2"
	usersv1alpha2 "github.com/prksu/publr/pkg/api/users/v1alpha2"
	"github.com/prksu/publr/pkg/service/client/users"
	"github.com/prksu/publr/pkg/service/server/sites/datastore"
)

// Server implement sitesv1alpha2.SiteServiceServer.
type Server struct {
	Site       datastore.SiteDatastore
	UserClient usersv1alpha2.UserServiceClient
}

// NewServiceServer create new sites service server.
// returns sitesv1alpha2.SiteServiceServer.
func NewServiceServer() sitesv1alpha2.SiteServiceServer {
	return newServiceServer(datastore.NewSiteDatastore(), users.MustNewServiceClient())
}

func newServiceServer(site datastore.SiteDatastore, userClient usersv1alpha2.UserServiceClient) sitesv1alpha2.SiteServiceServer {
	server := new(Server)
	server.Site = site
	server.UserClient = userClient
	return server
}

// CreateSite handler method.
func (s *Server) CreateSite(ctx context.Context, req *sitesv1alpha2.CreateSiteRequest) (*sitesv1alpha2.Site, error) {
	site := req.Site

	if site == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if site.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}

	if site.Domain == "" {
		return nil, status.Error(codes.InvalidArgument, "domain is required")
	}

	owner := site.Owner
	if owner == nil {
		return nil, status.Error(codes.InvalidArgument, "owner is required")
	}

	sitedomain := site.Domain
	if err := s.Site.Create(ctx, site); err != nil {
		return nil, err
	}

	owner.Role = 3
	ownerres, err := s.UserClient.CreateUser(ctx, &usersv1alpha2.CreateUserRequest{Parent: strings.Join([]string{"sites", sitedomain}, "/"), User: owner})
	if err != nil {
		return nil, s.Site.Delete(ctx, sitedomain)
	}

	res := new(sitesv1alpha2.Site)
	res = req.Site
	res.Name = strings.Join([]string{"sites", sitedomain}, "/")
	res.Owner = ownerres
	return res, nil
}

// GetSite handler method.
func (s *Server) GetSite(ctx context.Context, req *sitesv1alpha2.GetSiteRequest) (*sitesv1alpha2.Site, error) {
	name := req.Name
	sitedomain := strings.Split(name, "/")[1]
	res, err := s.Site.Get(ctx, sitedomain)
	if err != nil {
		return nil, err
	}

	res.Name = name
	return res, nil
}

// DeleteSite handler method.
func (s *Server) DeleteSite(ctx context.Context, req *sitesv1alpha2.DeleteSiteRequest) (*empty.Empty, error) {
	name := req.Name
	sitedomain := strings.Split(name, "/")[1]

	if err := s.Site.Delete(ctx, sitedomain); err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}
