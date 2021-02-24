/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SimStatus the model 'SimStatus'
type SimStatus string

// List of sim_status
const (
	SIMSTATUS_NEW         SimStatus = "new"
	SIMSTATUS_READY       SimStatus = "ready"
	SIMSTATUS_ACTIVE      SimStatus = "active"
	SIMSTATUS_SUSPENDED   SimStatus = "suspended"
	SIMSTATUS_DEACTIVATED SimStatus = "deactivated"
	SIMSTATUS_CANCELED    SimStatus = "canceled"
	SIMSTATUS_SCHEDULED   SimStatus = "scheduled"
	SIMSTATUS_UPDATING    SimStatus = "updating"
)
