package main

func isEqual(s []string, t []string) bool {
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}
