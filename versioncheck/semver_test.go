package versioncheck

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// table test for comparing versions
func Test_SemVer_table(t *testing.T) {
	var v1 SemVer
	var v2 string
	tableTestCases := []struct {
		expected int
		version1 string
		version2 string
	}{
		{0, "0.1", "0.1"},
		{-1, "0.1", "1.1"},
		{-1, "0.1", "1.1"},
		{-1, "0.1", "1.1"},
		{-1, "1.2", "1.2.9"},
		{-1, "1.3", "1.3.4"},
		{-1, "0.1", "0.10"},
		{0, "1.0.0.1", "1.0.0.1"},
		{0, "1.0.0.1", "1.0.0.1.0"},
	}

	for k, testCase := range tableTestCases {
		v1 = New(testCase.version1)
		v2 = testCase.version2
		assert.Equal(t,
			testCase.expected,
			v1.Compare(v2),
			fmt.Sprintf("test case #%d failed comparing %s and %s, got: %d expected: %d",
				k, v1, v2, v1.Compare(v2), testCase.expected))
	}
}
