package eliminate

func EliminateAdjacentDups(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	pos := 0
	for i, s := range strings {
		if s != strings[pos] && pos != i {
			pos++
			strings[pos] = s
		}
	}
	return strings[:pos+1]
}
