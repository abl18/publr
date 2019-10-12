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

package util

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var defaultPageTokenSalt = strconv.FormatInt(time.Now().Unix(), 10)

// PageToken a page token for list result
type PageToken interface {
	Generate(int) string
	Parse(string) (int, error)
}

// Options implement page token
type Options struct {
	salt string
}

// NewPageToken returns PageToken
func NewPageToken() PageToken {
	token := new(Options)
	token.salt = defaultPageTokenSalt
	return token
}

// Generate page token for next_page_token
func (o *Options) Generate(index int) string {
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s%d", o.salt, index)))
	return token
}

// Parse page token
func (o *Options) Parse(token string) (int, error) {
	if token == "" {
		return 0, nil
	}

	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return -1, err
	}

	if !strings.HasPrefix(string(b), o.salt) {
		return -1, status.Errorf(codes.InvalidArgument, "invalid page_token")
	}

	index, err := strconv.Atoi(strings.TrimPrefix(string(b), o.salt))
	if err != nil {
		return -1, status.Errorf(codes.InvalidArgument, "invalid page_token")
	}

	return index, nil
}
