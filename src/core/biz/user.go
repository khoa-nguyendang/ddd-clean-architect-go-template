package biz

import (
	"context"
	"time"

	ett "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/entities"
	ibs "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/interfaces/biz_services"
	ida "github.com/khoa-nguyendang/ddd-clean-architect-go-template/core/interfaces/data_acessors"
)

type userBiz struct {
	userDA     ida.UserDataAccessor
	ctxTimeout time.Duration
}

// Add implements biz_services.UserBiz
func (*userBiz) Add(ctx context.Context, model ett.User) (ett.User, error) {
	panic("unimplemented")
}

// Delete implements biz_services.UserBiz
func (*userBiz) Delete(ctx context.Context, userId string, softDelete bool) (bool, error) {
	panic("unimplemented")
}

// Get implements biz_services.UserBiz
func (*userBiz) Get(ctx context.Context, userId string) (ett.User, error) {
	panic("unimplemented")
}

// Search implements biz_services.UserBiz
func (*userBiz) Search(ctx context.Context, searchTerm string, pageNumber int, pageAmount int) ([]ett.User, error) {
	panic("unimplemented")
}

// Update implements biz_services.UserBiz
func (*userBiz) Update(ctx context.Context, model ett.User) (ett.User, error) {
	panic("unimplemented")
}

func NewUserBiz(u ida.UserDataAccessor, timeout time.Duration) ibs.UserBiz {
	return &userBiz{
		userDA:     u,
		ctxTimeout: timeout,
	}
}
