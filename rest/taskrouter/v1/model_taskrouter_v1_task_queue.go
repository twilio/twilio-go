/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

// TaskrouterV1TaskQueue struct for TaskrouterV1TaskQueue
type TaskrouterV1TaskQueue struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Activity to assign Workers when a task is assigned for them.
	AssignmentActivitySid *string `json:"assignment_activity_sid,omitempty"`
	// The name of the Activity to assign Workers when a task is assigned for them.
	AssignmentActivityName *string `json:"assignment_activity_name,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The maximum number of Workers to reserve for the assignment of a task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
	MaxReservedWorkers *int `json:"max_reserved_workers,omitempty"`
	// The SID of the Activity to assign Workers once a task is reserved for them.
	ReservationActivitySid *string `json:"reservation_activity_sid,omitempty"`
	// The name of the Activity to assign Workers once a task is reserved for them.
	ReservationActivityName *string `json:"reservation_activity_name,omitempty"`
	// The unique string that we created to identify the TaskQueue resource.
	Sid *string `json:"sid,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example `'\"language\" == \"spanish\"'` If no TargetWorkers parameter is provided, Tasks will wait in the TaskQueue until they are either deleted or moved to another TaskQueue. Additional examples on how to describing Worker selection criteria below. Defaults to 1==1.
	TargetWorkers *string `json:"target_workers,omitempty"`
	TaskOrder     *string `json:"task_order,omitempty"`
	// The absolute URL of the TaskQueue resource.
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the TaskQueue.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the Operating Unit that contains the TaskQueue.
	OperatingUnitSid *string `json:"operating_unit_sid,omitempty"`
}
