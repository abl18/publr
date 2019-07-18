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
)

var logger = newLogger()

// Info logs an Entry with a status of "info"
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logs an Entry with a status of "info"
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn logs an Entry with a status of "warn"
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf logs an Entry with a status of "warn"
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error logs an Entry with a status of "error"
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf logs an Entry with a status of "error"
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal logs an Entry with a status of "fatal" and exits the program
// with status code 1.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf logs an Entry with a status of "fatal" and exits the program
// with status code 1.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// WithFields ...
func WithFields(fields Fields) *logrus.Entry {
	return logger.WithFields(toLogrusFields(fields))
}
