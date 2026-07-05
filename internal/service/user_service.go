package service

import (
	"context"
	"fmt"
	"time"

	"github.com/f7rzen/user-cache-service/internal/cache"
	"github.com/f7rzen/user-cache-service/internal/client"
	"github.com/f7rzen/user-cache-service/internal/model"
)

type UserService struct {
	userClient *client.UserClient
	cache      *cache.Cache
	cacheTTL   time.Duration
}

func NewUserService(userClient *client.UserClient, cache *cache.Cache, cacheTTL time.Duration) *UserService {
	return &UserService{
		userClient: userClient,
		cache:      cache,
		cacheTTL:   cacheTTL,
	}
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.UserResponse, error) {
	cacheKey := "users:list"

	cachedValue, ok := s.cache.Get(cacheKey)
	if ok {
		users := cachedValue.([]model.UserResponse)
		return users, nil
	}

	externalUsers, err := s.userClient.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]model.UserResponse, 0, len(externalUsers))

	for _, externalUser := range externalUsers {
		user := model.UserResponse{
			ID:       externalUser.ID,
			FullName: externalUser.Name,
			Email:    externalUser.Email,
			City:     externalUser.Address.City,
			Company:  externalUser.Company.Name,
		}

		users = append(users, user)
	}

	s.cache.Set(cacheKey, users, s.cacheTTL)

	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (model.UserResponse, error) {
	cacheKey := fmt.Sprintf("users:id:%d", id)

	cachedValue, ok := s.cache.Get(cacheKey)
	if ok {
		user := cachedValue.(model.UserResponse)
		return user, nil
	}

	externalUser, err := s.userClient.GetUserByID(ctx, id)
	if err != nil {
		return model.UserResponse{}, err
	}

	user := model.UserResponse{
		ID:       externalUser.ID,
		FullName: externalUser.Name,
		Email:    externalUser.Email,
		City:     externalUser.Address.City,
		Company:  externalUser.Company.Name,
	}

	s.cache.Set(cacheKey, user, s.cacheTTL)

	return user, nil
}
