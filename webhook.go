package riminder

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Xalrandion/go-riminder-api/response"
)

// EventProfileParseSuccess is a string constant for corresponding event.
const EventProfileParseSuccess string = "profile.parse.success"

// EventProfileParseError is a string constant for corresponding event.
const EventProfileParseError string = "profile.parse.error"

// EventProfileScoreSuccess is a string constant for corresponding event.
const EventProfileScoreSuccess string = "profile.score.success"

// EventProfileScoreError is a string constant for corresponding event.
const EventProfileScoreError string = "profile.score.error"

// EventFilterTrainSuccess is a string constant for corresponding event.
const EventFilterTrainSuccess string = "filter.train.success"

// EventFilterTrainError is a string constant for corresponding event.
const EventFilterTrainError string = "filter.train.error"

// EventFilterTrainStart is a string constant for corresponding event.
const EventFilterTrainStart string = "filter.train.start"

// EventFilterScoreSuccess is a string constant for corresponding event.
const EventFilterScoreSuccess string = "filter.score.success"

// EventFilterScoreError is a string constant for corresponding event.
const EventFilterScoreError string = "filter.score.error"

// EventFilterScoreStart is a string constant for corresponding event.
const EventFilterScoreStart string = "filter.score.start"

// ActionStageSuccess is a string constant for corresponding event.
const ActionStageSuccess string = "action.stage.success"

// ActionStageError is a string constant for corresponding event.
const ActionStageError string = "action.stage.error"

// ActionRatingSuccess is a string constant for corresponding event.
const ActionRatingSuccess string = "action.rating.success"

// ActionRatingError is a string constant for corresponding event.
const ActionRatingError string = "action.rating.error"

const webhookHeaderKey string = "HTTP-RIMINDER-SIGNATURE"

// WebhookCallback is the callback called by any webhook event.
type WebhookCallback = func(string, interface{})
type webhookMessageparser = func(s string) interface{}

type webhookHandlerContainer struct {
	callback WebhookCallback
	factory  webhookMessageparser
}

// webhooks class contains methods to interact with the webhooks on riminder's api.
type webhooks struct {
	client     *clientw
	webhookKey string
	handlers   map[string]*webhookHandlerContainer
}

