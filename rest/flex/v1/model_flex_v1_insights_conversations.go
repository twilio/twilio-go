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

// FlexV1InsightsConversations struct for FlexV1InsightsConversations
type FlexV1InsightsConversations struct {
	// The id of the account.
	AccountId *string `json:"account_id,omitempty"`
	// The unique id of the conversation
	ConversationId *string `json:"conversation_id,omitempty"`
	// The count of segments for a conversation
	SegmentCount int `json:"segment_count,omitempty"`
	// The Segments of a conversation
	Segments *[]map[string]interface{} `json:"segments,omitempty"`
}
