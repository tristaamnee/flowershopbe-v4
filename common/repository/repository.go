package repository

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func GetByCondition[T interface{}](db *pg.DB, column string, value any) (*[]*T, error) {
	data := new([]*T)
	err := db.Model(data).Where("? = ANY(?)", pg.Ident(column), value).Select()
	if err != nil {
		log.Println("error getting data from %v ", column, ": %v", err)
		return nil, err
	}
	return data, nil
}
