package payload

import "github.com/megaqstar/web-core/meg-pkg/common"

type UserCreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Age       int64  `json:"age"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

type UserUpdateRequest struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Age       int64  `json:"age"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

type UserGetByIDRequest struct {
	UserID int64 `json:"userId" query:"userId"`
}

type UserGetListRequest struct {
	Paginator  common.Paginator   `json:"paginator,omitempty"`
	Conditions []common.Condition `json:"conditions,omitempty"`
	OrderBy    common.Order       `json:"orderBy,omitempty"`
	Unscoped   bool               `json:"unscoped,omitempty"`
}

type UserDeleteRequest struct {
	UserID int64 `json:"userId" query:"userId"`
}
