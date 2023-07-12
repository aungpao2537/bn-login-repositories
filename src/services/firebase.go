package services

type IireBaseService struct {
	FirebaseRepository repositories.ISpeakerRepository
	RedisRepository    repositories.IRedisConnectionRepository
}

type IFireBaseService interface {
}

func NewFirebaseService(repo0 repositories.IFireBaseService, cache0 repositories.IRedisConnectionRepository) ISpeakerService {
	return &speakerService{
		SpeakerRepository: repo0,
		RedisRepository:   cache0,
	}
}
