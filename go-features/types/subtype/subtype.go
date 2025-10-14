package subtype

type SString string

func (s SString) String() string {
	return string(s)
}
