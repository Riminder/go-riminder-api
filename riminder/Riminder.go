package riminder

const BASEURL string = "https://www.riminder.net/sf/public/api/v1.0/"

type Riminder struct {
	apiKey     string
	webhookKey string
	baseURL    string
	clientw    *clientw
	source     *source
}

func New(apiKey string) *Riminder {
	r := new(Riminder)

	r.apiKey = apiKey
	r.webhookKey = "NO"
	r.baseURL = BASEURL
	r.clientw = newClientw(r.apiKey, r.baseURL)

	r.source = newSource(r)
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
