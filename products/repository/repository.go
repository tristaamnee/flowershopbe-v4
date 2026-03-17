package repository

import (
	"github.com/go-pg/pg/v10"
	base "github.com/tristaamne/flowershopbe-v4/common/repository"
	"github.com/tristaamne/flowershopbe-v4/products/model"
)

func GetByCategory(db *pg.DB, category string) (*[]*model.Product, error) {
	return base.GetByCondition[model.Product](db, model.ColCategory, category)
}
