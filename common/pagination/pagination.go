package pagination

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PaginationQuery struct {
	Page   int64  `form:"page"`
	Limit  int64  `form:"limit"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
}

func (p *PaginationQuery) GetSkip() int64 {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return (p.Page - 1) * p.Limit
}

func (p *PaginationQuery) GetLimit() int64 {
	if p.Limit <= 0 {
		return 10
	}
	if p.Limit > 100 {
		return 100
	}
	return p.Limit
}

func ParsePagingOption(pg PaginationQuery) *options.FindOptions {
	opts := options.Find()

	opts.SetLimit(pg.GetLimit())
	opts.SetSkip(pg.GetSkip())

	if pg.SortBy != "" {
		direction := 1 // Tăng dần (asc)
		if strings.ToLower(pg.Order) == "desc" {
			direction = -1 // Giảm dần (desc)
		}
		opts.SetSort(bson.D{{Key: pg.SortBy, Value: direction}})
	}

	opts.SetProjection(bson.M{
		"password": 0,
		"otp":      0,
	})

	return opts
}
