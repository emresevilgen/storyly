package errors_test

import (
	"errors"
	testifyAssert "github.com/stretchr/testify/assert"
	commonErrors "storyly/pkg/errors"
	"testing"
)

func Test_it_should_create_error(t *testing.T) {
	// Given
	var (
		assert = testifyAssert.New(t)
	)

	// When
	err := commonErrors.CreateError(400, "error")

	// Then
	assert.NotNil(err)
	assert.Equal(400, err.StatusCode)
	assert.Equal("error", err.Message)
	assert.NotNil(err.Error())
}

func Test_it_should_create_api_call_failed_error(t *testing.T) {
	// Given
	var (
		assert = testifyAssert.New(t)
	)

	// When
	err := commonErrors.CreateApiCallFailedError(errors.New("error"))

	// Then
	assert.NotNil(err)
	assert.Equal(500, err.StatusCode)
	assert.Equal("error", err.Message)
}
