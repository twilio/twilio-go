/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// CommandStatus the model 'CommandStatus'
type CommandStatus string

// List of command_status
const (
	COMMANDSTATUS_QUEUED    CommandStatus = "queued"
	COMMANDSTATUS_SENT      CommandStatus = "sent"
	COMMANDSTATUS_DELIVERED CommandStatus = "delivered"
	COMMANDSTATUS_RECEIVED  CommandStatus = "received"
	COMMANDSTATUS_FAILED    CommandStatus = "failed"
)
