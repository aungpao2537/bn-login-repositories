package services

import (
	"bn-login-repositories/domain/repositories"
)

type loginService struct {
	LoginRepository repositories.IUserRepository
	// RedisRepository repositories.IRedisConnectionRepository
}

type ILoginService interface {
	Liff_Authentication() string
}

func NewLoginService(repo0 repositories.IUserRepository) ILoginService {
	return &loginService{
		LoginRepository: repo0,
		// RedisRepository: cache0,
	}
}

func (lsv loginService) Liff_Authentication() string {
	return "ok222"
	// redisData := lsv.RedisRepository.GetSpeakerData()
	// if redisData == nil {
	// 	mongoData, err := lsv.SpeakerRepository.FindAllSpeaker()
	// 	if err != nil {
	// 		log.Print("Error Cannot Get Speaker Data")
	// 		return nil, err
	// 	}
	// 	return mongoData, nil
	// }
	// return redisData, nil
}
