package controller

import (
	"database/sql"

	"services/internal/budgeting/db/model"
	"services/internal/budgeting/db/view"
	q "services/internal/utils/sql/querybuilder"
)

func CreateIncome(db *sql.DB, m model.Income) error {
	query, err := q.BuildCreateQuery("incomes", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateIncome(db *sql.DB, setM model.Income, whereM model.Income) error {
	query, err := q.BuildUpdateQuery("incomes", setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveIncome(db *sql.DB, m model.Income) ([]model.Income, error) {
	query, err := q.BuildRetrieveQuery("incomes", m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		return view.ViewIncome(rows), nil
	}
	return nil, err
}
