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

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials"

	postsv1alpha2 "github.com/prksu/publr/pkg/api/posts/v1alpha2"
	sitesv1alpha2 "github.com/prksu/publr/pkg/api/sites/v1alpha2"
	usersv1alpha2 "github.com/prksu/publr/pkg/api/users/v1alpha2"
	"github.com/prksu/publr/pkg/service"
)

var (
	// InecureServer insecure grpc server
	InecureServer bool
)

func init() {
	flag.BoolVar(&InecureServer, "insecure-server", false, "Insecure grpc server")
}

func run() error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithBalancerName(roundrobin.Name),
	}

	if InecureServer {
		opts = append(opts, grpc.WithInsecure())
	} else {
		ca, err := ioutil.ReadFile(service.CA)
		if err != nil {
			return err
		}

		CertPool := x509.NewCertPool()
		CertPool.AppendCertsFromPEM(ca)
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			RootCAs: CertPool,
		})))
	}

	postsv1alpha2.RegisterPostServiceHandlerFromEndpoint(context.Background(), mux, "dns:///posts:9000", opts)
	sitesv1alpha2.RegisterSiteServiceHandlerFromEndpoint(context.Background(), mux, "dns:///sites:9000", opts)
	usersv1alpha2.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "dns:///users:9000", opts)

	return http.ListenAndServe(":8000", mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
