package common

import "strings"

type Order struct {
	Field string   `json:"field"`
	Sort  SortType `json:"sort"`
}

type SortType string

const (
	SortTypeASC  SortType = "ASC"
	SortTypeDESC SortType = "DESC"
)

func (st SortType) IsValid() bool {
	switch SortType(strings.ToUpper(string(st))) {
	case SortTypeASC, SortTypeDESC:
		return true
	default:
		return false
	}
}

func (st *SortType) Format() {
	if !st.IsValid() {
		*st = SortTypeASC
	}
}
