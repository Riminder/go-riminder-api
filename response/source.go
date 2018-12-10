package riminderResponse

// SourceListContainer contains the response representation for source.List method.
type SourceListContainer struct {
	*Container

	Data []SourceListElem
}

// SourceListElem the data field response representation for source.List method.
type SourceListElem struct {
	SourceID     string   `json:"source_id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Archive      bool     `json:"archive"`
	CountSource  int      `json:"count_source"`
	DateCreation DateTime `json:"date_creation"`
}

// SourceGetContainer contains the response representation for source.Get method.
type SourceGetContainer struct {
	*Container

	Data SourceGetElem
}

// SourceGetElem the data field response representation for source.Get method.
type SourceGetElem struct {
	SourceID     string   `json:"source_id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Archive      bool     `json:"archive"`
	CountSource  int      `json:"count_source"`
	DateCreation DateTime `json:"date_creation"`
}
