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
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	sitesv1alpha3 "github.com/prksu/publr/pkg/api/sites/v1alpha3"
)

// DefaultAddress default sites service server address
var DefaultAddress = "dns:///sites:9000"

// MustNewServiceClient create new sites service client with panic if any errors.
func MustNewServiceClient() sitesv1alpha3.SiteServiceClient {
	client, err := NewServiceClient(DefaultAddress)
	if err != nil {
		panic(err)
	}
	return client
}

// NewServiceClient create new sites service client.
func NewServiceClient(address string) (sitesv1alpha3.SiteServiceClient, error) {
	// ca, err := ioutil.ReadFile(service.CA)
	// if err != nil {
	// 	return nil, err
	// }

	// CertPool := x509.NewCertPool()
	// CertPool.AppendCertsFromPEM(ca)

	opts := []grpc.DialOption{
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
		grpc.WithInsecure(),
		// grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		// 	RootCAs: CertPool,
		// })),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	client := sitesv1alpha3.NewSiteServiceClient(conn)

	return client, nil
}
