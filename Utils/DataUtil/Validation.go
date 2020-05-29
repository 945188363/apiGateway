package DataUtil

func Validate(pattern string, values ...interface{}) bool {
	switch pattern {
	case "NULL":
		for _, value := range values {
			if value == nil {
				return false
			}
		}
	}
	return true
}
