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

package user

import (
	"context"

	identitypb "github.com/opencicd/opencicd/api/proto/v1/identity"
)

// Service ...
type Service struct {
	controller *Controller
}

// NewService ...
func NewService(controller *Controller) *Service {
	return &Service{
		controller: controller,
	}
}

// Init ...
func (s *Service) Init() error {
	return nil
}

// Close ...
func (s *Service) Close() error {
	return nil
}

// CreateUser ...
func (s *Service) CreateUser(ctx context.Context, request *identitypb.CreateUserRequest) (*identitypb.CreateUserResponse, error) {
	return s.controller.CreateUser(ctx, request)
}

// ListUsers ...
func (s *Service) ListUsers(ctx context.Context, request *identitypb.ListUsersRequest) (*identitypb.ListUsersResponse, error) {
	return s.controller.ListUsers(ctx, request)
}

// GetUser ...
func (s *Service) GetUser(ctx context.Context, request *identitypb.GetUserRequest) (*identitypb.GetUserResponse, error) {
	return s.controller.GetUser(ctx, request)
}

// UpdateUser ...
func (s *Service) UpdateUser(ctx context.Context, request *identitypb.UpdateUserRequest) (*identitypb.UpdateUserResponse, error) {
	return s.controller.UpdateUser(ctx, request)
}

// DeleteUser ...
func (s *Service) DeleteUser(ctx context.Context, request *identitypb.DeleteUserRequest) (*identitypb.DeleteUserResponse, error) {
	return s.controller.DeleteUser(ctx, request)
}
