package client_test

import (
	"encoding/json"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/twilio/twilio-go/client"
)

const (
	errorCode     = 20001
	errorMessage  = "Bad request"
	errorMoreInfo = "https://www.twilio.com/docs/errors/20001"
	errorStatus   = 400
)

func TestTwilioRestError_Error(t *testing.T) {
	details := make(map[string]interface{})
	details["foo"] = "bar"
	err := &client.TwilioRestError{
		Code:     errorCode,
		Details:  details,
		Message:  errorMessage,
		MoreInfo: errorMoreInfo,
		Status:   errorStatus,
	}
	expected := "Status: 400 - ApiError 20001: Bad request ({\"foo\":\"bar\"}) More info: https://www.twilio.com/docs/errors/20001"
	assert.Equal(t, expected, err.Error())
}

func TestTwilioRestError_NoDetails(t *testing.T) {
	err := &client.TwilioRestError{}
	response := `{"code":20001,"message":"Bad request","more_info":"https://www.twilio.com/docs/errors/20001","status":400}`
	responseReader := strings.NewReader(response)
	decodeErr := json.NewDecoder(responseReader).Decode(err)
	assert.Nil(t, decodeErr)
	assert.Equal(t, err.Code, errorCode)
	assert.Equal(t, err.Status, errorStatus)
	assert.Equal(t, err.Message, errorMessage)
	assert.Equal(t, err.MoreInfo, errorMoreInfo)
	assert.Nil(t, err.Details)
}

// Tests for RestErrorV1 (Twilio API Standards V1.0)

func TestRestErrorV1_Error(t *testing.T) {
	params := make(map[string]interface{})
	params["twilioErrorCodeUrl"] = "https://www.twilio.com/docs/errors/20404"
	err := &client.RestErrorV1{
		Code:           20404,
		Message:        "The requested resource was not found",
		HttpStatusCode: 404,
		Params:         params,
		UserError:      true,
	}
	errorStr := err.Error()
	assert.Contains(t, errorStr, "Status: 404")
	assert.Contains(t, errorStr, "ApiError 20404")
	assert.Contains(t, errorStr, "The requested resource was not found")
	assert.Contains(t, errorStr, "twilioErrorCodeUrl")
	assert.Contains(t, errorStr, "userError: true")
}

func TestRestErrorV1_GetCode(t *testing.T) {
	err := &client.RestErrorV1{
		Code:           20404,
		Message:        "Not found",
		HttpStatusCode: 404,
	}
	assert.Equal(t, 20404, err.GetCode())
}

func TestRestErrorV1_GetMessage(t *testing.T) {
	err := &client.RestErrorV1{
		Code:           20404,
		Message:        "The requested resource was not found",
		HttpStatusCode: 404,
	}
	assert.Equal(t, "The requested resource was not found", err.GetMessage())
}

func TestRestErrorV1_GetHttpStatusCode(t *testing.T) {
	err := &client.RestErrorV1{
		Code:           20404,
		Message:        "Not found",
		HttpStatusCode: 404,
	}
	assert.Equal(t, 404, err.GetHttpStatusCode())
}

func TestRestErrorV1_GetParams(t *testing.T) {
	params := make(map[string]interface{})
	params["twilioErrorCodeUrl"] = "https://www.twilio.com/docs/errors/20404"
	params["resourceId"] = "RM123"
	err := &client.RestErrorV1{
		Code:           20404,
		Message:        "Not found",
		HttpStatusCode: 404,
		Params:         params,
	}
	resultParams := err.GetParams()
	assert.Equal(t, 2, len(resultParams))
	assert.Equal(t, "https://www.twilio.com/docs/errors/20404", resultParams["twilioErrorCodeUrl"])
	assert.Equal(t, "RM123", resultParams["resourceId"])
}

func TestRestErrorV1_GetUserError(t *testing.T) {
	err := &client.RestErrorV1{
		Code:           20404,
		Message:        "Not found",
		HttpStatusCode: 404,
		UserError:      true,
	}
	assert.True(t, err.GetUserError())

	err2 := &client.RestErrorV1{
		Code:           50001,
		Message:        "Internal server error",
		HttpStatusCode: 500,
		UserError:      false,
	}
	assert.False(t, err2.GetUserError())
}

func TestRestErrorV1_FromJSON(t *testing.T) {
	err := &client.RestErrorV1{}
	response := `{
		"code": 20404,
		"message": "The requested resource was not found",
		"httpStatusCode": 404,
		"userError": true,
		"params": {
			"twilioErrorCodeUrl": "https://www.twilio.com/docs/errors/20404"
		}
	}`
	responseReader := strings.NewReader(response)
	decodeErr := json.NewDecoder(responseReader).Decode(err)
	assert.Nil(t, decodeErr)
	assert.Equal(t, 20404, err.Code)
	assert.Equal(t, "The requested resource was not found", err.Message)
	assert.Equal(t, 404, err.HttpStatusCode)
	assert.True(t, err.UserError)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20404", err.Params["twilioErrorCodeUrl"])
}

func TestRestErrorV1_WithoutParams(t *testing.T) {
	err := &client.RestErrorV1{
		Code:           400,
		Message:        "Bad request",
		HttpStatusCode: 400,
		Params:         nil,
		UserError:      false,
	}
	errorStr := err.Error()
	assert.Contains(t, errorStr, "Status: 400")
	assert.Contains(t, errorStr, "ApiError 400")
	assert.Contains(t, errorStr, "Bad request")
	assert.Contains(t, errorStr, "params: null")
	assert.Contains(t, errorStr, "userError: false")

	// Test getters
	assert.Equal(t, 400, err.GetCode())
	assert.Equal(t, "Bad request", err.GetMessage())
	assert.Equal(t, 400, err.GetHttpStatusCode())
	assert.Nil(t, err.GetParams())
	assert.False(t, err.GetUserError())
}

func TestRestErrorV1_EmptyParams(t *testing.T) {
	err := &client.RestErrorV1{
		Code:           400,
		Message:        "Bad request",
		HttpStatusCode: 400,
		Params:         make(map[string]interface{}),
		UserError:      false,
	}
	resultParams := err.GetParams()
	assert.NotNil(t, resultParams)
	assert.Equal(t, 0, len(resultParams))
}
