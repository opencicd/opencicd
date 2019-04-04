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

// Resource ...
type Resource interface {
	CreateUser(context.Context, *identitypb.CreateUserRequest) (*identitypb.CreateUserResponse, error)
	ListUsers(context.Context, *identitypb.ListUsersRequest) (*identitypb.ListUsersResponse, error)
	GetUser(context.Context, *identitypb.GetUserRequest) (*identitypb.GetUserResponse, error)
	UpdateUser(context.Context, *identitypb.UpdateUserRequest) (*identitypb.UpdateUserResponse, error)
	DeleteUser(context.Context, *identitypb.DeleteUserRequest) (*identitypb.DeleteUserResponse, error)
}
