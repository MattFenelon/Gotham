package domain

type TrimmedString string

func ToTrimmedString(value string) TrimmedString {
	return TrimmedString(value)
}

func (t TrimmedString) String() string {
	return string(t)
}
