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

// ResourceMongo ...
type ResourceMongo struct {
}

// NewResourceMongo ...
func NewResourceMongo() *ResourceMongo {
	return &ResourceMongo{}
}

// Init ...
func (r *ResourceMongo) Init() error {
	return nil
}

// Close ...
func (r *ResourceMongo) Close() error {
	return nil
}

// CreateUser ...
func (r *ResourceMongo) CreateUser(ctx context.Context, request *identitypb.CreateUserRequest) (*identitypb.CreateUserResponse, error) {
	return nil, nil
}

// ListUsers ...
func (r *ResourceMongo) ListUsers(ctx context.Context, request *identitypb.ListUsersRequest) (*identitypb.ListUsersResponse, error) {
	return nil, nil
}

// GetUser ...
func (r *ResourceMongo) GetUser(ctx context.Context, request *identitypb.GetUserRequest) (*identitypb.GetUserResponse, error) {
	return nil, nil
}

// UpdateUser ...
func (r *ResourceMongo) UpdateUser(ctx context.Context, request *identitypb.UpdateUserRequest) (*identitypb.UpdateUserResponse, error) {
	return nil, nil
}

// DeleteUser ...
func (r *ResourceMongo) DeleteUser(ctx context.Context, request *identitypb.DeleteUserRequest) (*identitypb.DeleteUserResponse, error) {
	return nil, nil
}
