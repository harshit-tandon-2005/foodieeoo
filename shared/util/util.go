package util

import (
	"github.com/foodieeoo/models"
	"github.com/labstack/echo"
)

func ApiResponse(c echo.Context, status string, data interface{}, message string, httpCode int, meta *models.ResponsePatternMeta) error {

	resp := &models.ApiResponsePattern{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    httpCode,
	}

	if meta != nil {
		resp.Meta = meta
	}

	return c.JSON(httpCode, resp)
}

func SetUsecaseResponse(data interface{}, err error, statusCode int, errorCode, message string) models.ApiUsescaseResponse {
	return models.ApiUsescaseResponse{
		Data:       data,
		Error:      err,
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
	}
}
