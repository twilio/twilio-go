package error

import (
	"encoding/json"
	"fmt"

	//"github.com/pkg/errors"
)

// TwilioRestError provides information about an unsuccessful request.
type TwilioRestError struct {
	Code     int					`json:"code"`
	Details  map[string]interface{}	`json:"details"`
	Message  string					`json:"message"`
	MoreInfo string					`json:"more_info"`
	Status   int					`json:"status"`
}



func (err *TwilioRestError) Error() string {
	detailsJson, _ := json.Marshal(err.Details)
	return fmt.Sprintf("Status: %d - ApiError %d: %s (%s) More info: %s",
		err.Status, err.Code, err.Message, detailsJson, err.MoreInfo)
}

func (err *TwilioRestError) Is(target error) bool {
	t, ok := target.(*TwilioRestError)
	if !ok {
		return false
	}
	errDetails, _ := json.Marshal(err.Details)
	targetDetails, _ := json.Marshal(t.Details)
	return (err.Code == t.Code || t.Code == 0) &&
		(string(errDetails) == string(targetDetails) ||
			string(targetDetails) == "") &&
		(err.Message == t.Message || t.Message == "") &&
		(err.MoreInfo == t.MoreInfo || t.MoreInfo == "") &&
		(err.Status == t.Status || t.Status == 0)
}
