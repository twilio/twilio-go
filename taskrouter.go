// Package twilio - base package
package twilio

import (
	"time"
)

//Workspace refer: https://www.twilio.com/docs/taskrouter/api/workspace
type Workspace struct {
	Sid                  string            `json:"sid"`
	AccountSid           string            `json:"account_sid"`
	DateCreated          time.Time         `json:"date_created"`
	DateUpdated          time.Time         `json:"date_updated"`
	DefaultActivityName  string            `json:"default_activity_name"`
	DefaultActivitySid   string            `json:"default_activity_sid"`
	TimeoutActivityName  string            `json:"timeout_activity_name"`
	TimeoutActivitySid   string            `json:"timeout_activity_sid"`
	URL                  string            `json:"url"`
	Links                map[string]string `json:"links"`
	FriendlyName         string            `json:"friendly_name"`
	EventCallbackURL     string            `json:"event_callback_url"`
	EventsFilter         string            `json:"events_filter"`
	MultitaskEnabled     bool              `json:"multi_task_enabled"`
	Template             string            `json:"template"`
	PrioritizeQueueOrder string            `json:"prioritize_queue_order"`
}

//WorkspaceParams refer: https://www.twilio.com/docs/taskrouter/api/workspace
type WorkspaceParams struct {
	FriendlyName         string `url:"FriendlyName,omitempty"`
	EventCallbackURL     string `url:"EventCallbackUrl,omitempty"`
	EventsFilter         string `url:"EventsFilter,omitempty"`
	MultitaskEnabled     bool   `url:"MultitaskEnabled,omitempty"`
	Template             string `url:"Template,omitempty"`
	PrioritizeQueueOrder string `url:"PrioritizeQueueOrder,omitempty"`
}

// constants
const (
	//WorkspaceDefaultFriendlyName = "Flex Task Assignment"
	WorkspaceDefaultFriendlyName = "Flex Task Assignment"
	WorkflowDefaultFriendlyName  = "Assign to Anyone"
	TaskQueueDefaultFriendlyName = "Everyone"
	ActivityDefaultFriendlyName  = "Break"
)

// Workflow workflow refer: https://www.twilio.com/docs/taskrouter/api/workflow
type Workflow struct {
	AssignmentCallbackURL         string            `json:"assignment_callback_url"`
	Configuration                 string            `json:"configuration"`
	AccountSid                    string            `json:"account_sid"`
	DateCreated                   time.Time         `json:"date_created"`
	DateUpdated                   time.Time         `json:"date_updated"`
	DocumentContentType           string            `json:"document_content_type"`
	FallbackAssignmentCallbackURL string            `json:"fallback_assignment_callback_url"`
	FriendlyName                  string            `json:"friendly_name"`
	Sid                           string            `json:"sid"`
	TaskReservationTimeout        int               `json:"task_reservation_timeout"`
	WorkspaceSid                  string            `json:"workspace_sid"`
	URL                           string            `json:"url"`
	Links                         map[string]string `json:"links"`
}

// WorkflowParams workflow parameters
type WorkflowParams struct {
	WorkspaceSid                  string `url:"WorkspaceSid,omitempty"`
	FriendlyName                  string `url:"FriendlyName,omitempty"`
	Configuration                 string `url:"Configuration,omitempty"`
	AssignmentCallbackURL         string `url:"AssignmentCallbackURL,omitempty"`
	FallbackAssignmentCallbackURL string `url:"FallbackAssignmentCallbackURL,omitempty"`
	TaskReservationTimeout        int    `url:"TaskReservationTimeout,omitempty"`
}

// TaskQueue taskqueue refer: https://www.twilio.com/docs/taskrouter/api/task-queue
type TaskQueue struct {
	AccountSid              string            `json:"account_sid"`
	AssignmentActivitySid   string            `json:"assignment_activity_sid"`
	AssignmentActivityName  string            `json:"assignment_activity_name"`
	DateCreated             time.Time         `json:"date_created"`
	DateUpdated             time.Time         `json:"date_updated"`
	FriendlyName            string            `json:"friendly_name"`
	MaxReservedWorkers      int               `json:"max_reserved_workers"`
	ReservationActivitySid  string            `json:"reservation_activity_sid"`
	ReservationActivityName string            `json:"reservation_activity_name"`
	Sid                     string            `json:"sid"`
	TargetWorkers           string            `json:"target_workers"`
	TaskOrder               string            `json:"task_order"`
	URI                     string            `json:"url"`
	WorkspaceSid            string            `json:"workspace_sid"`
	Links                   map[string]string `json:"links"`
}

// TaskQueueParams taskQueue parameters refer: https://www.twilio.com/docs/taskrouter/api/task-queue
type TaskQueueParams struct {
	FriendlyName           string `url:"FriendlyName,omitempty"`
	AssignmentActivitySid  string `url:"AssignmentActivitySid,omitempty"`
	MaxReservedWorkers     int    `url:"MaxReservedWorkers,omitempty"`
	TargetWorkers          string `url:"TargetWorkers,omitempty"`
	TaskOrder              string `url:"TaskOrder,omitempty"`
	ReservationActivitySid string `url:"ReservationActivitySid,omitempty"`
}

// Activity activity refer: https://www.twilio.com/docs/taskrouter/api/activity
type Activity struct {
	AccountSid   string    `json:"account_sid"`
	Available    string    `json:"available"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	FriendlyName string    `json:"friendly_name"`
	Sid          string    `json:"sid"`
	WorkspaceSid string    `json:"workspace_sid"`
	URI          string    `json:"url"`
}

// ActivityParams activity refer: https://www.twilio.com/docs/taskrouter/api/activity
type ActivityParams struct {
	Available    string `url:"Available,omitempty"`
	FriendlyName string `url:"FriendlyName,omitempty"`
	WorkspaceSid string `url:"WorkspaceSid,omitempty"`
}
