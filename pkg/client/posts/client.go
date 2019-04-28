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

package posts

import (
	"net"

	"google.golang.org/grpc"

	postsv1alpha1 "github.com/prksu/publr/pkg/api/posts/v1alpha1"
	"github.com/prksu/publr/pkg/server/posts"
)

// NewClient create new posts service client.
func NewClient() (postsv1alpha1.PostServiceClient, error) {
	host, port, err := net.SplitHostPort(posts.ServiceAddress)
	if err != nil {
		return nil, err
	}

	host = posts.ServiceName
	address := net.JoinHostPort(host, port)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	client := postsv1alpha1.NewPostServiceClient(conn)

	return client, nil
}
