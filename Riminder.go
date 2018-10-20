package riminder

// BaseURL is the base api url of riminder's api.
const BaseURL string = "https://www.riminder.net/sf/public/api/v1.0/"

// Riminder is the class that permit interaction with riminder's api.
type Riminder struct {
	apiKey     string
	webhookKey string
	baseURL    string
	clientw    *clientw
	source     *source
	filter     *filter
	profile    *profile
	webhooks   *webhooks
}

// New creates a new Riminder class.
func New(apiKey string) *Riminder {
	r := new(Riminder)

	r.apiKey = apiKey
	r.webhookKey = "NO"
	r.baseURL = BaseURL
	r.clientw = newClientw(r.apiKey, r.baseURL)

	r.source = newSource(r)
	r.filter = newfilter(r)
	r.profile = newprofile(r)
	r.webhooks = newWebhooks(r)
	return r
}

// SetBaseURL changes the baseURL for the nexts calls.
func (r *Riminder) SetBaseURL(url string) {
	r.baseURL = url
	r.clientw.SetBaseURL(r.baseURL)
}

// SetWebhookKey set a webhookKey.
func (r *Riminder) SetWebhookKey(wkey string) {
	r.webhookKey = wkey
	r.webhooks.setWebhookKey(r.webhookKey)
}

// Source permits to access source related calls.
func (r *Riminder) Source() *source {
	return r.source
}

// Filter permits to access filter related calls.
func (r *Riminder) Filter() *filter {
	return r.filter
}

// Profile permits to access profile related calls.
func (r *Riminder) Profile() *profile {
	return r.profile
}

// Webhooks permits to access webhooks related calls.
func (r *Riminder) Webhooks() *webhooks {
	return r.webhooks
}
