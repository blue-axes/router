package errno

import (
	"fmt"
)

type (
	Errno uint32
)

const (
	OK                Errno = 200
	NoContent         Errno = 204
	InvalidArgument   Errno = 400
	Unauthorized      Errno = 401
	Forbidden         Errno = 403
	NotFound          Errno = 404
	Timeout           Errno = 408
	EntityToLarge     Errno = 413
	UnknownMediaType  Errno = 415
	ServerError       Errno = 500
	ServerUnavailable Errno = 503

	AlreadyExist Errno = 20001
)

var (
	msgMap = map[Errno]string{
		OK:                "Success",
		NoContent:         "No Content",
		InvalidArgument:   "Argument:%s Is Invalid",
		Unauthorized:      "Operation Unauthorized",
		Forbidden:         "Operation Forbidden",
		NotFound:          "%s Not Found",
		Timeout:           "Operation Timeout",
		EntityToLarge:     "Request Entity Out Of Limit",
		UnknownMediaType:  "Not Support Media Type",
		ServerError:       "Internal Server Error",
		ServerUnavailable: "Server Unavailable",
		AlreadyExist:      "%s Already Exist",
	}
)

func GenErrMsg(no Errno, name string) string {
	return fmt.Sprintf(msgMap[no], name)
}
