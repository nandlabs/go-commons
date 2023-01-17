package semver

import (
	"fmt"
	"strings"
)

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

	pre1 := len(v1) > 1
	pre2 := len(v2) > 1

	if pre1 && pre2 {
		return strings.Compare(v1, v2), nil
	}
	if pre1 {
		return -1, nil
	}

	if pre2 {
		return 1, nil
	}

	return 0, fmt.Errorf("no pre-release versions present")
}
