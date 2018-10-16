package response

type WebhookCheckContainer struct {
	*ResponseContainer

	Data []WebhookCheckElem
}

type WebhookCheckElem struct {
	TeamName   string `json:"team_name"`
	WebhookURL string `json:"webhook_url"`
}

type WebhookMessageContainer struct {
	Type string `json:"type"`
}

type WebhookProfile struct {
	ProfileID        string `json:"profile_id"`
	ProfileReference string `json:"profile_reference"`
}

type WebhookFilter struct {
	FilterID        string `json:"filter_id"`
	FilterReference string `json:"filter_reference"`
}

type WebhookMessageProfileParse struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
}

type WebhookMessageProfileScore struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
	Filter  WebhookFilter  `json:"filter"`
	Score   float64        `json:"score"`
}

type WebhookMessageFilterTrain struct {
	*WebhookCheckContainer

	Message string        `json:"message"`
	Filter  WebhookFilter `json:"filter"`
}

type WebhookMessageFilterScore struct {
	*WebhookCheckContainer

	Message string        `json:"message"`
	Filter  WebhookFilter `json:"filter"`
}

type WebhookMessageActionStage struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
	Filter  WebhookFilter  `json:"filter"`
	Stage   string         `json:"stage"`
}

type WebhookMessageActionRating struct {
	*WebhookCheckContainer

	Message string         `json:"message"`
	Profile WebhookProfile `json:"profile"`
	Filter  WebhookFilter  `json:"filter"`
	Rating  string         `json:"rating"`
}
