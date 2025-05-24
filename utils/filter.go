package utils

func FilterEmptyStrings(arr []string) []string {
	var filtered []string
	for _, str := range arr {
		if len(str) > 0 {
			filtered = append(filtered, str)
		}
	}
	return filtered
}