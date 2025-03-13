/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// ApiV2010UsageRecordLastMonth struct for ApiV2010UsageRecordLastMonth
type ApiV2010UsageRecordLastMonth struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that accrued the usage.
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Usage records up to date as of this timestamp, formatted as YYYY-MM-DDTHH:MM:SS+00:00. All timestamps are in GMT
	AsOf     *string `json:"as_of,omitempty"`
	Category *string `json:"category,omitempty"`
	// The number of usage events, such as the number of calls.
	Count *string `json:"count,omitempty"`
	// The units in which `count` is measured, such as `calls` for calls or `messages` for SMS.
	CountUnit *string `json:"count_unit,omitempty"`
	// A plain-language description of the usage category.
	Description *string `json:"description,omitempty"`
	// The last date for which usage is included in the UsageRecord. The date is specified in GMT and formatted as `YYYY-MM-DD`.
	EndDate *string `json:"end_date,omitempty"`
	// The total price of the usage in the currency specified in `price_unit` and associated with the account.
	Price *float32 `json:"price,omitempty"`
	// The currency in which `price` is measured, in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format, such as `usd`, `eur`, and `jpy`.
	PriceUnit *string `json:"price_unit,omitempty"`
	// The first date for which usage is included in this UsageRecord. The date is specified in GMT and formatted as `YYYY-MM-DD`.
	StartDate *string `json:"start_date,omitempty"`
	// A list of related resources identified by their URIs. For more information, see [List Subresources](https://www.twilio.com/docs/usage/api/usage-record#list-subresources).
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
	// The amount used to bill usage and measured in units described in `usage_unit`.
	Usage *string `json:"usage,omitempty"`
	// The units in which `usage` is measured, such as `minutes` for calls or `messages` for SMS.
	UsageUnit *string `json:"usage_unit,omitempty"`
}

func (response *ApiV2010UsageRecordLastMonth) UnmarshalJSON(bytes []byte) (err error) {
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

	*response = ApiV2010UsageRecordLastMonth{
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
