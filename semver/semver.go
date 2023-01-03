package semver

import (
	"errors"
	"strings"
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
	major, err := processNextVersion(parsed.major)
	if err != nil {
		return "", err
	}
	parsed.major = major
	builtVersion, err := buildSemver(parsed)
	if err != nil {
		return "", err
	}
	return builtVersion, nil
}

func GetNextMinor(version string) (string, error) {
	parsed, ok := parse(version)
	if !ok {
		return "", errors.New("error parsing semantic version")
	}
	minor, err := processNextVersion(parsed.minor)
	if err != nil {
		return "", err
	}
	parsed.minor = minor
	builtVersion, err := buildSemver(parsed)
	if err != nil {
		return "", err
	}
	return builtVersion, nil
}

func GetNextPatch(version string) (string, error) {
	parsed, ok := parse(version)
	if !ok {
		return "", errors.New("error parsing semantic version")
	}
	patch, err := processNextVersion(parsed.patch)
	if err != nil {
		return "", err
	}
	parsed.minor = patch
	builtVersion, err := buildSemver(parsed)
	if err != nil {
		return "", err
	}
	return builtVersion, nil
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

func buildSemver(input version) (string, error) {
	versionArr := []string{input.major, input.minor, input.patch}
	builtVersion := strings.Join(versionArr, ".")
	builtVersion = "v" + builtVersion
	if input.preRelease != "" {
		builtVersion = builtVersion + "-" + input.preRelease
	}
	if input.build != "" {
		builtVersion = builtVersion + "+" + input.build
	}
	return builtVersion, nil
}
