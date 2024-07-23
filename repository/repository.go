package repository

import (
	"fmt"

	"github.com/henrique77/api-quote/model"
	"gorm.io/gorm"
)

type QuoteRepository interface {
	Save([]*model.Quote) error
	GetMetrics(lastQuotes int) (*model.Metrics, error)
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

func (r *quoteRepository) GetMetrics(lastQuotes int) (*model.Metrics, error) {
	metrics := &model.Metrics{}

	if err := r.getResultsPerCarrier(lastQuotes, metrics); err != nil {
		return nil, err
	}

	if err := r.getTotalFinalPricePerCarrier(lastQuotes, metrics); err != nil {
		return nil, err
	}

	if err := r.getAverageFinalPricePerCarrier(lastQuotes, metrics); err != nil {
		return nil, err
	}

	if err := r.getLeastExpensiveShippinh(lastQuotes, metrics); err != nil {
		return nil, err
	}

	if err := r.getMostExpensiveShipping(lastQuotes, metrics); err != nil {
		return nil, err
	}

	return metrics, nil

}

func (r *quoteRepository) getResultsPerCarrier(lastQuotes int, metrics *model.Metrics) error {
	find := []*model.ResultsPerCarrier{}

	query := fmt.Sprintf("SELECT name, COUNT(*) AS quantity FROM %s GROUP BY name;", r.validateSubqueryLastQuotes(lastQuotes))

	if err := r.db.Raw(query).Scan(&find).Error; err != nil {
		return err
	}

	metrics.ResultsPerCarrier = make(map[string]int)

	for _, r := range find {
		metrics.ResultsPerCarrier[r.Name] = r.Quantity
	}

	return nil
}

func (r *quoteRepository) getTotalFinalPricePerCarrier(lastQuotes int, metrics *model.Metrics) error {
	find := []*model.TotalFinalPrice{}

	querry := fmt.Sprintf("SELECT name, ROUND(SUM(price), 2) AS total FROM %s GROUP BY name;", r.validateSubqueryLastQuotes(lastQuotes))

	if err := r.db.Raw(querry).Scan(&find).Error; err != nil {
		return err
	}

	metrics.TotalFinalPrice = make(map[string]float64)

	for _, r := range find {
		metrics.TotalFinalPrice[r.Name] = r.Total
	}

	return nil
}

func (r *quoteRepository) getAverageFinalPricePerCarrier(lastQuotes int, metrics *model.Metrics) error {
	find := []*model.AverageFinalPrice{}

	query := fmt.Sprintf("SELECT name, ROUND(AVG(price), 2) AS average FROM %s GROUP BY name", r.validateSubqueryLastQuotes(lastQuotes))

	if err := r.db.Raw(query).Scan(&find).Error; err != nil {
		return err
	}

	metrics.AverageFinalPrice = make(map[string]float64)

	for _, r := range find {
		metrics.AverageFinalPrice[r.Name] = r.Average
	}

	return nil
}

func (r *quoteRepository) getLeastExpensiveShippinh(lastQuotes int, metrics *model.Metrics) error {
	query := fmt.Sprintf("SELECT MIN(price) FROM %s;", r.validateSubqueryLastQuotes(lastQuotes))

	return r.db.Raw(query).Scan(&metrics.LeastExpensiveShipping).Error
}

func (r *quoteRepository) getMostExpensiveShipping(lastQuotes int, metrics *model.Metrics) error {
	query := fmt.Sprintf("SELECT MAX(price) FROM %s;", r.validateSubqueryLastQuotes(lastQuotes))

	return r.db.Raw(query).Scan(&metrics.MostExpensiveShipping).Error
}

func (r *quoteRepository) validateSubqueryLastQuotes(lastQuotes int) string {
	if lastQuotes > 0 {
		return fmt.Sprintf(`(SELECT * FROM quotes ORDER BY id DESC LIMIT %d) AS subquery`, lastQuotes)
	}

	return "quotes"
}
