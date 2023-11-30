package server

import (
	"context"
	"sync"

	buildsv1 "github.com/chaitanyapantheor/golang-grpc-http/proto/builds/v1"
	usersv1 "github.com/chaitanyapantheor/golang-grpc-http/proto/users/v1"
	"github.com/gofrs/uuid"
)

// Backend implements the protobuf interface
type Backend struct {
	mu     *sync.RWMutex
	users  []*usersv1.User
	builds []*buildsv1.Build
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

// AddUser adds a user to the in-memory store.
func (b *Backend) AddUser(ctx context.Context, _ *usersv1.AddUserRequest) (*usersv1.AddUserResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	user := &usersv1.User{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.users = append(b.users, user)

	return &usersv1.AddUserResponse{
		User: user,
	}, nil
}

// ListUsers lists all users in the store.
func (b *Backend) ListUsers(_ *usersv1.ListUsersRequest, srv usersv1.UserService_ListUsersServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, user := range b.users {
		err := srv.Send(&usersv1.ListUsersResponse{User: user})
		if err != nil {
			return err
		}
	}

	return nil
}

// AddUser adds a build to the in-memory store.
func (b *Backend) AddBuild(ctx context.Context, breq *buildsv1.AddBuildRequest) (*buildsv1.AddBuildResponse, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	build := &buildsv1.Build{
		Id:    uuid.Must(uuid.NewV4()).String(),
		Label: breq.GetLabel(),
	}
	b.builds = append(b.builds, build)

	return &buildsv1.AddBuildResponse{
		Build: build,
	}, nil
}

// ListUsers lists all builds in the store.
func (b *Backend) ListBuilds(_ *buildsv1.ListBuildsRequest, srv buildsv1.BuildService_ListBuildsServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, build := range b.builds {
		err := srv.Send(&buildsv1.ListBuildsResponse{Build: build})
		if err != nil {
			return err
		}
	}

	return nil
}
