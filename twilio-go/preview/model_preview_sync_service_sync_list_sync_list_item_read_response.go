/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PreviewSyncServiceSyncListSyncListItemReadResponse struct for PreviewSyncServiceSyncListSyncListItemReadResponse
type PreviewSyncServiceSyncListSyncListItemReadResponse struct {
	Items []PreviewSyncServiceSyncListSyncListItem `json:"items,omitempty"`
	Meta PreviewBulkExportsExportDayReadResponseMeta `json:"meta,omitempty"`
}
