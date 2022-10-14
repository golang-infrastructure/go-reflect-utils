package reflect_utils

import (
	"github.com/CC11001100/go-reflect-util/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStructPtrUnExportedField(t *testing.T) {

	bar1 := "CC11001100"
	bar2 := "bar2"

	// case 001. 尝试从struct上获取一个存在的未导出字段
	field := GetStructPtrUnExportedField(test.NewFoo(bar1, bar2), "bar1")
	assert.Equal(t, bar1, field.String())

	// case 002. 尝试从struct上获取一个存在的导出字段
	field = GetStructPtrUnExportedField(test.NewFoo(bar1, bar2), "Bar2")
	assert.Equal(t, bar2, field.String())

	// case 003. 尝试从struct上获取一个不存在的字段
	field = GetStructPtrUnExportedField(test.NewFoo(bar1, bar2), "bar3")
	assert.False(t, field.IsValid())

}
