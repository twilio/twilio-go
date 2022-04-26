/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// ApiV2010RecordingTranscription struct for ApiV2010RecordingTranscription
type ApiV2010RecordingTranscription struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the transcription
	ApiVersion *string `json:"api_version,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The duration of the transcribed audio in seconds.
	Duration *string `json:"duration,omitempty"`
	// The charge for the transcription
	Price *float32 `json:"price,omitempty"`
	// The currency in which price is measured
	PriceUnit *string `json:"price_unit,omitempty"`
	// The SID that identifies the transcription's recording
	RecordingSid *string `json:"recording_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the transcription
	Status *string `json:"status,omitempty"`
	// The text content of the transcription.
	TranscriptionText *string `json:"transcription_text,omitempty"`
	// The transcription type
	Type *string `json:"type,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}

func (response *ApiV2010RecordingTranscription) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid        *string      `json:"account_sid"`
		ApiVersion        *string      `json:"api_version"`
		DateCreated       *string      `json:"date_created"`
		DateUpdated       *string      `json:"date_updated"`
		Duration          *string      `json:"duration"`
		Price             *interface{} `json:"price"`
		PriceUnit         *string      `json:"price_unit"`
		RecordingSid      *string      `json:"recording_sid"`
		Sid               *string      `json:"sid"`
		Status            *string      `json:"status"`
		TranscriptionText *string      `json:"transcription_text"`
		Type              *string      `json:"type"`
		Uri               *string      `json:"uri"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = ApiV2010RecordingTranscription{
		AccountSid:        raw.AccountSid,
		ApiVersion:        raw.ApiVersion,
		DateCreated:       raw.DateCreated,
		DateUpdated:       raw.DateUpdated,
		Duration:          raw.Duration,
		PriceUnit:         raw.PriceUnit,
		RecordingSid:      raw.RecordingSid,
		Sid:               raw.Sid,
		Status:            raw.Status,
		TranscriptionText: raw.TranscriptionText,
		Type:              raw.Type,
		Uri:               raw.Uri,
	}

	responsePrice, err := client.UnmarshalFloat32(raw.Price)
	if err != nil {
		return err
	}
	response.Price = responsePrice

	return
}
