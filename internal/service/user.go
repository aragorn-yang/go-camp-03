package service

import (
	"context"
	"strconv"

	api "github.com/aragorn-yang/go-camp-03/api/user/v1"
	"github.com/aragorn-yang/go-camp-03/internal/domain"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

type UserService struct {
	api.UserServiceServer

	usecase domain.IUserUsecase
}

func NewUserService(usecase domain.IUserUsecase) *UserService {
	return &UserService{usecase: usecase}
}

func (u *UserService) GetUserInfo(ctx context.Context, req *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("get metadata err")
	}

	data := md.Get("uid")
	if len(data) != 1 {
		return nil, errors.Errorf("user id lens not 1, metadata: %v", data)
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.Errorf("user id not a num, data: %v", data)
	}

	user, err := u.usecase.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &api.GetUserInfoResponse{
		Username: user.Name,
		City:     user.City,
	}

	return resp, nil
}
