package twilio

import (
	"time"
)

//Workspace refer: https://www.twilio.com/docs/taskrouter/api/workspace
type Workspace struct {
	Sid                  string            `json:"sid,omitempty"`
	AccountSid           string            `json:"account_sid,omitempty"`
	DateCreated          time.Time         `json:"date_created,omitempty"`
	DateUpdated          time.Time         `json:"date_updated,omitempty"`
	DefaultActivityName  string            `json:"default_activity_name,omitempty"`
	DefaultActivitySid   string            `json:"default_activity_sid,omitempty"`
	TimeoutActivityName  string            `json:"timeout_activity_name,omitempty"`
	TimeoutActivitySid   string            `json:"timeout_activity_sid,omitempty"`
	URL                  string            `json:"url,omitempty"`
	Links                map[string]string `json:"links,omitempty"`
	FriendlyName         string            `json:"friendly_name,omitempty"`
	EventCallbackURL     string            `json:"event_callback_url,omitempty"`
	EventsFilter         string            `json:"events_filter,omitempty"`
	MultitaskEnabled     bool              `json:"multi_task_enabled,omitempty"`
	Template             string            `json:"template,omitempty"`
	PrioritizeQueueOrder string            `json:"prioritize_queue_order,omitempty"`
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
	WorkspaceDefaultFriendlyName = "Flex Task Assignment5"
	WorkflowDefaultFriendlyName  = "Assign to Anyone"
	TaskQueueDefaultFriendlyName = "Everyone"
	ActivityDefaultFriendlyName  = "Break"
)

// Workflow workflow refer: https://www.twilio.com/docs/taskrouter/api/workflow
type Workflow struct {
	AssignmentCallbackURL         string            `json:"assignment_callback_url,omitempty"`
	Configuration                 string            `json:"configuration,omitempty"`
	AccountSid                    string            `json:"account_sid,omitempty"`
	DateCreated                   time.Time         `json:"date_created,omitempty"`
	DateUpdated                   time.Time         `json:"date_updated,omitempty"`
	DocumentContentType           string            `json:"document_content_type,omitempty"`
	FallbackAssignmentCallbackURL string            `json:"fallback_assignment_callback_url,omitempty"`
	FriendlyName                  string            `json:"friendly_name,omitempty"`
	Sid                           string            `json:"sid,omitempty"`
	TaskReservationTimeout        int               `json:"task_reservation_timeout,omitempty"`
	WorkspaceSid                  string            `json:"workspace_sid,omitempty"`
	URL                           string            `json:"url,omitempty"`
	Links                         map[string]string `json:"links,omitempty"`
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
	AccountSid              string            `json:"account_sid,omitempty"`
	AssignmentActivitySid   string            `json:"assignment_activity_sid,omitempty"`
	AssignmentActivityName  string            `json:"assignment_activity_name,omitempty"`
	DateCreated             time.Time         `json:"date_created,omitempty"`
	DateUpdated             time.Time         `json:"date_updated,omitempty"`
	FriendlyName            string            `json:"friendly_name,omitempty"`
	MaxReservedWorkers      int               `json:"max_reserved_workers,omitempty"`
	ReservationActivitySid  string            `json:"reservation_activity_sid,omitempty"`
	ReservationActivityName string            `json:"reservation_activity_name,omitempty"`
	Sid                     string            `json:"sid,omitempty"`
	TargetWorkers           string            `json:"target_workers,omitempty"`
	TaskOrder               string            `json:"task_order,omitempty"`
	URI                     string            `json:"url,omitempty"`
	WorkspaceSid            string            `json:"workspace_sid,omitempty"`
	Links                   map[string]string `json:"links,omitempty"`
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
	AccountSid   string    `json:"account_sid,omitempty"`
	Available    string    `json:"available,omitempty"`
	DateCreated  time.Time `json:"date_created,omitempty"`
	DateUpdated  time.Time `json:"date_updated,omitempty"`
	FriendlyName string    `json:"friendly_name,omitempty"`
	Sid          string    `json:"sid,omitempty"`
	WorkspaceSid string    `json:"workspace_sid,omitempty"`
	URI          string    `json:"url,omitempty"`
}

// ActivityParams activity refer: https://www.twilio.com/docs/taskrouter/api/activity
type ActivityParams struct {
	Available    string `url:"Available,omitempty"`
	FriendlyName string `url:"FriendlyName,omitempty"`
	WorkspaceSid string `url:"WorkspaceSid,omitempty"`
}
