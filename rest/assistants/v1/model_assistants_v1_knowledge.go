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
	"time"
)

// AssistantsV1Knowledge struct for AssistantsV1Knowledge
type AssistantsV1Knowledge struct {
	// The type of knowledge source.
	Description string `json:"description,omitempty"`
	// The description of knowledge.
	Id string `json:"id"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Knowledge resource.
	AccountSid string `json:"account_sid,omitempty"`
	// The details of the knowledge source based on the type.
	KnowledgeSourceDetails map[string]interface{} `json:"knowledge_source_details,omitempty"`
	// The name of the knowledge source.
	Name string `json:"name"`
	// The status of processing the knowledge source ('QUEUED', 'PROCESSING', 'COMPLETED', 'FAILED')
	Status string `json:"status,omitempty"`
	// The type of knowledge source ('Web', 'Database', 'Text', 'File')
	Type string `json:"type"`
	// The url of the knowledge resource.
	Url string `json:"url,omitempty"`
	// The embedding model to be used for the knowledge source.
	EmbeddingModel string `json:"embedding_model,omitempty"`
	// The date and time in GMT when the Knowledge was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated time.Time `json:"date_created"`
	// The date and time in GMT when the Knowledge was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated time.Time `json:"date_updated"`
}
