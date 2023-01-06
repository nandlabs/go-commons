package semver

import (
	"fmt"
	"testing"
)

func TestSemver(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{
			name:    "TestParseSemver_1",
			version: "v1.2.3",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSemver(tt.version)
			if got != tt.want {
				t.Errorf("Error in parsing version :: got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestCompareSemver(t *testing.T) {
	tests := []struct {
		name string
		ver1 string
		ver2 string
		want int
	}{
		{
			name: "TestCompareSemver_1",
			ver1: "v1.2.3",
			ver2: "v1.2.4",
			want: -1,
		},
		{
			name: "TestCompareSemver_2",
			ver1: "v1.0.0-alpha",
			ver2: "v1.0.0-alpha.1",
			want: -1,
		},
		{
			name: "TestCompareSemver_2",
			ver1: "v1.0.0-alpha.1",
			ver2: "v1.0.0-alpha",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareSemver(tt.ver1, tt.ver2)
			if got != tt.want {
				t.Errorf("Error in comparing version :: got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestGetNextMajor(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    error
	}{
		{
			name:    "TestGetNextMajor_1",
			version: "v1.2.3",
			want:    nil,
		},
		{
			name:    "TestGetNextMajor_2",
			version: "v9.1.1",
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextMajor(tt.version)
			if tt.want != nil && err.Error() != tt.want.Error() {
				t.Errorf("Error in comparing version :: got %d, want %d", err, tt.want)
			}
			if got != "" {
				fmt.Println(got)
			}
		})
	}
}

func TestGetNextMinor(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    error
	}{
		{
			name:    "TestGetNextMinor_1",
			version: "v1.2.3",
			want:    nil,
		},
		{
			name:    "TestGetNextMinor_2",
			version: "v9.1.1",
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextMinor(tt.version)
			if tt.want != nil && err.Error() != tt.want.Error() {
				t.Errorf("Error in comparing version :: got %d, want %d", err, tt.want)
			}
			if got != "" {
				fmt.Println(got)
			}
		})
	}
}

func TestGetNextPatch(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    error
	}{
		{
			name:    "TestGetNextPatch_1",
			version: "v1.2.3",
			want:    nil,
		},
		{
			name:    "TestGetNextPatch_2",
			version: "v9.1.1",
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextPatch(tt.version)
			if tt.want != nil && err.Error() != tt.want.Error() {
				t.Errorf("Error in comparing version :: got %d, want %d", err, tt.want)
			}
			if got != "" {
				fmt.Println(got)
			}
		})
	}
}

func TestIsPreRelease(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{
			name:    "TestIsPreRelease_1",
			version: "v1.2.3-beta",
			want:    true,
		},
		{
			name:    "TestIsPreRelease_2",
			version: "v1.2.3",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPreRelease(tt.version)
			if tt.want != got {
				t.Errorf("Error in testing IsPreRelease :: got %t, want %t", got, tt.want)
			}
		})
	}
}
