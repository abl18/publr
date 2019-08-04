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
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/prksu/publr/pkg/log"
)

// ServerLoggingInterceptor is unary server interceptor for logging
func ServerLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	grpcstatus, _ := status.FromError(err)

	logMsg := "call grpc request"
	logFields := log.WithFields(
		log.Fields{
			"grpc.service":        ServiceName,
			"grpc.method":         info.FullMethod,
			"grpc.time.start":     start.Format(time.RFC3339),
			"grpc.time.duration":  time.Since(start),
			"grpc.status.code":    grpcstatus.Code(),
			"grpc.status.message": grpcstatus.Message(),
			"grpc.status.detail":  grpcstatus.Details(),
		},
	)

	switch log.GRPCCode(grpcstatus.Code()).ToLevel() {
	case log.InfoLevel:
		logFields.Info(logMsg)
	case log.WarnLevel:
		logFields.Warn(logMsg)
	case log.ErrorLevel:
		logFields.Error(logMsg)
	case log.FatalLevel:
		logFields.Fatal(logMsg)
	}

	return h, err
}
