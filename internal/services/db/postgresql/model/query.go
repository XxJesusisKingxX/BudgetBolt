package model
type QueryParameters struct {
	Asc     bool
	Desc    bool
	OrderBy string
}

type Querys struct {
	Select QueryParameters
	Update QueryParameters
	Delete QueryParameters
	Insert QueryParameters
}

