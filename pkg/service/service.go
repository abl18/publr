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

package service

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/prksu/publr/pkg/log"
	"github.com/prksu/publr/pkg/service/logging"
)

// Server global var
var (
	ServerAddress string
	ServerTLS     bool
	ServerCert    string
	ServerKey     string
	CA            string
)

// Service struct
type Service struct {
	grpc *grpc.Server
	http *http.Server
	tls  *tls.Config
}

func init() {
	flag.StringVar(&ServerAddress, "server-address", "0.0.0.0:9443", "Server address")
	flag.BoolVar(&ServerTLS, "server-tls", false, "Enable server TLS")
	flag.StringVar(&ServerCert, "server-cert", "", "Server certifiate")
	flag.StringVar(&ServerKey, "server-key", "", "Server key")
	flag.StringVar(&CA, "ca", "", "Certificate authority")
}

// NewService create new service instance
func NewService() (*Service, error) {
	service := new(Service)

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logging.ServerInterceptor),
	}

	if ServerTLS {
		ServerKeyPair, err := tls.LoadX509KeyPair(ServerCert, ServerKey)
		if err != nil {
			return nil, err
		}

		ca, err := ioutil.ReadFile(CA)
		if err != nil {
			return nil, err
		}

		CertPool := x509.NewCertPool()
		CertPool.AppendCertsFromPEM(ca)

		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&ServerKeyPair)))

		service.tls = &tls.Config{
			Certificates: []tls.Certificate{ServerKeyPair},
			NextProtos:   []string{"h2"},
			RootCAs:      CertPool,
		}

	}

	service.grpc = grpc.NewServer(opts...)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	service.http = &http.Server{
		Handler:   handlerFunc(service.grpc, mux),
		TLSConfig: service.tls,
	}

	return service, nil
}

func handlerFunc(g *grpc.Server, h http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			g.ServeHTTP(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// GRPC returns grpc.Server
func (s *Service) GRPC() *grpc.Server { return s.grpc }

// ListenAndServe ...
func (s *Service) ListenAndServe() error {
	listener, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		return err
	}

	if s.tls == nil {
		log.Warnf("serve insecure service on %s", listener.Addr())
		return s.http.Serve(listener)

	}

	log.Infof("serve service on %s", listener.Addr())
	return s.http.Serve(tls.NewListener(listener, s.tls))
}
