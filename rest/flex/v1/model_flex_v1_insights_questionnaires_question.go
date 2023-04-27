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
// FlexV1InsightsQuestionnairesQuestion struct for FlexV1InsightsQuestionnairesQuestion
type FlexV1InsightsQuestionnairesQuestion struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Insights resource and owns this resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The unique ID of the question
	QuestionId *string `json:"question_id,omitempty"`
		// The question.
	Question *string `json:"question,omitempty"`
		// The description for the question.
	Description *string `json:"description,omitempty"`
		// The Category for the question.
	Category *interface{} `json:"category,omitempty"`
		// The answer_set for the question.
	AnswerSetId *string `json:"answer_set_id,omitempty"`
		// The flag  to enable for disable NA for answer.
	AllowNa *bool `json:"allow_na,omitempty"`
		// Integer value that tells a particular question is used by how many questionnaires
	Usage *int `json:"usage,omitempty"`
		// Set of answers for the question
	AnswerSet *interface{} `json:"answer_set,omitempty"`
	Url *string `json:"url,omitempty"`
}


