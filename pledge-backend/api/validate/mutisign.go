package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/config"
)

type MutiSign struct{}

func NewMutiSign() *MutiSign {
	return &MutiSign{}
}

func (v *MutiSign) SetMultiSign(c *gin.Context, req *request.SetMultiSign) int {

	err := c.ShouldBind(req)
	if !config.SupportedChain(req.ChainId) {
		return statecode.ChainIdErr
	}
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return statecode.ParameterEmptyErr
		}
		for _, e := range errs {
			if e.Field() == "SpName" && e.Tag() == "required" {
				return statecode.PNameEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}

func (v *MutiSign) GetMultiSign(c *gin.Context, req *request.GetMultiSign) int {

	err := c.ShouldBind(req)
	if !config.SupportedChain(req.ChainId) {
		return statecode.ChainIdErr
	}
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

	return statecode.CommonSuccess
}
