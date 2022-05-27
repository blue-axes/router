package http

import (
	"github.com/blue-axes/router/entity/constants"
	"github.com/blue-axes/router/pkg/util"
	"github.com/labstack/echo/v4"
)

func Pre(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := util.UUIDString()
		c.Set(constants.CtxKeyRequestID, requestID)
		return next(c)
	}
}
