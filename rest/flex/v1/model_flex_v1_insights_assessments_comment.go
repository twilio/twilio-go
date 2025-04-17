/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// FlexV1InsightsAssessmentsComment struct for FlexV1InsightsAssessmentsComment
type FlexV1InsightsAssessmentsComment struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Insights resource and owns this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the assessment.
	AssessmentSid *string `json:"assessment_sid,omitempty"`
	// The comment added for assessment.
	Comment *map[string]interface{} `json:"comment,omitempty"`
	// The offset
	Offset *float32 `json:"offset,omitempty"`
	// The flag indicating if this assessment is part of report
	Report *bool `json:"report,omitempty"`
	// The weightage given to this comment
	Weight *float32 `json:"weight,omitempty"`
	// The id of the agent.
	AgentId *string `json:"agent_id,omitempty"`
	// The id of the segment.
	SegmentId *string `json:"segment_id,omitempty"`
	// The name of the user.
	UserName *string `json:"user_name,omitempty"`
	// The email id of the user.
	UserEmail *string `json:"user_email,omitempty"`
	// The timestamp when the record is inserted
	Timestamp *float32 `json:"timestamp,omitempty"`
	Url       *string  `json:"url,omitempty"`
}

func (response *FlexV1InsightsAssessmentsComment) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid    *string                 `json:"account_sid"`
		AssessmentSid *string                 `json:"assessment_sid"`
		Comment       *map[string]interface{} `json:"comment"`
		Offset        *interface{}            `json:"offset"`
		Report        *bool                   `json:"report"`
		Weight        *interface{}            `json:"weight"`
		AgentId       *string                 `json:"agent_id"`
		SegmentId     *string                 `json:"segment_id"`
		UserName      *string                 `json:"user_name"`
		UserEmail     *string                 `json:"user_email"`
		Timestamp     *interface{}            `json:"timestamp"`
		Url           *string                 `json:"url"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = FlexV1InsightsAssessmentsComment{
		AccountSid:    raw.AccountSid,
		AssessmentSid: raw.AssessmentSid,
		Comment:       raw.Comment,
		Report:        raw.Report,
		AgentId:       raw.AgentId,
		SegmentId:     raw.SegmentId,
		UserName:      raw.UserName,
		UserEmail:     raw.UserEmail,
		Url:           raw.Url,
	}

	responseOffset, err := client.UnmarshalFloat32(raw.Offset)
	if err != nil {
		return err
	}
	response.Offset = responseOffset

	responseWeight, err := client.UnmarshalFloat32(raw.Weight)
	if err != nil {
		return err
	}
	response.Weight = responseWeight

	responseTimestamp, err := client.UnmarshalFloat32(raw.Timestamp)
	if err != nil {
		return err
	}
	response.Timestamp = responseTimestamp

	return
}
