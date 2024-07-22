package repository

import (
	"github.com/henrique77/api-quote/model"
	"gorm.io/gorm"
)

type QuoteRepository interface {
	Save([]*model.Quote) error
}

type quoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) QuoteRepository {
	return &quoteRepository{
		db: db,
	}
}

func (r *quoteRepository) Save(quotes []*model.Quote) error {
	return r.db.CreateInBatches(quotes, 100).Error
}
