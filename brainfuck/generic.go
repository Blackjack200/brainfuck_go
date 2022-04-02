package brainfuck

func inArray[T comparable](arr []T, elem T) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
