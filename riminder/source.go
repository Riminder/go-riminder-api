package riminder

import "github.com/Xalrandion/go-riminder-api/riminder/response"

type source struct {
	client *clientw
}

func newSource(riminder *Riminder) *source {
	s := new(source)

	s.client = riminder.clientw
	return s
}

func (s *source) List() ([]response.SourceListElem, error) {

	resp := response.SourceListContainer{}
	tmpResp, err := s.client.Get("sources", map[string]string{}, resp)
	if tmpResp == nil || err != nil {
		return nil, err
	}
	return tmpResp.(*response.SourceListContainer).Data, nil
}

func (s *source) Get(options map[string]interface{}) (response.SourceGetElem, error) {

	query := map[string]string{
		"source_id": options["source_id"].(string),
	}

	resp := response.SourceGetContainer{}
	tmpResp, err := s.client.Get("source", query, resp)
	if tmpResp == nil || err != nil {
		return response.SourceGetElem{}, err
	}
	return tmpResp.(*response.SourceGetContainer).Data, nil
}
