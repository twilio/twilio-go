/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ServiceGeoMatchLevel the model 'ServiceGeoMatchLevel'
type ServiceGeoMatchLevel string

// List of service_geo_match_level
const (
	SERVICEGEOMATCHLEVEL_AREA_CODE ServiceGeoMatchLevel = "area-code"
	SERVICEGEOMATCHLEVEL_OVERLAY   ServiceGeoMatchLevel = "overlay"
	SERVICEGEOMATCHLEVEL_RADIUS    ServiceGeoMatchLevel = "radius"
	SERVICEGEOMATCHLEVEL_COUNTRY   ServiceGeoMatchLevel = "country"
)
