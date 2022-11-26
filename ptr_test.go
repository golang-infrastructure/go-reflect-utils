package reflect_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUintPtrE(t *testing.T) {
	s := ""
	ptr, err := GetUintPtrE(s)
	assert.NotNil(t, err)
	assert.Equal(t, uintptr(0), ptr)

	ptr, err = GetUintPtrE(&s)
	assert.Nil(t, err)
	t.Log(ptr)

}

func TestGetUintPtr(t *testing.T) {

}
