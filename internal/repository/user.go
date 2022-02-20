package repository

import (
	"context"
	usererror "github.com/aragorn-yang/go-camp-03/errors"
	"github.com/aragorn-yang/go-camp-03/internal/domain"
	"github.com/aragorn-yang/go-camp-03/internal/repository/ent"
	"github.com/pkg/errors"
)

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) domain.IUserRepo {
	return &repository{client: client}
}

func (r *repository) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	user, err := r.client.User.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, errors.Wrapf(usererror.UserNotFound, "user not found, id: %d, err: %+v", id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(usererror.Unknown, "db query err: %+v, id: %d,", err, id)
	}

	return &domain.User{
		Name: user.Name,
		City: user.City,
		ID:   user.ID,
	}, nil
}
