/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Lookup
 * Lookup APIs for individual, bulk and job based requests  Discussion topics: - API version to use - Using or not lookup in the path or just as lookups subdomain
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ReassignedNumberResponse struct for ReassignedNumberResponse
type ReassignedNumberResponse struct {
	LastVerifiedDate   string `json:"last_verified_date,omitempty"`
	IsNumberReassigned string `json:"is_number_reassigned,omitempty"`
	ErrorCode          string `json:"error_code,omitempty"`
}