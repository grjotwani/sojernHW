package versioncheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Receiver_lower_version(t *testing.T) {
	v1 := New("5.0.1.11.3.2")
	assert.Equal(t, -1, v1.Compare("5.0.1.11.4"))
}

func Test_Receiver_lower_version2(t *testing.T) {
	v1 := New("5.0.1")
	assert.Equal(t, -1, v1.Compare("5.0.1.1"))
}

func Test_same_version(t *testing.T) {
	v1 := New("1.0.0.1")
	assert.Equal(t, 0, v1.Compare("1.0.0.1"))
}

func Test_same_version2(t *testing.T) {

	v1 := New("5.0.1")
	assert.Equal(t, 0, v1.Compare("5.0.1.0"))
}
