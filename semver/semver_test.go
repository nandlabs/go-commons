package semver

import (
	"testing"
)

func TestSemver(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    SemVer
	}{
		{
			name:    "TestParseSemver_1",
			version: "v1.2.3",
			want: SemVer{
				major:      1,
				minor:      2,
				patch:      3,
				preRelease: "",
				build:      "",
			},
		},
		{
			name:    "TestParseSemver_2",
			version: "1.2.3",
			want: SemVer{
				major:      1,
				minor:      2,
				patch:      3,
				preRelease: "",
				build:      "",
			},
		},
		{
			name:    "TestParseSemver_3",
			version: "1.2.3-alpha.1+build.1",
			want: SemVer{
				major:      1,
				minor:      2,
				patch:      3,
				preRelease: "alpha.1",
				build:      ".1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSemver(tt.version)
			if err != nil {
				t.Errorf("Error in parsing version :: got %t", err)
			}
			if got != tt.want {
				t.Errorf("Invalid output :: want %+v, got :: %+v", tt.want, got)
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
			name: "TestCompareSemver_3",
			ver1: "v1.0.0-alpha.1",
			ver2: "v1.0.0-alpha",
			want: 1,
		},
		{
			name: "TestCompareSemver_4",
			ver1: "v1.2.5",
			ver2: "v1.2.4",
			want: 1,
		},
		{
			name: "TestCompareSemver_5",
			ver1: "v1.3.5",
			ver2: "v1.2.4",
			want: 1,
		},
		{
			name: "TestCompareSemver_6",
			ver1: "v2.3.5",
			ver2: "v1.2.4",
			want: 1,
		},
		{
			name: "TestCompareSemver_7",
			ver1: "v1.2.5",
			ver2: "v1.3.4",
			want: -1,
		},
		{
			name: "TestCompareSemver_8",
			ver1: "v1.2.5",
			ver2: "v2.3.4",
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareSemver(tt.ver1, tt.ver2)
			if err != nil {
				t.Errorf("Error comparing version :: got %t", err)
			}
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
		want    string
	}{
		{
			name:    "TestGetNextMajor_1",
			version: "v1.2.3",
			want:    "2.0.0",
		},
		{
			name:    "TestGetNextMajor_2",
			version: "v9.1.1",
			want:    "10.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextMajor(tt.version)
			if err != nil {
				t.Errorf("Error fetching next major :: got %t", err)
			}
			if got != tt.want {
				t.Errorf("invalid next major :: got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextMinor(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
	}{
		{
			name:    "TestGetNextMinor_1",
			version: "v1.2.3",
			want:    "1.3.0",
		},
		{
			name:    "TestGetNextMinor_2",
			version: "v9.1.1",
			want:    "9.2.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextMinor(tt.version)
			if err != nil {
				t.Errorf("Error fetching next minor :: got %t", err)
			}
			if got != tt.want {
				t.Errorf("invalid next minor :: got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextPatch(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
	}{
		{
			name:    "TestGetNextPatch_1",
			version: "v1.2.3",
			want:    "1.2.4",
		},
		{
			name:    "TestGetNextPatch_2",
			version: "v9.1.1",
			want:    "9.1.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextPatch(tt.version)
			if err != nil {
				t.Errorf("Error fetching next patch :: got %t", err)
			}
			if got != tt.want {
				t.Errorf("invalid next patch :: got %v, want %v", got, tt.want)
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
			version: "v1.2.3-beta.1",
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
			got, _ := IsPreRelease(tt.version)
			if tt.want != got {
				t.Errorf("Error in testing IsPreRelease :: got %t, want %t", got, tt.want)
			}
		})
	}
}
