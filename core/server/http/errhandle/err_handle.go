package errhandle

import (
	"github.com/blue-axes/router/entity/constants"
	"github.com/blue-axes/router/entity/errno"
	"github.com/blue-axes/router/pkg/errors"
	"github.com/labstack/echo/v4"
)

type (
	ResponseData struct {
		Code      uint32      `json:"Code"`
		Message   string      `json:"Message"`
		RequestID string      `json:"RequestID"`
		Data      interface{} `json:"Data;omitempty"`
	}
)

func ErrorHandler(err error, c echo.Context) {
	requestID, _ := c.Get(constants.CtxKeyRequestID).(string)

	resp := ResponseData{}
	switch valErr := err.(type) {
	case *errors.Error:
		resp = ResponseData{
			Code:      uint32(valErr.Code()),
			Message:   valErr.Message(),
			RequestID: requestID,
		}
	default:
		resp = ResponseData{
			Code:      uint32(errno.ServerError),
			Message:   valErr.Error(),
			RequestID: requestID,
		}
	}
	_ = c.JSON(int(resp.Code), resp)
}
