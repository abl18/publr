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

package users

import (
	"net"

	"google.golang.org/grpc"

	usersv1alpha1 "github.com/prksu/publr/pkg/api/users/v1alpha1"
	"github.com/prksu/publr/pkg/service/server/users"
)

// NewServiceClient create new users service client.
func NewServiceClient() (usersv1alpha1.UserServiceClient, error) {
	host, port, err := net.SplitHostPort(users.ServiceAddress)
	if err != nil {
		return nil, err
	}

	host = users.ServiceName
	address := net.JoinHostPort(host, port)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	client := usersv1alpha1.NewUserServiceClient(conn)

	return client, nil
}
