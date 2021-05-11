/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ChallengeFactorTypes the model 'ChallengeFactorTypes'
type ChallengeFactorTypes string

// List of challenge_factor_types
const (
	CHALLENGEFACTORTYPES_PUSH ChallengeFactorTypes = "push"
	CHALLENGEFACTORTYPES_TOTP ChallengeFactorTypes = "totp"
)
