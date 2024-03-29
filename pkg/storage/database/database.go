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

package database

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	"contrib.go.opencensus.io/integrations/ocsql"
	_ "github.com/go-sql-driver/mysql" // mysql driver

	"github.com/prksu/publr/pkg/log"
)

// Database global var
var (
	Host     string
	User     string
	Password string
	Name     string
)

func init() {
	flag.StringVar(&Host, "database-host", "127.0.0.1", "Database host")
	flag.StringVar(&User, "database-user", "root", "Database user")
	flag.StringVar(&Password, "database-password", "", "Database password")
	flag.StringVar(&Name, "database-name", "publr", "Database name")
}

// Database interface
type Database interface {
	WithDriver(driver string) Database
	WithDSN(dsn string) Database
	Connect() *sql.DB
}

// Options implement database
type Options struct {
	DSN    string
	Driver string
}

// NewDatabase create new database configuration
func NewDatabase() Database {
	database := new(Options)
	database.DSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?autocommit=true&parseTime=true", User, Password, Host, Name)
	return database
}

// WithDriver set database driver.
func (o *Options) WithDriver(driver string) Database {
	o.Driver = driver
	return o
}

// WithDSN set database driver, This method will be overide the DSN from database configuration variable.
func (o *Options) WithDSN(dsn string) Database {
	o.DSN = dsn
	return o
}

// Connect open database connection.
func (o *Options) Connect() *sql.DB {
	// Open sql connection
	if o.Driver == "" {
		o.Driver = "mysql"
	}

	driver, err := ocsql.Register(o.Driver, ocsql.WithAllTraceOptions(), ocsql.WithPing(false), ocsql.WithDisableErrSkip(true))
	if err != nil {
		log.Fatal(err)
	}

	database, err := sql.Open(driver, o.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	database.SetMaxOpenConns(5)
	database.SetMaxIdleConns(5)
	database.SetConnMaxLifetime(time.Hour)
	return database
}
