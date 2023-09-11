package meghttp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type MegError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
	IsSentry  bool
}

func (e MegError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

func (e MegError) Is(target error) bool {
	if e.Raw != nil {
		return errors.Is(e.Raw, target)
	}

	return strings.Contains(e.Error(), target.Error())
}

func NewError(err error, httpCode int, errCode string, message string, isSentry bool) MegError {
	return MegError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
		IsSentry:  isSentry,
	}
}

func ErrUnauthorized(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "000001",
		Message:   fmt.Sprintf("%s: Unauthorized!", module),
		IsSentry:  false,
	}
}

func ErrNoPermission() MegError {
	return MegError{
		Raw:       nil,
		HTTPCode:  http.StatusForbidden,
		ErrorCode: "000002",
		Message:   "No permission.",
		IsSentry:  false,
	}
}

func ErrInvalidParams(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "000003",
		Message:   fmt.Sprintf("%s: Invalid param.", module),
		IsSentry:  false,
	}
}

func ErrNotFound(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "000004",
		Message:   fmt.Sprintf("%s: Not found.", module),
		IsSentry:  false,
	}
}

func ErrBind(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10006",
		Message:   fmt.Sprintf("%s: bind error.", module),
		IsSentry:  false,
	}
}

func ErrParseInt(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10007",
		Message:   fmt.Sprintf("%s: parse int error.", module),
		IsSentry:  false,
	}
}

func ErrGet(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000005",
		Message:   fmt.Sprintf("%s: Failed to get.", module),
		IsSentry:  true,
	}
}

func ErrCreate(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000006",
		Message:   fmt.Sprintf("%s: Failed to create.", module),
		IsSentry:  true,
	}
}

func ErrUpdate(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000007",
		Message:   fmt.Sprintf("%s: Failed to update.", module),
		IsSentry:  true,
	}
}

func ErrDelete(err error, module string) MegError {
	return MegError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000008",
		Message:   fmt.Sprintf("%s: Failed to delete.", module),
		IsSentry:  true,
	}
}
