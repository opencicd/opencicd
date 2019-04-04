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

package main

import (
	"github.com/opencicd/opencicd/internal/identity-svc/user"
	"github.com/opencicd/opencicd/pkg/v1/stack"
)

func main() {
	s := stack.New()
	defer s.MustClose()

	resource := user.NewResourceMongo()
	s.MustInit(resource)

	controller := user.NewController(resource)
	s.MustInit(controller)

	service := user.NewService(controller)
	s.MustInit(service)

	s.MustRun()
}
