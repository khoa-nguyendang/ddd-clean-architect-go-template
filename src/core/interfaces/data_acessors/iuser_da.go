package data_accessors

import (
	"context"

	ett "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/entities"
)

type UserDataAccessor interface {
	Search(ctx context.Context, searchTerm string, pageNumber, pageAmount int) ([]ett.User, error)
	Get(ctx context.Context, userId string) (ett.User, error)
	Add(ctx context.Context, model ett.User) (ett.User, error)
	Update(ctx context.Context, model ett.User) (ett.User, error)
	Delete(ctx context.Context, userId string, softDelete bool) (bool, error)
}
