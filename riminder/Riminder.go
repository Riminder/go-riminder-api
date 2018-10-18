package riminder

const BASEURL string = "https://www.riminder.net/sf/public/api/v1.0/"

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

func New(apiKey string) *Riminder {
	r := new(Riminder)

	r.apiKey = apiKey
	r.webhookKey = "NO"
	r.baseURL = BASEURL
	r.clientw = newClientw(r.apiKey, r.baseURL)

	r.source = newSource(r)
	r.filter = newfilter(r)
	r.profile = newprofile(r)
	r.webhooks = newWebhooks(r)
	return r
}

func (r *Riminder) SetBaseURL(url string) {
	r.baseURL = url
	r.clientw.SetBaseURL(r.baseURL)
}

func (r *Riminder) SetWebhookKey(wkey string) {
	r.webhookKey = wkey
}

func (r *Riminder) Source() *source {
	return r.source
}

func (r *Riminder) Filter() *filter {
	return r.filter
}

func (r *Riminder) Profile() *profile {
	return r.profile
}

func (r *Riminder) Webhooks() *webhooks {
	return r.webhooks
}
