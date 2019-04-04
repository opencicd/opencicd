// Copyright 2019 OpenCICD Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logrus

import (
	"io"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	defaultLevel     = "info"
	defaultFormatter = "json"
	defaultOutput    = "stderr"
)

// Config ...
type Config struct {
	Level     logrus.Level     // Logging level (trace, debug, info, etc).
	Formatter logrus.Formatter // Log formatter (text or json).
	Output    io.Writer        // Output writer (stderr or stdout).
}

// Initializes the configuration from environment variables.
// Basically just logs the output of ParseEnv().
func NewConfigFromEnv() *Config {
	level, formatter, output := ParseEnv()

	// Temporary logger.
	logger := NewLogger(level, formatter, output)
	logger.WithFields(logrus.Fields{
		"level":     level,
		"formatter": reflect.TypeOf(formatter).Elem().String(),
		"output":    reflect.TypeOf(output).Elem().String(),
	}).Debug("Logrus Config initialized")

	return &Config{
		Level:     level,
		Formatter: formatter,
		Output:    output,
	}
}

// Parses the environment variables.
// This is done in a separate method, which can be used as parameter for NewLogger().
func ParseEnv() (logrus.Level, logrus.Formatter, io.Writer) {
	v := viper.New()
	v.SetEnvPrefix("LOGRUS")
	v.AutomaticEnv()

	v.SetDefault("LEVEL", defaultLevel)
	level, err := logrus.ParseLevel(v.GetString("LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}

	v.SetDefault("FORMATTER", defaultFormatter)
	var formatter logrus.Formatter
	switch v.GetString("FORMATTER") {
	case "text":
		formatter = &logrus.TextFormatter{}
	case "text_clr":
		formatter = &logrus.TextFormatter{ForceColors: true}
	case "json":
		fallthrough
	default:
		formatter = &logrus.JSONFormatter{}
	}

	v.SetDefault("OUTPUT", defaultOutput)
	var output io.Writer
	switch v.GetString("OUTPUT") {
	case "stdout":
		output = os.Stdout
	case "stderr":
		fallthrough
	default:
		output = os.Stderr
	}

	return level, formatter, output
}
