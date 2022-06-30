package versioncheck

import (
	"strings"
)

// SemVer represents a versioning scheme
type SemVer string

// New returns a new SemVer
func New(v string) SemVer {
	return SemVer(v)
}

func (s SemVer) String() string {
	return string(s)
}

// Compare returns a
// 0 if receiver and s2 are the same version.
// 1 if receiver is higher
// -1 if receiver is lower
func (s SemVer) Compare(s2 SemVer) int {
	versions1 := strings.Split(s.String(), ".")
	versions2 := strings.Split(s2.String(), ".")
	len1 := len(versions1)
	len2 := len(versions2)
	var minlengthForComparision, max, i int
	if len1 <= len2 {
		minlengthForComparision = len1
		max = len2
	} else {
		minlengthForComparision = len2
		max = len1
	}

	for i = 0; i < max; i++ {
		// compare versions until we can't
		if i < minlengthForComparision {
			if versions1[i] == versions2[i] {
				continue
			} else if versions1[i] > versions2[i] {
				return 1
			} else if versions1[i] < versions2[i] {
				return -1
			}
		} else {
			// longer version with non-zero patch is bigger
			if len1 > len2 && versions1[i] != "0" {
				return 1
			} else if len1 < len2 && versions2[i] != "0" {
				return -1
			}
		}
	}

	return 0
}
