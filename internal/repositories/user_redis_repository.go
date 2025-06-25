package repositories

import (
	"context"
	"encoding/json"
	"go-echo-clean-architecture/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type UserRedisRepository interface {
	Create(ctx context.Context, user *models.User)
	FindByUsernameOrEmailOrPhone(query string, ctx context.Context) *models.User
	Update(existingUser *models.User, newUser *models.User, ctx context.Context)
	Delete(user *models.User, ctx context.Context)
}

type UserRedisRepositoryImpl struct {
	redisClient *redis.Client
}

func NewUserRedisRepositoryImpl(redisClient *redis.Client) *UserRedisRepositoryImpl {
	return &UserRedisRepositoryImpl{redisClient: redisClient}
}

func (u *UserRedisRepositoryImpl) Create(ctx context.Context, user *models.User) {
	logrus.Infof("Creating user cache %s", user.ID.String())

	userJson, err := json.Marshal(user)
	if err != nil {
		logrus.Errorf("Error marshalling user: %v", err)
	}

	err = u.redisClient.Set(ctx, "user:"+user.ID.String(), userJson, 10*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Error creating user index: %v", err)
	}
	err = u.redisClient.Set(ctx, "user:"+user.Username, userJson, 10*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Error creating user index: %v", err)
	}
	err = u.redisClient.Set(ctx, "user:"+user.Email, userJson, 10*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Error creating user index: %v", err)
	}
	err = u.redisClient.Set(ctx, "user:"+user.Phone, userJson, 10*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Error creating user index: %v", err)
	}
	return
}

func (u *UserRedisRepositoryImpl) FindByUsernameOrEmailOrPhone(query string, ctx context.Context) *models.User {
	logrus.Infof("Finding user cache %s", query)

	// find by index
	user, err := u.redisClient.Get(ctx, "user:"+query).Result()
	if err != nil {
		logrus.Errorf("Error finding user cached %v", err)
	}

	// serialize user (string) to model user
	var userModel *models.User
	err = json.Unmarshal([]byte(user), &userModel)
	if err != nil {
		return nil
	}

	return userModel
}

func (u *UserRedisRepositoryImpl) Update(existingUser *models.User, newUser *models.User, ctx context.Context) {
	logrus.Infof("Updating user cache %s", existingUser.ID.String())
	// clear all existing user index
	u.Delete(existingUser, ctx)

	// creating new index
	u.Create(ctx, newUser)
}

func (u *UserRedisRepositoryImpl) Delete(user *models.User, ctx context.Context) {
	logrus.Infof("Deleting user cache %s", user.ID.String())

	// clear all existing user index
	u.redisClient.Del(ctx, "user:"+user.ID.String())
	u.redisClient.Del(ctx, "user:"+user.Username)
	u.redisClient.Del(ctx, "user:"+user.Email)
	u.redisClient.Del(ctx, "user:"+user.Phone)

	return
}
