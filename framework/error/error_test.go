package error_test

import (
	"encoding/json"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/twilio/twilio-go/framework/error"
)

const errorCode = 20001
const errorMessage = "Bad request"
const errorMoreInfo = "https://www.twilio.com/docs/errors/20001"
const errorStatus = 400

func TestTwilioRestError_Error(t *testing.T) {
	details := make(map[string]interface{})
	details["foo"] = "bar"
	err := &error.TwilioRestError{
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
	err := &error.TwilioRestError{}
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
