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

package log

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// toLogrusFields convert Fields to logrus Fields
func toLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}

// Level aliasing logurs.Level
type Level logrus.Level

// GRPCCode aliasing grpc/codes.Code
type GRPCCode codes.Code

// Log level
const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// ToLevel convert grpc code to level
func (c GRPCCode) ToLevel() Level {
	code := codes.Code(c)
	switch code {
	case codes.OK:
		return InfoLevel
	case codes.Canceled:
		return InfoLevel
	case codes.Unknown:
		return ErrorLevel
	case codes.InvalidArgument:
		return InfoLevel
	case codes.DeadlineExceeded:
		return WarnLevel
	case codes.NotFound:
		return InfoLevel
	case codes.AlreadyExists:
		return InfoLevel
	case codes.PermissionDenied:
		return WarnLevel
	case codes.Unauthenticated:
		return InfoLevel
	case codes.ResourceExhausted:
		return WarnLevel
	case codes.FailedPrecondition:
		return WarnLevel
	case codes.Aborted:
		return WarnLevel
	case codes.OutOfRange:
		return WarnLevel
	case codes.Unimplemented:
		return ErrorLevel
	case codes.Internal:
		return ErrorLevel
	case codes.Unavailable:
		return WarnLevel
	case codes.DataLoss:
		return ErrorLevel
	default:
		return ErrorLevel
	}
}
