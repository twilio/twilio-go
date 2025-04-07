/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Knowledge
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// KnowledgeV1KnowledgeChunk struct for KnowledgeV1KnowledgeChunk
type KnowledgeV1KnowledgeChunk struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Knowledge resource.
	AccountSid string `json:"account_sid,omitempty"`
	// The chunk content.
	Content string `json:"content,omitempty"`
	// The metadata of the chunk.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The date and time in GMT when the Chunk was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Chunk was updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated time.Time `json:"date_updated,omitempty"`
}
