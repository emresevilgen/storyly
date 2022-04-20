package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	notFoundMessage = "Record not found!"
)

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func CreateError(statusCode int, message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

func CreateDbError(err error) ErrorResponse {
	return CreateError(http.StatusInternalServerError, err.Error())
}

func CreateNotFoundError() ErrorResponse {
	return CreateError(http.StatusNotFound, notFoundMessage)
}

func CreateApiCallFailedError(err error) ErrorResponse {
	return CreateError(http.StatusInternalServerError, err.Error())
}

func CreateAuthenticationError() ErrorResponse {
	return CreateError(http.StatusUnauthorized, "Unauthorized credentials")
}

func (e ErrorResponse) Error() string {
	errorResponse, _ := json.Marshal(e)
	return fmt.Sprintf("%s", errorResponse)
}
