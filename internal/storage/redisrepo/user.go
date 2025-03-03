package redisrepo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/application/userapp"
)

var _ userapp.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	client *redis.Client
}

type RedisUser struct {
	ID   string `redis:"id"`
	Name string `redis:"name"`
}

const userKeyFormat = "user:%s"

func (user RedisUser) userKey() string {
	return fmt.Sprintf(userKeyFormat, user.ID)
}

func (repo *UserRepository) Save(ctx context.Context, user userapp.UserDTO) error {
	redisUser := RedisUser{
		ID:   user.ID,
		Name: user.Name,
	}

	if err := repo.client.HMSet(ctx, redisUser.userKey(), redisUser).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Remove(ctx context.Context, id string) error {
	redisUser := RedisUser{ID: id}
	userKey := redisUser.userKey()

	exists, err := repo.client.Exists(ctx, userKey).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return ErrUserNotFound
	}
	return repo.client.HDel(ctx, userKey).Err()
}

func (repo *UserRepository) GetByID(ctx context.Context, id string) (userapp.UserDTO, error) {
	redisUser := RedisUser{ID: id}
	userKey := redisUser.userKey()

	exists, err := repo.client.Exists(ctx, userKey).Result()
	if err != nil {
		return userapp.UserDTO{}, err
	}
	if exists == 0 {
		return userapp.UserDTO{}, ErrUserNotFound
	}

	if err := repo.client.HGetAll(ctx, userKey).Scan(&redisUser); err != nil {
		return userapp.UserDTO{}, err
	}

	return userapp.UserDTO{
		ID:   redisUser.ID,
		Name: redisUser.Name,
	}, nil
}
