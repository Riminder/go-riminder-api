package rimindertest

import (
	"testing"

	"github.com/Xalrandion/go-riminder-api/riminder/response"

	"github.com/Xalrandion/go-riminder-api/riminder"
)

func TestWebhooksCheckNoErrorFID(t *testing.T) {
	t.Skip("Some server issue")
	api := riminder.New(GetTestHelper().APIKey)
	_, err := api.Webhooks().Check()

	if err != nil {
		t.Fatalf("Invalid response: %v", err)
	}
}

func TestWebhookHandle(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	api.SetWebhookKey(GetTestHelper().WebhookKey)

	pomme := 0
	pID := ""
	cb := func(Type string, message interface{}) {
		pomme = 1
		trueMessage := message.(riminderResponse.WebhookMessageProfileParse)
		pID = trueMessage.Profile.ProfileID
	}
	err := api.Webhooks().SetHandler(riminder.EventProfileParseSuccess, cb)
	if err != nil {
		t.Fatalf("'%v' should be a valid webhook event: %v.", riminder.EventProfileParseSuccess, err)
	}

	req := map[string]string{
		"HTTP-RIMINDER-SIGNATURE": GetTestHelper().genWebhookRequest(),
	}
	err = api.Webhooks().Handle(req)
	if err != nil {
		t.Fatalf("Error during webhook handle: %v", err)
	}
	if pomme != 1 || pID != "hey" {
		t.Fatal("Callback not called")
	}
}

func TestWebhookHandleHeaderNoMap(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	api.SetWebhookKey(GetTestHelper().WebhookKey)

	pomme := 0
	pID := ""
	cb := func(Type string, message interface{}) {
		pomme = 1
		trueMessage := message.(riminderResponse.WebhookMessageProfileParse)
		pID = trueMessage.Profile.ProfileID
	}
	err := api.Webhooks().SetHandler(riminder.EventProfileParseSuccess, cb)
	if err != nil {
		t.Fatalf("'%v' should be a valid webhook event: %v.", riminder.EventProfileParseSuccess, err)
	}

	req := map[string]string{
		"HTTP-RIMINDER-SIGNATURE": GetTestHelper().genWebhookRequest(),
	}
	err = api.Webhooks().Handle(req["HTTP-RIMINDER-SIGNATURE"])
	if err != nil {
		t.Fatalf("Error during webhook handle: %v", err)
	}
	if pomme != 1 || pID != "hey" {
		t.Fatal("Callback not called")
	}
}

func TestWebhookHandleNoHandler(t *testing.T) {
	api := riminder.New(GetTestHelper().APIKey)
	api.SetWebhookKey(GetTestHelper().WebhookKey)

	pomme := 0
	pID := ""
	cb := func(Type string, message interface{}) {
		pomme = 1
		trueMessage := message.(riminderResponse.WebhookMessageProfileParse)
		pID = trueMessage.Profile.ProfileID
	}
	err := api.Webhooks().SetHandler(riminder.EventProfileParseSuccess, cb)
	if err != nil {
		t.Fatalf("'%v' should be a valid webhook event: %v.", riminder.EventProfileParseSuccess, err)
	}

	req := map[string]string{
		"HTTP-RIMINDER-SIGNATURE": GetTestHelper().genWebhookRequest(),
	}

	if !api.Webhooks().IsHandlerPresent(riminder.EventProfileParseSuccess) {
		t.Fatal("Handler is enabled")
	}
	api.Webhooks().RemoveHandler(riminder.EventProfileParseSuccess)
	if api.Webhooks().IsHandlerPresent(riminder.EventProfileParseSuccess) {
		t.Fatal("Handler is not enabled")
	}
	err = api.Webhooks().Handle(req["HTTP-RIMINDER-SIGNATURE"])
	if err != nil {
		t.Fatalf("Error during webhook handle: %v", err)
	}
	if pomme == 1 || pID == "hey" {
		t.Fatal("Callback  called")
	}
}
