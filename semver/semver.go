package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	RegexSemver     = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-([0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*))?(\+([0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*))?$`
	RegexPreRelease = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)-([0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)`
)

type SemVer struct {
	major      int
	minor      int
	patch      int
	preRelease string
	build      string
}

func ParseSemver(input string) (SemVer, error) {
	parsed, err := parseSemver(input)
	return parsed, err
}

// CompareSemver returns three values -1, 0, +1
// -1 denotes ver1 < ver2
// 0 denotes invalid input
// +1 denotes ver1 > ver2
func CompareSemver(ver1, ver2 string) (int, error) {
	ok, err := compare(ver1, ver2)
	return ok, err
}

func GetNextMajor(version string) (string, error) {
	parsed, err := parseSemver(version)
	if err != nil {
		return "", err
	}
	major := parsed.major
	// increment the major version and reset minor and patch to 0
	major++
	return fmt.Sprintf("%d.0.0", major), nil
}

func GetNextMinor(version string) (string, error) {
	parsed, err := parseSemver(version)
	if err != nil {
		return "", err
	}
	minor := parsed.minor
	// increment the minor version and reset patch to 0
	minor++
	return fmt.Sprintf("%d.%d.0", parsed.major, minor), nil
}

func GetNextPatch(version string) (string, error) {
	parsed, err := parseSemver(version)
	if err != nil {
		return "", err
	}
	patch := parsed.patch
	// increment the patch version
	patch++
	return fmt.Sprintf("%d.%d.%d", parsed.major, parsed.minor, patch), nil
}

func IsPreRelease(input string) (bool, error) {
	input = strings.TrimPrefix(input, "v")
	input = strings.TrimPrefix(input, " ")
	semverRegex := regexp.MustCompile(RegexPreRelease)
	match := semverRegex.FindStringSubmatch(input)
	if match == nil {
		return false, fmt.Errorf("invalid semantic version string")
	}
	return true, nil
}

func parseSemver(version string) (SemVer, error) {

	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, " ")

	semverRegex := regexp.MustCompile(RegexSemver)
	match := semverRegex.FindStringSubmatch(version)
	if match == nil {
		return SemVer{}, fmt.Errorf("invalid semantic version string")
	}

	major, err := strconv.Atoi(match[1])
	if err != nil {
		return SemVer{}, err
	}

	minor, err := strconv.Atoi(match[2])
	if err != nil {
		return SemVer{}, err
	}

	patch, err := strconv.Atoi(match[3])
	if err != nil {
		return SemVer{}, err
	}

	preRelease := match[5]
	build := match[8]

	return SemVer{
		major:      major,
		minor:      minor,
		patch:      patch,
		preRelease: preRelease,
		build:      build,
	}, nil
}
