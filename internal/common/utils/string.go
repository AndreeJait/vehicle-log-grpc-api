package utils

func PtrStr(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}
