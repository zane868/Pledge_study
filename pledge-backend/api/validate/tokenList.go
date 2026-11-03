package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/config"
)

type TokenList struct{}

func NewTokenList() *TokenList {
	return &TokenList{}
}

func (v *TokenList) TokenList(c *gin.Context, req *request.TokenList) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return statecode.ParameterEmptyErr
		}
		for _, e := range errs {
			if e.Field() == "ChainId" && e.Tag() == "required" {
				return statecode.ChainIdEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	if !config.SupportedChain(req.ChainId) {
		return statecode.ChainIdErr
	}

	return statecode.CommonSuccess
}
