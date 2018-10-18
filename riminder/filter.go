package riminder

import "github.com/Xalrandion/go-riminder-api/riminder/response"

// filter class contains methods to interact with the filters on riminder's api.
type filter struct {
	client *clientw
}

func newfilter(riminder *Riminder) *filter {
	s := new(filter)

	s.client = riminder.clientw
	return s
}

// List returns the list of filters your team can access.
func (s *filter) List() ([]response.FilterListElem, error) {

	resp := response.FilterListContainer{}
	err := s.client.Get("filters", map[string]string{}, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// Get returns a specific filter given a filter_id or filter_reference.
func (s *filter) Get(options map[string]interface{}) (response.FilterGetElem, error) {

	query := map[string]string{}
	AddIfNotEmptyStrMap(&query, options, "filter_id")
	AddIfNotEmptyStrMap(&query, options, "filter_reference")

	resp := response.FilterGetContainer{}
	err := s.client.Get("filter", query, &resp)
	if err != nil {
		return response.FilterGetElem{}, err
	}
	return resp.Data, nil
}
