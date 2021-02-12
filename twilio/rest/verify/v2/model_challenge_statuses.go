/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ChallengeStatuses the model 'ChallengeStatuses'
type ChallengeStatuses string

// List of challenge_statuses
const (
	CHALLENGESTATUSES_PENDING ChallengeStatuses = "pending"
	CHALLENGESTATUSES_EXPIRED ChallengeStatuses = "expired"
	CHALLENGESTATUSES_APPROVED ChallengeStatuses = "approved"
	CHALLENGESTATUSES_DENIED ChallengeStatuses = "denied"
)
