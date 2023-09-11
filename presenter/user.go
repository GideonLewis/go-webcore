package presenter

import "time"

type UserResponseWrapper struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int64     `json:"age"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListUserResponseWrapper struct {
	Users []UserResponseWrapper  `json:"users"`
	Meta  map[string]interface{} `json:"meta"`
}
