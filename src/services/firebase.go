package services

import (
	"github.com/aungpao2537/bn-login-repositories/domain/repositories"
)

type FireBaseService struct {
	FirebaseRepository repositories.IFirebaseRepository
	RedisRepository    repositories.IRedisConnectionRepository
}

type IFireBaseService interface {
}

func NewFirebaseService(repo0 repositories.IFirebaseRepository, cache0 repositories.IRedisConnectionRepository) IFireBaseService {
	return &FireBaseService{
		FirebaseRepository: repo0,
		RedisRepository:    cache0,
	}
}
