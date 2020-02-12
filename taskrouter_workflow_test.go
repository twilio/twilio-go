package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestTaskRouterWorkflow_marshall(t *testing.T) {
	testJSONMarshal(t, TaskRouterWorkflow{}, "{}")

	got := &TaskRouterWorkflow{
		AccountSID:                    String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AssignmentCallbackURL:         String("https://example.com/"),
		Configuration:                 nil,
		DateCreated:                   &time.Time{},
		DateUpdated:                   &time.Time{},
		DocumentContentType:           String("application/json"),
		FallbackAssignmentCallbackURL: String("https://example2.com/"),
		FriendlyName:                  String("Sales, Marketing, Support Workflow"),
		SID:                           String("WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		TaskReservationTimeout:        Int(120),
		URL:                           String("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		Links: map[string]*string{
			"statistics": String("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Statistics"),
		},
		WorkspaceSID: String("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
	}

	want := `{
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"assignment_callback_url": "https://example.com/",
		"configuration": null,
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"document_content_type": "application/json",
		"fallback_assignment_callback_url": "https://example2.com/",
		"friendly_name": "Sales, Marketing, Support Workflow",
		"sid": "WWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"task_reservation_timeout": 120,
		"url": "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"workspace_sid": "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"links": {
		  "statistics": "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Statistics"
		}
	  }`

	testJSONMarshal(t, got, want)
}

func TestTaskrouterWorkflow_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		// testFormValues(t, r, values{"FriendlyName": "TaskRouterWorkflow", "Configuration": `"{\"task_routing\":\"JSON\"}"`})
		response := `{"friendly_name":"TaskRouterWorkflow", "configuration": "{\"task_routing\":\"JSON\"}"} }`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workflows.Create("WS123", &TaskRouterWorkflowParams{
		FriendlyName:  String("TaskRouterWorkflow"),
		Configuration: String(`"{\"task_routing\":\"JSON\"}"`),
	})

	if err != nil {
		t.Errorf("TaskRouterWorkflow.Create returned error: %v", err)
	}

	want := &TaskRouterWorkflow{FriendlyName: String("TaskRouterWorkflow"), Configuration: String(`{"task_routing":"JSON"}`)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkflow.Create returned %+v, want %+v", got, want)
	}

}

func TestTaskRouterWorkflow_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Workflows/WF123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"WF123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workflows.Fetch("WS123", "WF123")

	if err != nil {
		t.Errorf("TaskRouterWorkflow.Fetch returned error: %v", err)
	}

	want := &TaskRouterWorkflow{SID: String("WF123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkflow.Fetch returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterWorkflow_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"workflows":[{"sid": "WF123"}]}`
		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workflows.Read("WS123", nil)

	if err != nil {
		t.Errorf("TaskRouterWorkflow.Read returned error: %v", err)
	}

	workflow := &TaskRouterWorkflow{SID: String("WF123")}
	want := &TaskRouterWorkflowList{Workflows: []*TaskRouterWorkflow{workflow}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkflow.Read returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterWorkflow_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Workflows/WF123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"WF123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workflows.Update("WS123", "WF123", &TaskRouterWorkflowParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("TaskRouterWorkflow.Update returned error: %v", err)
	}

	want := &TaskRouterWorkflow{SID: String("WF123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkflow.Update returned %+v, want %+v", got, want)
	}

}

func TestTaskRouterWorkflow_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Workflows/WF123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TaskRouter.Workflows.Delete("WS123", "WF123")

	if err != nil {
		t.Errorf("TaskRouterWorkflow.Delete returned error: %v", err)
	}
}
