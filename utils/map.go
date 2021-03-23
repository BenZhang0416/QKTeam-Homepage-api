package utils

func MergeMap(source map[string]interface{}, target map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range source {
		result[key] = value
	}
	for key, value := range target {
		result[key] = value
	}
	return result
}
