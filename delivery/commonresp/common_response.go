package commonresp

import (
	"enigmacamp.com/goacc/utils"
	"errors"
	"net/http"
)

const (
	SuccessCode         = "00"
	SuccessMessage      = "Success"
	DefaultErrorCode    = "XX"
	DefaultErrorMessage = "Something went wrong"
)

type AppHttpResponse interface {
	Send()
	Get() (int, ApiResponse)
}

type Status struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

type ApiResponse struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessMessage(data interface{}) (httpStatusCode int, apiResponse ApiResponse) {
	status := Status{
		ResponseCode:    SuccessCode,
		ResponseMessage: SuccessMessage,
	}
	httpStatusCode = http.StatusOK
	apiResponse = ApiResponse{
		status, data,
	}
	return
}
func NewErrorMessage(err error) (httpStatusCode int, apiResponse ApiResponse) {
	var userError *utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponseCode:    userError.ErrorCode,
			ResponseMessage: userError.ErrorMessage,
		}
		httpStatusCode = userError.ErrorType
	} else {
		status = Status{
			ResponseCode:    DefaultErrorCode,
			ResponseMessage: DefaultErrorMessage,
		}
		httpStatusCode = http.StatusInternalServerError
	}
	apiResponse = ApiResponse{
		status, nil,
	}

	return
}
