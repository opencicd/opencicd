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

package provider

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

// Utility function that allows waiting for a provider to run.
// Mostly usable from other providers that have a dependency.
func WaitForRunningProvider(p RunProvider, timeoutSeconds time.Duration) error {
	if p.IsRunning() {
		// No need to wait if provider is already running.
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	name := Name(p)
	logrus.WithField("timeout", timeoutSeconds).Debugf("Waiting for %s to run...", name)
	for {
		if p.IsRunning() {
			logrus.Debugf("%s is running", name)
			return nil
		}
		if ctx.Err() != nil {
			return fmt.Errorf("time exceeded for %s to run", name)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// Utility function to get a provider's name.
func Name(provider Provider) string {
	return reflect.TypeOf(provider).Elem().String()
}
