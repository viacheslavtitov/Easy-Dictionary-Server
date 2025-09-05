package utils

func Deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func DerefInt(s *int) int {
	if s == nil {
		return 0
	}
	return *s
}
