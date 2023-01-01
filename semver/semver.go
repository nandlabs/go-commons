package semver

import (
	"errors"
	"strconv"
)

type version struct {
	major      string
	minor      string
	patch      string
	preRelease string
	build      string
}

func ParseSemver(input string) bool {
	_, ok := parse(input)
	return ok
}

func CompareSemver(ver1, ver2 string) int {
	ok := compare(ver1, ver2)
	return ok
}

func GetNextMajor(version string) (string, error) {
	parsed, ok := parse(version)
	if !ok {
		return "", errors.New("error parsing semantic version")
	}
	if parsed.major > "0" && parsed.major < "9" {
		nm, err := strconv.Atoi(parsed.major)
		if err != nil {

		}
		nm = nm + 1
		return strconv.Itoa(nm), nil
	} else {
		return "", errors.New("cannot generate next acceptable major version")
	}
}

func GetNextMinor() {

}

func GetNextPatch() {

}

func IsPreRelease() bool {
	return false
}

func parse(input string) (v version, ok bool) {
	if input == "" || input[0] != 'v' {
		return
	}
	v.major, input, ok = parseInt(input[1:])
	if !ok {
		return
	}
	if input == "" {
		v.minor = "0"
		v.patch = "0"
		return
	}
	if input[0] != '.' {
		ok = false
		return
	}
	v.minor, input, ok = parseInt(input[1:])
	if !ok {
		return
	}
	if input == "" {
		v.patch = "0"
		return
	}
	if input[0] != '.' {
		ok = false
		return
	}
	v.patch, input, ok = parseInt(input[1:])
	if !ok {
		return
	}
	if len(input) > 0 && input[0] == '-' {
		v.preRelease, input, ok = parsePreRelease(input)
		if !ok {
			return
		}
	}
	if len(input) > 0 && input[0] == '+' {
		v.build, input, ok = parseBuild(input)
		if !ok {
			return
		}
	}
	if input != "" {
		ok = false
		return
	}
	ok = true
	return
}

func buildSemver(input string) (string, error) {

}
