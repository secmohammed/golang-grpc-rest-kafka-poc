package utils

import "github.com/siruspen/logrus"

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ValidationErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func CreateApiError(status int, err error) (int, *ApiError) {
	logrus.Info(err.Error())
	message := err.Error()
	return status, &ApiError{
		Status:  status,
		Message: message,
	}
}
