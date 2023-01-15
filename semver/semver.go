package semver

import (
	"fmt"
	"regexp"
	"strconv"
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

func CompareSemver(ver1, ver2 string) int {
	ok := compare(ver1, ver2)
	return ok
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
	semverRegex := regexp.MustCompile(RegexPreRelease)
	match := semverRegex.FindStringSubmatch(input)
	if match == nil {
		return false, fmt.Errorf("invalid semantic version string")
	}
	return true, nil
}

func parseSemver(version string) (SemVer, error) {
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
	build := match[9]

	return SemVer{
		major:      major,
		minor:      minor,
		patch:      patch,
		preRelease: preRelease,
		build:      build,
	}, nil
}

//func buildSemver(input version) (string, error) {
//	versionArr := []string{input.major, input.minor, input.patch}
//	builtVersion := strings.Join(versionArr, ".")
//	builtVersion = "v" + builtVersion
//	if input.preRelease != "" {
//		builtVersion = builtVersion + "-" + input.preRelease
//	}
//	if input.build != "" {
//		builtVersion = builtVersion + "+" + input.build
//	}
//	return builtVersion, nil
//}
