package model

type QueryParameters struct {
	OrderBy OrderBy
	GreaterThanEq GreaterThanEq
}

type OrderBy struct {
	Asc     bool
	Desc    bool
	Column  string
}

type GreaterThanEq struct {
	Column  string
	Value   string
}

type Querys struct {
	Select QueryParameters
	Update QueryParameters
	Delete QueryParameters
	Insert QueryParameters
}

