package semver

import "fmt"

// compare returns an integer with 3 possible values, -1, 0, +1
func compare(ver1, ver2 string) (int, error) {
	p1, err := parseSemver(ver1)
	if err != nil {
		return 0, err
	}
	p2, err := parseSemver(ver2)
	if err != nil {
		return 0, err
	}

	// compare major version
	if p1.major != p2.major {
		if p1.major > p2.major {
			return 1, nil
		} else {
			return -1, nil
		}
	}

	// compare minor version
	if p1.minor != p2.minor {
		if p1.minor > p2.minor {
			return 1, nil
		} else {
			return -1, nil
		}
	}

	// compare patch version
	if p1.patch != p2.patch {
		if p1.patch > p2.patch {
			return 1, nil
		} else {
			return -1, nil
		}
	}
	return comparePreRelease(p1.preRelease, p2.preRelease)
}

func comparePreRelease(v1, v2 string) (int, error) {
	if v1 == v2 {
		return 0, fmt.Errorf("equal pre-release versions")
	}
	if v1 == "" && v2 != "" {
		return 1, nil
	}
	if v2 == "" && v1 != "" {
		return -1, nil
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
					return -1, nil
				} else {
					return 1, nil
				}
			}
			if ix {
				if len(x) < len(y) {
					return -1, nil
				}
				if len(x) > len(y) {
					return 1, nil
				}
			}
			if x < y {
				return -1, nil
			} else {
				return 1, nil
			}
		}
	}
	if v1 == "" {
		return -1, nil
	} else {
		return 1, nil
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
