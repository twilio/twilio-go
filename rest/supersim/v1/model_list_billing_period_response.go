/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListBillingPeriodResponse struct for ListBillingPeriodResponse
type ListBillingPeriodResponse struct {
	BillingPeriods []SupersimV1BillingPeriod   `json:"billing_periods,omitempty"`
	Meta           ListEsimProfileResponseMeta `json:"meta,omitempty"`
}
