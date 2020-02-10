package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTaskrouterWorkflow_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "TaskRouterWorkflow", "Configuration": `"{\"task_routing\":\"JSON\"}"`})
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

	expected := &TaskRouterWorkflow{FriendlyName: String("TaskRouterWorkflow"), Configuration: String(`{"task_routing":"JSON"}`)}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkflow.Create returned %+v, expected %+v", got, expected)
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

	expected := &TaskRouterWorkflow{Sid: String("WF123")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkflow.Fetch returned %+v, expected %+v", got, expected)
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

	workflow := &TaskRouterWorkflow{Sid: String("WF123")}
	expected := &TaskRouterWorkflowList{Workflows: []*TaskRouterWorkflow{workflow}}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkflow.Read returned %+v, expected %+v", got, expected)
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

	expected := &TaskRouterWorkflow{Sid: String("WF123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkflow.Update returned %+v, expected %+v", got, expected)
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
