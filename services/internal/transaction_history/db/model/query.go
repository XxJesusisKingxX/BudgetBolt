package model

type QueryParameters struct {
	OrderBy OrderBy
	GreaterThanEq GreaterThanEq
	Equal Equal
}

type OrderBy struct {
	Asc     bool
	Desc    bool
	Column  string
}

type GreaterThanEq struct {
	Column  string
	Value   interface{}
}

type Equal struct {
	Column  string
	Value   interface{}
}

type Querys struct {
	Select QueryParameters
	Update QueryParameters
	Delete QueryParameters
	Insert QueryParameters
}

