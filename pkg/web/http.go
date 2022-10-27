package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// ResponseFormatter returning formatted JSON response
func ResponseFormatter(ctx echo.Context, code int, message string, body interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    nil,
			"error":   err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
		}
	}

	return ctx.JSON(code, response)
}

// ResponseErrValidation returning formatted JSON response
func ResponseErrValidation(ctx echo.Context, message string, errMap map[string]interface{}) error {

	var b strings.Builder
	for k, v := range errMap {
		b.WriteString(fmt.Sprintf("%s : %v, ", k, v))
	}
	errorString := strings.TrimRight(b.String(), ", ")

	response := map[string]interface{}{
		"message":          message,
		"data":             nil,
		"error_validation": errMap,
		"error":            errorString,
	}

	return ctx.JSON(http.StatusBadRequest, response)
}

// ResponseErrWithFormValidation returning formatted JSON response
func ResponseErrWithFormatValidation(ctx echo.Context, message string, validation map[string]interface{}) error {
	return ctx.JSON(
		http.StatusBadRequest,
		struct {
			Message        string                 `json:"message"`
			FormValidation map[string]interface{} `json:"form_validation"`
		}{Message: message, FormValidation: validation},
	)
}

func ResponseErrValidationWithCode(ctx echo.Context, message string, errMap map[string]interface{}, code int) error {

	var msg string

	if len(errMap) == 0 {
		msg = message
	} else {
		for _, value := range errMap {
			msg = fmt.Sprintf("%s", value)
			break
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"data":             nil,
		"error_validation": errMap,
	}

	return ctx.JSON(code, response)
}

// ResponseFormatter returning formatted JSON response with meta
func ResponseFormatterWithMeta(ctx echo.Context, code int, message string, body interface{}, meta interface{}, err error) error {
	var response map[string]interface{}

	if err != nil {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   err.Error(),
			"meta":    meta,
		}
	} else {
		response = map[string]interface{}{
			"message": message,
			"data":    body,
			"error":   nil,
			"meta":    meta,
		}
	}

	return ctx.JSON(code, response)
}

func ResponseErrValidationWithDefaultMessage(ctx echo.Context, message string, errMap map[string]interface{}, code int) error {
	var msg string

	if len(errMap) == 0 {
		msg = message
	} else {
		for _, value := range errMap {
			msg = value.(string)
			break
		}
	}

	response := map[string]interface{}{
		"message":          msg,
		"error_validation": errMap,
	}

	return ctx.JSON(code, response)
}

func ResponseFormaterWithDefaultEmptyObject(ctx echo.Context, message, errMessage string, data, meta interface{}, errMap map[string]interface{}, code int) error {

	var status bool
	if data == nil || len(errMap) == 0 {
		status = true
	}

	if len(errMap) == 0 {
		errMap = map[string]interface{}{}
	} else {
		var b strings.Builder
		for k, v := range errMap {
			b.WriteString(fmt.Sprintf("%s : %v, ", k, v))
		}
		errMessage = strings.TrimRight(b.String(), ", ")
	}

	// make empty object instead null value
	if data == nil {
		data = make(map[string]string, 0)
	}

	response := map[string]interface{}{
		"status":           status,
		"msg":              message,
		"message":          message,
		"data":             data,
		"error_validation": errMap,
		"error":            errMessage,
	}

	if meta != nil {
		response["meta"] = meta
	}

	return ctx.JSON(code, response)
}
