package response

type SourceListContainer struct {
	*ResponseContainer

	Data []SourceListElem
}

type SourceListElem struct {
	SourceID     string   `json:"source_id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Archive      bool     `json:"archive"`
	CountSource  int      `json:"count_source"`
	DateCreation DateTime `json:"date_creation"`
}

type SourceGetContainer struct {
	*ResponseContainer

	Data SourceGetElem
}

type SourceGetElem struct {
	SourceID     string   `json:"source_id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Archive      bool     `json:"archive"`
	CountSource  int      `json:"count_source"`
	DateCreation DateTime `json:"date_creation"`
}
