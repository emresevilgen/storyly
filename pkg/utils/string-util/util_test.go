package string_util_test

import (
	testifyAssert "github.com/stretchr/testify/assert"
	"storyly/pkg/utils/string-util"
	"testing"
)

func Test_it_should_check_is_blank(t *testing.T) {
	//Given
	var (
		assert = testifyAssert.New(t)
	)

	//Then
	assert.True(string_util.IsBlank(""))
	assert.False(string_util.IsBlank("1"))
}
