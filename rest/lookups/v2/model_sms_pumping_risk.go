/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Lookups
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// SmsPumpingRisk struct for SmsPumpingRisk
type SmsPumpingRisk struct {
	CarrierRiskCategory      string    `json:"carrier_risk_category,omitempty"`
	NumberBlocked            bool      `json:"number_blocked,omitempty"`
	NumberBlockedDate        time.Time `json:"number_blocked_date,omitempty"`
	NumberBlockedLast3Months bool      `json:"number_blocked_last_3_months,omitempty"`
	SmsPumpingRiskScore      int       `json:"sms_pumping_risk_score,omitempty"`
	ErrorCode                int       `json:"error_code,omitempty"`
}
