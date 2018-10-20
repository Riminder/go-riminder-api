package rimindertest

import (
	"testing"

	"github.com/Xalrandion/go-riminder-api"
)

func TestFilterListNoError(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	resp, err := api.Filter().List()

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
	// will fail if team have no filter, tests response parsing also.
	if len(resp) == 0 {
		t.Fatal("No filter retrieved")
	}
}

func TestFilterListExpError(t *testing.T) {
	api := riminder.New("an obviously invalid key")
	_, err := api.Filter().List()

	if err == nil {
		t.Fatalf("A valid response was not expected.")
	}
}

func TestFilterGetNoErrorFID(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	resp, err := api.Filter().Get(map[string]interface{}{"filter_id": GetTestHelper().FilterID})

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
	// tests response parsing also (only on this one because is the more complexe one).
	if resp.FilterID != GetTestHelper().FilterID {
		t.Fatalf("Missmatch filter ids (exp: '%s' got: '%s'", GetTestHelper().FilterID, resp.FilterID)
	}
	if resp.FilterReference == "" || resp.Template.Name == "" {
		t.Fatal("Som field should not be empty (filterReference; Template.Name")
	}
}

func TestFilterGetNoErrorFRef(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	resp, err := api.Filter().Get(map[string]interface{}{"filter_id": GetTestHelper().FilterID})

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
	// tests response parsing also (only on this one because is the more complexe one).
	if resp.FilterReference != GetTestHelper().FilterReference {
		t.Fatalf("Missmatch filter ids (exp: '%s' got: '%s'", GetTestHelper().FilterID, resp.FilterID)
	}
	if resp.FilterID == "" || resp.Template.Name == "" {
		t.Fatal("Som field should not be empty (filterReference; Template.Name")
	}
}
