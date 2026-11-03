package services

import (
	"crypto/subtle"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/utils"
)

type UserService struct{}

func NewUser() *UserService {
	return &UserService{}
}

func (s *UserService) Login(req *request.Login, result *response.Login) int {

	if req.Name == config.Config.DefaultAdmin.Username && len(config.Config.DefaultAdmin.Password) >= 12 && subtle.ConstantTimeCompare([]byte(req.Password), []byte(config.Config.DefaultAdmin.Password)) == 1 {
		token, err := utils.CreateToken(req.Name)
		if err != nil {
			log.Logger.Error("CreateToken" + err.Error())
			return statecode.CommonErrServerErr
		}
		result.TokenId = token
		//save to redis
		if err := db.RedisSetString("session:"+req.Name, token, config.Config.Jwt.ExpireTime); err != nil {
			return statecode.CommonErrServerErr
		}
		return statecode.CommonSuccess
	} else {
		return statecode.NameOrPasswordErr
	}
}
