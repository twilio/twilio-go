/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// ApiV2010UsageRecordMonthly struct for ApiV2010UsageRecordMonthly
type ApiV2010UsageRecordMonthly struct {
	// The SID of the Account accrued the usage
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the resource
	ApiVersion *string `json:"api_version,omitempty"`
	// Usage records up to date as of this timestamp
	AsOf *string `json:"as_of,omitempty"`
	// The category of usage
	Category *string `json:"category,omitempty"`
	// The number of usage events
	Count *string `json:"count,omitempty"`
	// The units in which count is measured
	CountUnit *string `json:"count_unit,omitempty"`
	// A plain-language description of the usage category
	Description *string `json:"description,omitempty"`
	// The last date for which usage is included in the UsageRecord
	EndDate *string `json:"end_date,omitempty"`
	// The total price of the usage
	Price *float32 `json:"price,omitempty"`
	// The currency in which `price` is measured
	PriceUnit *string `json:"price_unit,omitempty"`
	// The first date for which usage is included in this UsageRecord
	StartDate *string `json:"start_date,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// The amount of usage
	Usage *string `json:"usage,omitempty"`
	// The units in which usage is measured
	UsageUnit *string `json:"usage_unit,omitempty"`
}

func (response *ApiV2010UsageRecordMonthly) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid      *string                 `json:"account_sid"`
		ApiVersion      *string                 `json:"api_version"`
		AsOf            *string                 `json:"as_of"`
		Category        *string                 `json:"category"`
		Count           *string                 `json:"count"`
		CountUnit       *string                 `json:"count_unit"`
		Description     *string                 `json:"description"`
		EndDate         *string                 `json:"end_date"`
		Price           *interface{}            `json:"price"`
		PriceUnit       *string                 `json:"price_unit"`
		StartDate       *string                 `json:"start_date"`
		SubresourceUris *map[string]interface{} `json:"subresource_uris"`
		Uri             *string                 `json:"uri"`
		Usage           *string                 `json:"usage"`
		UsageUnit       *string                 `json:"usage_unit"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = ApiV2010UsageRecordMonthly{
		AccountSid:      raw.AccountSid,
		ApiVersion:      raw.ApiVersion,
		AsOf:            raw.AsOf,
		Category:        raw.Category,
		Count:           raw.Count,
		CountUnit:       raw.CountUnit,
		Description:     raw.Description,
		EndDate:         raw.EndDate,
		PriceUnit:       raw.PriceUnit,
		StartDate:       raw.StartDate,
		SubresourceUris: raw.SubresourceUris,
		Uri:             raw.Uri,
		Usage:           raw.Usage,
		UsageUnit:       raw.UsageUnit,
	}

	responsePrice, err := client.UnmarshalFloat32(raw.Price)
	if err != nil {
		return err
	}
	response.Price = responsePrice

	return
}
