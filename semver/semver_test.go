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
			name:    "ParseSemver_1",
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
			name: "CompareSemver_1",
			ver1: "v1.2.3",
			ver2: "v1.2.4",
			want: -1,
		},
		{
			name: "CompareSemver_2",
			ver1: "v1.0.0-alpha",
			ver2: "v1.0.0-alpha.1",
			want: -1,
		},
		{
			name: "CompareSemver_2",
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
			name:    "GetNextMajor_1",
			version: "v1.2.3",
			want:    nil,
		},
		{
			name:    "GetNextMajor_2",
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
