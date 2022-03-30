package commonresp

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       ApiResponse
}

func (j *JsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}
func (j *JsonResponse) Get() (int, ApiResponse) {
	return j.httpStatusCode, j.response
}
func NewSuccessJsonResponse(c *gin.Context, data interface{}) AppHttpResponse {
	httpStatusCode, resp := NewSuccessMessage(data)
	return &JsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}

func NewErrorJsonResponse(c *gin.Context, err error) AppHttpResponse {
	fmt.Println("===>", err)
	httpStatusCode, resp := NewErrorMessage(err)
	return &JsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}

func NewJsonResponse(c *gin.Context, httpStatusCode int, response ApiResponse) AppHttpResponse {
	return &JsonResponse{
		c,
		httpStatusCode,
		response,
	}
}
