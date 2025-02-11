package util

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ValidationErr map[string]string
}

type ValidationErrorStruct struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ValidationErr map[string]string
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFoundErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusNotFound, Message: msg[0]}
	}
	return &AppError{Code: http.StatusNotFound, Message: "Not found"}
}

func BadRequestErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusBadRequest, Message: msg[0]}
	}
	return &AppError{Code: http.StatusBadRequest, Message: "Bad request"}
}

func UnauthorizedErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusUnauthorized, Message: msg[0]}
	}
	return &AppError{Code: http.StatusBadRequest, Message: "Unauthorized"}
}

func InternalServerErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusInternalServerError, Message: msg[0]}
	}
	return &AppError{Code: http.StatusInternalServerError, Message: "Internal server error"}
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			switch e := err.(type) {
			case *AppError:
				c.AbortWithStatusJSON(e.Code, gin.H{
					"StatusCode": e.Code,
					"Error":      e.Message,
					"Messages":   e.ValidationErr,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"StatusCode": http.StatusInternalServerError,
					"Error":      e.Error(),
				})
			}
		}
	}
}

func ValidationErr(err error, obj interface{}) *AppError {
	errorMessages := make(map[string]string)

	// Type assertion to get validation errors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		objType := reflect.TypeOf(obj)
		for _, fieldErr := range validationErrors {

			tag := fieldErr.Tag()
			structFieldName := fieldErr.Field()

			var jsonTag string
			if objField, found := objType.Elem().FieldByName(structFieldName); found {
				jsonTag = objField.Tag.Get("json")
			} else {
				jsonTag = structFieldName
			}

			switch tag {
			case "required":
				errorMessages[jsonTag] = "This field is required."
			case "email":
				errorMessages[jsonTag] = "Invalid email format."
			case "url":
				errorMessages[jsonTag] = "Invalid URL format."
			case "min":
				errorMessages[jsonTag] = "Value must be at least " + fieldErr.Param() + "."
			case "max":
				errorMessages[jsonTag] = "Value must not exceed " + fieldErr.Param() + "."
			case "len":
				errorMessages[jsonTag] = "Value must be exactly " + fieldErr.Param() + " characters."
			case "oneof":
				errorMessages[jsonTag] = "Value must be one of the following: " + fieldErr.Param() + "."
			case "gt":
				errorMessages[jsonTag] = "Value must be greater than " + fieldErr.Param() + "."
			case "gte":
				errorMessages[jsonTag] = "Value must be greater than or equal to " + fieldErr.Param() + "."
			case "lt":
				errorMessages[jsonTag] = "Value must be less than " + fieldErr.Param() + "."
			case "lte":
				errorMessages[jsonTag] = "Value must be less than or equal to " + fieldErr.Param() + "."
			case "alpha":
				errorMessages[jsonTag] = "Value must contain only letters."
			case "alphanum":
				errorMessages[jsonTag] = "Value must contain only letters and numbers."
			case "numeric":
				errorMessages[jsonTag] = "Value must be a valid number."
			case "uuid":
				errorMessages[jsonTag] = "Value must be a valid UUID."
			case "ipv4":
				errorMessages[jsonTag] = "Value must be a valid IPv4 address."
			case "ipv6":
				errorMessages[jsonTag] = "Value must be a valid IPv6 address."
			case "ip":
				errorMessages[jsonTag] = "Value must be a valid IP address."
			case "contains":
				errorMessages[jsonTag] = "Value must contain " + fieldErr.Param() + "."
			case "excludes":
				errorMessages[jsonTag] = "Value must not contain " + fieldErr.Param() + "."
			case "startswith":
				errorMessages[jsonTag] = "Value must start with " + fieldErr.Param() + "."
			case "endswith":
				errorMessages[jsonTag] = "Value must end with " + fieldErr.Param() + "."
			case "boolean":
				errorMessages[jsonTag] = "Value must be true or false."
			case "datetime":
				errorMessages[jsonTag] = "Value must be a valid datetime in the format: " + fieldErr.Param() + "."
			case "base64":
				errorMessages[jsonTag] = "Value must be a valid base64 string."
			case "hexadecimal":
				errorMessages[jsonTag] = "Value must be a valid hexadecimal string."
			case "json":
				errorMessages[jsonTag] = "Value must be a valid JSON string."
			case "required_if":
				errorMessages[jsonTag] = "This field is required when " + fieldErr.Param() + " is present."
			case "required_unless":
				errorMessages[jsonTag] = "This field is required unless " + fieldErr.Param() + " is present."
			case "required_with":
				errorMessages[jsonTag] = "This field is required when " + fieldErr.Param() + " is present."
			case "required_with_all":
				errorMessages[jsonTag] = "This field is required when all of " + fieldErr.Param() + " are present."
			case "required_without":
				errorMessages[jsonTag] = "This field is required when " + fieldErr.Param() + " is not present."
			case "required_without_all":
				errorMessages[jsonTag] = "This field is required when none of " + fieldErr.Param() + " are present."
			case "unique":
				errorMessages[jsonTag] = "This value must be unique."
			default:
				errorMessages[jsonTag] = "Invalid value."
			}
		}
	}

	return &AppError{Code: http.StatusBadRequest, Message: "Validation error", ValidationErr: errorMessages}
}
