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

func (s *SemVer) String() string {
	if s.preRelease != "" && s.build != "" {
		return fmt.Sprintf("%d.%d.%d-%s+%s", s.major, s.minor, s.patch, s.preRelease, s.build)
	} else if s.preRelease != "" {
		return fmt.Sprintf("%d.%d.%d-%s", s.major, s.minor, s.patch, s.preRelease)
	} else if s.build != "" {
		return fmt.Sprintf("%d.%d.%d+%s", s.major, s.minor, s.patch, s.build)
	}
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}

func Parse(input string) (*SemVer, error) {
	parsed, err := parse(input)
	return parsed, err
}

// CompareRaw returns three values -1, 0, +1
// -1 denotes ver1 < ver2
// 0 denotes invalid input
// +1 denotes ver1 > ver2
func CompareRaw(ver1, ver2 string) (int, error) {
	c1, err := Parse(ver1)
	if err != nil {
		return 0, err
	}
	c2, err := Parse(ver2)
	if err != nil {
		return 0, err
	}
	ok, err := compare(c1, c2)
	return ok, err
}

// Compare returns three values -1, 0, +1
// -1 denotes ver1 < ver2
// 0 denotes invalid input
// +1 denotes ver1 > ver2
func (s *SemVer) Compare(v *SemVer) (int, error) {
	ok, err := compare(s, v)
	return ok, err
}

func (s *SemVer) GetNextMajor() string {
	major := s.major
	// increment the major version and reset minor and patch to 0
	major++
	s.major = major
	s.minor = 0
	s.patch = 0
	return s.String()
}

func (s *SemVer) GetNextMinor() string {
	minor := s.minor
	// increment the minor version and reset patch to 0
	minor++
	s.minor = minor
	s.patch = 0
	return s.String()
}

func (s *SemVer) GetNextPatch() string {
	patch := s.patch
	// increment the patch version
	patch++
	s.patch = patch
	return s.String()
}

func (s *SemVer) IsPreRelease() bool {
	input := s.String()
	input = strings.TrimPrefix(input, "v")
	input = strings.TrimPrefix(input, " ")
	semverRegex := regexp.MustCompile(RegexPreRelease)
	match := semverRegex.FindStringSubmatch(input)
	if match == nil {
		return false
	}
	return true
}

func parse(version string) (*SemVer, error) {

	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, " ")

	semverRegex := regexp.MustCompile(RegexSemver)
	match := semverRegex.FindStringSubmatch(version)
	if match == nil {
		return &SemVer{}, fmt.Errorf("invalid semantic version string")
	}

	major, err := strconv.Atoi(match[1])
	if err != nil {
		return &SemVer{}, err
	}

	minor, err := strconv.Atoi(match[2])
	if err != nil {
		return &SemVer{}, err
	}

	patch, err := strconv.Atoi(match[3])
	if err != nil {
		return &SemVer{}, err
	}

	preRelease := match[5]
	build := match[8]

	return &SemVer{
		major:      major,
		minor:      minor,
		patch:      patch,
		preRelease: preRelease,
		build:      build,
	}, nil
}
