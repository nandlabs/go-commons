package semver

func isIdentChar(c byte) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '-'
}

func isBadNum(input string) bool {
	i := 0
	for i < len(input) && '0' <= input[i] && input[i] <= '9' {
		i++
	}
	return i == len(input) && i > 1 && input[0] == '0'
}

// compare returns an integer with 3 possible values, -1, 0, +1
func compare(ver1, ver2 string) int {
	p1, ok1 := parse(ver1)
	p2, ok2 := parse(ver2)
	if !ok1 && !ok2 {
		return 0
	}
	if !ok1 {
		return -1
	}
	if !ok2 {
		return 1
	}
	if comp := compareInt(p1.major, p2.major); comp != 0 {
		return comp
	}
	if comp := compareInt(p1.minor, p2.minor); comp != 0 {
		return comp
	}
	if comp := compareInt(p1.patch, p2.patch); comp != 0 {
		return comp
	}
	return comparePreRelease(p1.preRelease, p2.preRelease)
}

func compareInt(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}
	if len(v1) < len(v2) {
		return -1
	}
	if len(v1) > len(v2) {
		return 1
	}
	if v1 < v2 {
		return -1
	} else {
		return 1
	}
}

func comparePreRelease(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}
	if v1 == "" {
		return 1
	}
	if v2 == "" {
		return -1
	}

	for v1 != "" && v2 != "" {
		v1 = v1[1:]
		v2 = v2[1:]
		var x, y string
		x, v1 = nextIdent(v1)
		y, v2 = nextIdent(v2)
		if x != y {
			ix := isNum(x)
			iy := isNum(y)
			if ix != iy {
				if ix {
					return -1
				} else {
					return 1
				}
			}
			if ix {
				if len(x) < len(y) {
					return -1
				}
				if len(x) > len(y) {
					return 1
				}
			}
			if x < y {
				return -1
			} else {
				return 1
			}
		}
	}
	if v1 == "" {
		return -1
	} else {
		return 1
	}
}

func nextIdent(x string) (dx, rest string) {
	i := 0
	for i < len(x) && x[i] != '.' {
		i++
	}
	return x[:i], x[i:]
}

func isNum(v string) bool {
	i := 0
	for i < len(v) && '0' <= v[i] && v[i] <= '9' {
		i++
	}
	return i == len(v)
}
