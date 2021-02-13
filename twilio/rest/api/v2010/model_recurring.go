/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Recurring the model 'Recurring'
type Recurring string

// List of recurring
const (
	RECURRING_DAILY   Recurring = "daily"
	RECURRING_MONTHLY Recurring = "monthly"
	RECURRING_YEARLY  Recurring = "yearly"
	RECURRING_ALLTIME Recurring = "alltime"
)
