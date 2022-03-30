package api

import (
	"enigmacamp.com/goacc/delivery/commonresp"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) Success(c *gin.Context, data interface{}) {
	commonresp.NewSuccessJsonResponse(c, data).Send()
}

func (b *BaseApi) Failed(c *gin.Context, err error) {
	c.Error(err)
}
