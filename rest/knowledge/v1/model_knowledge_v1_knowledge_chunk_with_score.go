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
	"encoding/json"
	"time"

	"github.com/twilio/twilio-go/client"
)

// KnowledgeV1KnowledgeChunkWithScore struct for KnowledgeV1KnowledgeChunkWithScore
type KnowledgeV1KnowledgeChunkWithScore struct {
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
	// The score associated with the chunk.
	Score float32 `json:"score,omitempty"`
	// The knowledge ID associated with the chunk.
	KnowledgeId string `json:"knowledge_id,omitempty"`
}

func (response *KnowledgeV1KnowledgeChunkWithScore) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid  string                 `json:"account_sid"`
		Content     string                 `json:"content"`
		Metadata    map[string]interface{} `json:"metadata"`
		DateCreated time.Time              `json:"date_created"`
		DateUpdated time.Time              `json:"date_updated"`
		Score       interface{}            `json:"score"`
		KnowledgeId string                 `json:"knowledge_id"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = KnowledgeV1KnowledgeChunkWithScore{
		AccountSid:  raw.AccountSid,
		Content:     raw.Content,
		Metadata:    raw.Metadata,
		DateCreated: raw.DateCreated,
		DateUpdated: raw.DateUpdated,
		KnowledgeId: raw.KnowledgeId,
	}

	responseScore, err := client.UnmarshalFloat32(&raw.Score)
	if err != nil {
		return err
	}
	response.Score = *responseScore

	return
}
