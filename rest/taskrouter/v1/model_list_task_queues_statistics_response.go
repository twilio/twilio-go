/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTaskQueuesStatisticsResponse struct for ListTaskQueuesStatisticsResponse
type ListTaskQueuesStatisticsResponse struct {
	Meta                 ListWorkspaceResponseMeta          `json:"meta,omitempty"`
	TaskQueuesStatistics []TaskrouterV1TaskQueuesStatistics `json:"task_queues_statistics,omitempty"`
}
