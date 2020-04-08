package data

import (
	"database/sql"

	"github.com/pkg/errors"
)

func (db *postgresRepo) AlterationMethodList() ([]AlterationMethod, error) {
	rows, err := db.Query("select id,default_name from alteration_method;")
	if err != nil {
		return nil, errors.Wrap(
			err,
			"AlterationMethodList: select query failed",
		)
	}
	defer rows.Close()
	list := []AlterationMethod{}
	for rows.Next() {
		a := AlterationMethod{}
		err := rows.Scan(
			&a.ID,
			&a.DefaultName,
		)
		if err != nil {
			return nil, errors.Wrap(
				err,
				"AlterationMethodList: scan failed",
			)
		}
		list = append(
			list,
			a,
		)
	}
	return list, nil
}
func (db *postgresRepo) AlterationMethodGet(id int) (*AlterationMethod, error) {
	a := &AlterationMethod{}
	err := db.QueryRow(
		"select id,default_name from alteration_method where id=$1;",
		id,
	).Scan(
		&a.ID,
		&a.DefaultName,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(
			err,
			"AlterationMethodList: scan failed",
		)
	}

	return a, nil
}
