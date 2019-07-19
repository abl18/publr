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
	"flag"
	"fmt"
	"os"

	usersv1alpha2 "github.com/prksu/publr/pkg/api/users/v1alpha2"
	"github.com/prksu/publr/pkg/service"
	"github.com/prksu/publr/pkg/service/server/users"
	"github.com/prksu/publr/pkg/storage/database"
)

func init() {
	flag.StringVar(&database.Host, "database-host", "127.0.0.1", "Database host")
	flag.StringVar(&database.User, "database-user", "root", "Database user")
	flag.StringVar(&database.Password, "database-password", "", "Database password")
	flag.StringVar(&database.Name, "database-name", "publr", "Database name")
}

func run() error {
	s, err := service.NewService()
	if err != nil {
		return err
	}

	usersv1alpha2.RegisterUserServiceServer(s.GRPC(), users.NewServiceServer())
	return s.ListenAndServe()
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
