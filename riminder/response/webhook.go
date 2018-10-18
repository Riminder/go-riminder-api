package response

// WebhookCheckContainer contains the response representation for webhook.Check method.
type WebhookCheckContainer struct {
	*Container

	Data WebhookCheckElem
}

// WebhookCheckElem the data field response representation for webhook.Check method.
type WebhookCheckElem struct {
	TeamName   string `json:"team_name"`
	WebhookURL string `json:"webhook_url"`
}

// WebhookMessageContainer is the base representation of a webhook message.
type WebhookMessageContainer struct {
	Type string `json:"type"`
}

// WebhookProfile is the representaion of a profile in a webhook message.
type WebhookProfile struct {
	ProfileID        string `json:"profile_id"`
	ProfileReference string `json:"profile_reference"`
}

// WebhookFilter is the representaion of a filter in a webhook message.
type WebhookFilter struct {
	FilterID        string `json:"filter_id"`
	FilterReference string `json:"filter_reference"`
}

// WebhookMessageProfileParse is the message representation for webhook event profile parse.
type WebhookMessageProfileParse struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
}

// WebhookMessageProfileScore is the message representation for webhook event profile score.
type WebhookMessageProfileScore struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
	Filter  WebhookFilter  `json:"filter"`
	Score   float64        `json:"score"`
}

// WebhookMessageFilterTrain is the message representation for webhook event filter train.
type WebhookMessageFilterTrain struct {
	*WebhookCheckContainer

	Message string        `json:"message"`
	Filter  WebhookFilter `json:"filter"`
}

// WebhookMessageFilterScore is the message representation for webhook event filter score.
type WebhookMessageFilterScore struct {
	*WebhookCheckContainer

	Message string        `json:"message"`
	Filter  WebhookFilter `json:"filter"`
}

// WebhookMessageActionStage is the message representation for webhook event action stage.
type WebhookMessageActionStage struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
	Filter  WebhookFilter  `json:"filter"`
	Stage   string         `json:"stage"`
}

// WebhookMessageActionRating is the message representation for webhook event action rating.
type WebhookMessageActionRating struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
	Filter  WebhookFilter  `json:"filter"`
	Rating  string         `json:"rating"`
}
