package data_accessors

import (
	"context"
	"time"

	ett "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/entities"
	ida "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/interfaces/data_acessors"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDAMongo struct {
	MgCl       *mongo.Client
	ctxTimeout time.Duration
}

// Add implements data_accessors.UserDataAccessor
func (*userDAMongo) Add(ctx context.Context, model ett.User) (ett.User, error) {
	panic("unimplemented")
}

// Delete implements data_accessors.UserDataAccessor
func (*userDAMongo) Delete(ctx context.Context, userId string, softDelete bool) (bool, error) {
	panic("unimplemented")
}

// Get implements data_accessors.UserDataAccessor
func (*userDAMongo) Get(ctx context.Context, userId string) (ett.User, error) {
	panic("unimplemented")
}

// Search implements data_accessors.UserDataAccessor
func (*userDAMongo) Search(ctx context.Context, searchTerm string, pageNumber int, pageAmount int) ([]ett.User, error) {
	panic("unimplemented")
}

// Update implements data_accessors.UserDataAccessor
func (*userDAMongo) Update(ctx context.Context, model ett.User) (ett.User, error) {
	panic("unimplemented")
}

func NewUserDAMongo(mgcl *mongo.Client, timeout time.Duration) ida.UserDataAccessor {
	return &userDAMongo{
		MgCl:       mgcl,
		ctxTimeout: timeout,
	}
}
