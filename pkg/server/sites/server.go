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

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sitesv1alpha1 "github.com/prksu/publr/pkg/api/sites/v1alpha1"
)

// Sites service
var (
	ServiceName    = "sites"
	ServiceAddress = "0.0.0.0:9000"
	ServiceVersion = "v1alpha1"
)

// Server implement sitesv1alpha1.SiteServiceServer.
type Server struct {
	DS SiteDatastore
}

// NewServer create new sites service server.
// returns sitesv1alpha1.SiteServiceServer.
func NewServer() sitesv1alpha1.SiteServiceServer {
	server := new(Server)
	server.DS = NewDatastore()
	return server
}

// CreateSite handler method.
func (s *Server) CreateSite(ctx context.Context, req *sitesv1alpha1.CreateSiteRequest) (*sitesv1alpha1.Site, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// GetSite handler method.
func (s *Server) GetSite(ctx context.Context, req *sitesv1alpha1.GetSiteRequest) (*sitesv1alpha1.Site, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}

// DeleteSite handler method.
func (s *Server) DeleteSite(ctx context.Context, req *sitesv1alpha1.DeleteSiteRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "not implement yet")
}