func newWebhooks(riminder *Riminder) *webhooks {
	s := new(webhooks)

	s.client = riminder.clientw
	s.webhookKey = riminder.webhookKey
	s.handlers = map[string]*webhookHandlerContainer{
		EventProfileParseSuccess: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageProfileParse{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventProfileParseError: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageProfileParse{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventProfileScoreSuccess: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageProfileScore{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventProfileScoreError: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageProfileScore{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventFilterTrainSuccess: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageFilterTrain{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventFilterTrainError: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageFilterTrain{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventFilterTrainStart: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageFilterTrain{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventFilterScoreSuccess: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageFilterScore{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventFilterScoreError: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageFilterScore{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		EventFilterScoreStart: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageFilterScore{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		ActionStageSuccess: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageActionStage{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		ActionStageError: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageActionStage{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		ActionRatingSuccess: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageActionRating{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
		ActionRatingError: {factory: func(s string) interface{} {
			r := riminderResponse.WebhookMessageActionRating{}
			json.Unmarshal([]byte(s), &r)
			return r
		}, callback: nil},
	}
	return s
}

func (w *webhooks) setWebhookKey(key string) {
	w.webhookKey = key
}

// Check checks if the webhook integration is enabled.
func (w *webhooks) Check() (riminderResponse.WebhookCheckElem, error) {

	resp := riminderResponse.WebhookCheckContainer{}
	err := w.client.Post("webhook/check", map[string]interface{}{}, &resp)
	if err != nil {
		return riminderResponse.WebhookCheckElem{}, err
	}
	return resp.Data, nil
}

// SetHandler adds an handler for the specified event.
func (w *webhooks) SetHandler(eventName string, callback WebhookCallback) error {
	if _, ok := w.handlers[eventName]; !ok {
		return fmt.Errorf("webhook: '%s' is not a valid webhook event", eventName)
	}
	w.handlers[eventName].callback = callback
	return nil
}

// RemoveHandler removes the handler for the specified event.
func (w *webhooks) RemoveHandler(eventName string) error {
	if _, ok := w.handlers[eventName]; !ok {
		return fmt.Errorf("webhook: '%s' is not a valid webhook event", eventName)
	}
	w.handlers[eventName].callback = nil
	return nil
}

// IsHandlerPresent checks if there is a handler for the specified event.
func (w *webhooks) IsHandlerPresent(eventName string) bool {
	if _, ok := w.handlers[eventName]; !ok {
		return false
	}
	return w.handlers[eventName].callback != nil
}

func getEncodedHeader(receivedHeader interface{}) (string, error) {

	response := ""
	switch rh := receivedHeader.(type) {
	case map[string]string:
		r, ok := rh[webhookHeaderKey]
		if !ok {
			return "", fmt.Errorf("%s key should be part of the header", webhookHeaderKey)
		}
		response = r
	case map[string]interface{}:
		r, ok := rh[webhookHeaderKey].(string)
		if !ok {
			return "", fmt.Errorf("%s key should be part of the header", webhookHeaderKey)
		}
		response = r
	case string:
		response = rh
	case fmt.Stringer:
		response = rh.String()
	default:
		return "", fmt.Errorf("%v is not a valid encoded header (must be a map[string]string or string)", receivedHeader)
	}
	return response, nil
}

func customStrStr(s, toReplace, by string) string {
	res := ""
	for _, c := range s {
		tmpc := byte(c)
		trIDX := strings.Index(toReplace, string(c))
		if trIDX != -1 && trIDX < len(by) {
			tmpc = by[trIDX]
		}
		res += string(tmpc)
	}
	return res
}

func base64UrlDecode(s string) (string, error) {
	decodedStr := customStrStr(s, "-_", "+/")
	decodedStrB, err := base64.StdEncoding.DecodeString(decodedStr)
	if err != nil {
		return "", fmt.Errorf("webhooks.base64decode: %v", err)
	}
	return string(decodedStrB), nil
}

func (w *webhooks) isSignatureValid(payload, sign string) bool {
	hasher := hmac.New(sha256.New, []byte(w.webhookKey))

	hasher.Write([]byte(payload))
	bExpectedSign := hasher.Sum(nil)
	return hmac.Equal(bExpectedSign, []byte(sign))
}

// Handle start a callback for the received header.
func (w *webhooks) Handle(receivedHeaders interface{}) error {
	encodedMessage, err := getEncodedHeader(receivedHeaders)
	if err != nil {
		return fmt.Errorf("webhooks.handle:%v", err)
	}
	tmp := strings.SplitN(encodedMessage, ".", 2)
	sign, err := base64UrlDecode(tmp[0])
	jsonPayload, err1 := base64UrlDecode(tmp[1])
	if err != nil || err1 != nil {
		return fmt.Errorf("webhooks.handle:%v || %v", err, err1)
	}
	if !w.isSignatureValid(jsonPayload, sign) {
		return fmt.Errorf("webhook.handle: invalid signature")
	}

	tmpMessage := riminderResponse.WebhookMessageContainer{}
	err = json.Unmarshal([]byte(jsonPayload), &tmpMessage)
	if err != nil {
		return fmt.Errorf("webhook.handle:cannot parse webhook message: %v", err)
	}
	if _, ok := w.handlers[tmpMessage.Type]; !ok {
		return fmt.Errorf("webhook.handle:'%s' unknown message type", tmpMessage.Type)
	}
	if w.handlers[tmpMessage.Type].callback == nil {
		return nil
	}
	w.handlers[tmpMessage.Type].callback(tmpMessage.Type, w.handlers[tmpMessage.Type].factory(jsonPayload))
	return nil
}
