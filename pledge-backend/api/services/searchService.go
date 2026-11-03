package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/log"
)

type SearchService struct{}

func NewSearch() *SearchService {
	return &SearchService{}
}

func (c *SearchService) Search(req *request.Search) (int, int64, []models.Pool) {

	err, total, data := models.NewPool().Pagination(req)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr, 0, nil
	}
	return 0, total, data
}
