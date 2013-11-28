package sss

func equal(a, b polynomial) bool {
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
	return false
}
