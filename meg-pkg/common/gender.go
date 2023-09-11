package common

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

func (g Gender) IsValid() bool {
	switch g {
	case Male, Female, Other:
		return true
	default:
		return false
	}
}

func (g Gender) ToString() string {
	return string(g)
}

func IsGenderValid(s string) bool {
	switch s {
	case Male.ToString(), Female.ToString(), Other.ToString():
		return true
	default:
		return false
	}
}
