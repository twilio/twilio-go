/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskQueueTaskOrder the model 'TaskQueueTaskOrder'
type TaskQueueTaskOrder string

// List of task_queue_task_order
const (
	TASKQUEUETASKORDER_FIFO TaskQueueTaskOrder = "FIFO"
	TASKQUEUETASKORDER_LIFO TaskQueueTaskOrder = "LIFO"
)
