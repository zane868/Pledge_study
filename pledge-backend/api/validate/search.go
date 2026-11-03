package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/config"
)

type Search struct{}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) Search(c *gin.Context, req *request.Search) int {

	err := c.ShouldBindJSON(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return statecode.ParameterEmptyErr
		}
		for _, e := range errs {
			if e.Field() == "ChainID" && e.Tag() == "required" {
				return statecode.ChainIdEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	if !config.SupportedChain(req.ChainID) {
		return statecode.ChainIdErr
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 20
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}
	return statecode.CommonSuccess
}
