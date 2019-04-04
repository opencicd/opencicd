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

// Controller ...
type Controller struct {
	resource Resource
}

// NewController ...
func NewController(resource Resource) *Controller {
	return &Controller{
		resource: resource,
	}
}

// Init ...
func (c *Controller) Init() error {
	return nil
}

// Close ...
func (c *Controller) Close() error {
	return nil
}

// CreateUser ...
func (c *Controller) CreateUser(ctx context.Context, request *identitypb.CreateUserRequest) (*identitypb.CreateUserResponse, error) {
	res, err := c.resource.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListUsers ...
func (c *Controller) ListUsers(ctx context.Context, request *identitypb.ListUsersRequest) (*identitypb.ListUsersResponse, error) {
	res, err := c.resource.ListUsers(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetUser ...
func (c *Controller) GetUser(ctx context.Context, request *identitypb.GetUserRequest) (*identitypb.GetUserResponse, error) {
	res, err := c.resource.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateUser ...
func (c *Controller) UpdateUser(ctx context.Context, request *identitypb.UpdateUserRequest) (*identitypb.UpdateUserResponse, error) {
	res, err := c.resource.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteUser ...
func (c *Controller) DeleteUser(ctx context.Context, request *identitypb.DeleteUserRequest) (*identitypb.DeleteUserResponse, error) {
	res, err := c.resource.DeleteUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
