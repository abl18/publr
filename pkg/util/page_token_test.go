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
	"testing"
)

func TestOptions_Generate(t *testing.T) {
	pageToken := NewPageToken()

	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test generate page token",
			args: args{index: 1},
			want: base64.StdEncoding.EncodeToString([]byte(defaultPageTokenSalt + "1")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pageToken.Generate(tt.args.index); got != tt.want {
				t.Errorf("Options.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_Parse(t *testing.T) {
	pageToken := NewPageToken()

	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test parse token",
			args:    args{token: base64.StdEncoding.EncodeToString([]byte(defaultPageTokenSalt + "1"))},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test parse empty token",
			args: args{token: ""},
			want: 0,
		},
		{
			name:    "Test invalid token",
			args:    args{token: "invalid"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Test invalid token without salt",
			args:    args{token: base64.StdEncoding.EncodeToString([]byte("invalid"))},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Test invalid token not parseable",
			args:    args{token: base64.StdEncoding.EncodeToString([]byte(defaultPageTokenSalt + "invalid"))},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pageToken.Parse(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Options.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Options.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
