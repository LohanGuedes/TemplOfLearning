package userservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lohanguedes/templOfLearning/internal/database"
)

type UserRepo interface {
	CreateUser(context.Context, database.CreateUserParams) (database.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (database.User, error)
}

type UserService struct {
	UserRepo UserRepo
}

func New(ur UserRepo) UserService {
	return UserService{
		UserRepo: ur,
	}
}

func (us UserService) Create(ctx context.Context, b io.Reader) (database.User, error) {
	var data database.CreateUserParams
	err := json.NewDecoder(b).Decode(&data)
	if err != nil {
		return database.User{}, fmt.Errorf("invalid request-body json format: %w", err)
	}

	u, err := us.UserRepo.CreateUser(ctx, data)
	if err != nil {
		return database.User{}, err
	}
	return u, nil
}

func (us UserService) GetUserByID(ctx context.Context, id string) (database.User, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(id)
	if err != nil {
		return database.User{}, fmt.Errorf("invalid id: %s", id)
	}
	u, err := us.UserRepo.GetUserByID(ctx, uuid)
	if err != nil {
		return database.User{}, err
	}

	return u, nil
}
