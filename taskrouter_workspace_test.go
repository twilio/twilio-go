package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTaskrouterWorkspace_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "TaskRouterWorkspace"})
		response := `{"friendly_name":"TaskRouterWorkspace"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Create(&TaskRouterWorkspaceParams{
		FriendlyName: String("TaskRouterWorkspace"),
	})

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Create returned error: %v", err)
	}

	expected := &TaskRouterWorkspace{FriendlyName: String("TaskRouterWorkspace")}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkspace.Create returned %+v, expected %+v", got, expected)
	}

}

func TestTaskRouterWorkspace_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"WF123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Fetch("WS123")

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Fetch returned error: %v", err)
	}

	expected := &TaskRouterWorkspace{Sid: String("WF123")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkspace.Fetch returned %+v, expected %+v", got, expected)
	}
}

func TestTaskRouterWorkspace_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"workspaces":[{"sid": "WF123"}]}`
		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Read(nil)

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Read returned error: %v", err)
	}

	workflow := &TaskRouterWorkspace{Sid: String("WF123")}
	expected := &TaskRouterWorkspaceList{Workspaces: []*TaskRouterWorkspace{workflow}}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkspace.Read returned %+v, expected %+v", got, expected)
	}
}

func TestTaskRouterWorkspace_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"WF123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Update("WS123", &TaskRouterWorkspaceParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Update returned error: %v", err)
	}

	expected := &TaskRouterWorkspace{Sid: String("WF123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterWorkspace.Update returned %+v, expected %+v", got, expected)
	}

}

func TestTaskRouterWorkspace_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TaskRouter.Workspaces.Delete("WS123")

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Delete returned error: %v", err)
	}
}
