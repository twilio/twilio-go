/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Assistants
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

// AssistantsV1AssistantSendMessageRequest struct for AssistantsV1AssistantSendMessageRequest
type AssistantsV1AssistantSendMessageRequest struct {
	// The unique identity of user for the session.
	Identity string `json:"identity"`
	// The unique name for the session.
	SessionId string `json:"session_id,omitempty"`
	// The query to ask the assistant.
	Body string `json:"body"`
	// The webhook url to call after the assistant has generated a response or report an error.
	Webhook string `json:"webhook,omitempty"`
	// one of the modes 'chat', 'email' or 'voice'
	Mode string `json:"mode,omitempty"`
}
