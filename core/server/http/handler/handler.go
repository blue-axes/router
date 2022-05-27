package handler

import (
	"github.com/blue-axes/router/core/server/http/errhandle"
	"github.com/blue-axes/router/entity/constants"
	"github.com/blue-axes/router/entity/errno"
	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
	}
)

func (h Handler) Resp(c echo.Context, data interface{}) error {
	return h.RespWithError(c, nil, data)
}

func (h Handler) RespWithError(c echo.Context, err error, data interface{}) error {
	if err != nil {
		errhandle.ErrorHandler(err, c)
		return nil
	}
	requestID, _ := c.Get(constants.CtxKeyRequestID).(string)

	resp := errhandle.ResponseData{
		Code:      uint32(errno.OK),
		Message:   errno.GenErrMsg(errno.OK, ""),
		RequestID: requestID,
		Data:      data,
	}

	return c.JSON(int(errno.OK), resp)
}
