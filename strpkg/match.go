package strpkg

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func match(pattern, str string, noCase bool) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		case '*':
			for len(pattern) > 1 && pattern[1] == '*' {
				pattern = pattern[1:]
			}
			if len(pattern) == 1 {
				return true /* match */
			}
			for len(str) > 0 {
				if match(pattern[1:], str, noCase) {
					return true /* match */
				}
				str = str[1:]
			}
			return false /* no match */
		case '?':
			if len(str) == 0 {
				return false /* no match */
			}
			str = str[1:]
		case '[':
			var not, match bool
			pattern = pattern[1:]
			not = pattern[0] == '^'
			if not {
				pattern = pattern[1:]
			}
			for {
				if pattern[0] == '\\' {
					pattern = pattern[1:]
					if pattern[0] == str[0] {
						match = true
					}
				} else if pattern[0] == ']' {
					break
				} else if len(pattern) == 0 {
					pattern = pattern[1:]
					break
				} else if pattern[1] == '-' && len(pattern) >= 3 {
					var start = pattern[0]
					var end = pattern[2]
					var c = str[0]
					if start > end {
						start, end = end, start
					}
					if noCase {
						start = toLower(start)
						end = toLower(end)
						c = toLower(c)
					}
					pattern = pattern[2:]
					if c >= start && c <= end {
						match = true
					}
				} else {
					if !noCase {
						if pattern[0] == str[0] {
							match = true
						}
					} else {
						if toLower(pattern[0]) == toLower(str[0]) {
							match = true
						}
					}
				}
				pattern = pattern[1:]
			}
			if not {
				match = !match
			}
			if !match {
				return false /* no match */
			}
			str = str[1:]
		case '\\':
			if len(pattern) >= 2 {
				pattern = pattern[1:]
			}
			fallthrough
			/* fall through */
		default:
			if !noCase {
				if pattern[0] != str[0] {
					return false /* no match */
				}
			} else {
				if toLower(pattern[0]) != toLower(str[0]) {
					return false /* no match */
				}
			}
			str = str[1:]
			break
		}
		pattern = pattern[1:]
		if len(str) == 0 {
			for len(pattern) > 0 && pattern[0] == '*' {
				pattern = pattern[1:]
			}
			break
		}
	}
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}
	return false
}

// Check str is match the pattern
func Match(pattern string, str string) bool {
	if pattern == "*" {
		return true
	}
	return match(pattern, str, false)
}
