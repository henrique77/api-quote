package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique77/api-quote/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, err
}

const (
	resultsPerCarrier               = "SELECT name, COUNT(*) AS quantity FROM quotes GROUP BY name;"
	totalFinalPrice                 = "SELECT name, ROUND(SUM(price), 2) AS total FROM quotes GROUP BY name;"
	averageFinalPrice               = "SELECT name, ROUND(AVG(price), 2) AS average FROM quotes GROUP BY name;"
	leastExpensiveShipping          = "SELECT MIN(price) FROM quotes;"
	mostExpensiveShipping           = "SELECT MAX(price) FROM quotes;"
	resultsPerCarrierWithLimit      = "SELECT name, COUNT(*) AS quantity FROM (SELECT * FROM quotes ORDER BY id DESC LIMIT 10) AS subquery  GROUP BY name;"
	totalFinalPriceWithLimit        = "SELECT name, ROUND(SUM(price), 2) AS total FROM (SELECT * FROM quotes ORDER BY id DESC LIMIT 10) AS subquery  GROUP BY name;"
	averageFinalPriceWithLimit      = "SELECT name, ROUND(AVG(price), 2) AS average FROM (SELECT * FROM quotes ORDER BY id DESC LIMIT 10) AS subquery  GROUP BY name;"
	leastExpensiveShippingWithLimit = "SELECT MIN(price) FROM (SELECT * FROM quotes ORDER BY id DESC LIMIT 10) AS subquery ;"
	mostExpensiveShippingWithLimit  = "SELECT MAX(price) FROM (SELECT * FROM quotes ORDER BY id DESC LIMIT 10) AS subquery ;"
	queryInsert                     = "INSERT INTO `quotes` (`name`,`service`,`deadline`,`price`) VALUES (?,?,?,?)"
)

func Test_quoteRepository_Save(t *testing.T) {
	require := require.New(t)
	db, mock, err := getMockDB()
	require.NoError(err)
	repository := NewQuoteRepository(db)

	quotes := []*model.Quote{
		{
			Name:     "transportadora_1",
			Service:  "servico_1",
			Deadline: 1,
			Price:    1.11,
		},
	}

	t.Run("quando houver falha ao salvar, deverá retornar erro", func(t *testing.T) {
		expectedError := errors.New("save error")

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(queryInsert)).WithArgs(quotes[0].Name, quotes[0].Service, quotes[0].Deadline, quotes[0].Price).WillReturnError(expectedError)
		mock.ExpectRollback()

		err = repository.Save(quotes)
		require.Error(err)
		require.Equal(expectedError, err)

		require.NoError(mock.ExpectationsWereMet())
	})

	t.Run("quando não houver falha ao salvar, não deverá retornar erro", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(queryInsert)).WithArgs(quotes[0].Name, quotes[0].Service, quotes[0].Deadline, quotes[0].Price).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repository.Save(quotes)
		require.NoError(err)
		require.NoError(mock.ExpectationsWereMet())
	})
}

func Test_quoteRepository_GetMetrics(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	db, mock, err := getMockDB()
	require.NoError(err)
	repository := NewQuoteRepository(db)

	expectedResultsPerCarrier := map[string]int{"transportadora_1": 1}
	expectedTotalFinalPrice := map[string]float64{"transportadora_1": 1.10}
	expectedAverageFinalPrice := map[string]float64{"transportadora_1": 2.20}
	expectedLeastExpensiveShipping := 1.10
	expectedMostExpensiveShipping := 2.20

	t.Run("Quando o lastQuotes for igual a 0, deverá executar as querys sem limite", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrier)).WillReturnRows(sqlmock.NewRows([]string{"name", "quantity"}).AddRow("transportadora_1", 1))
		mock.ExpectQuery(regexp.QuoteMeta(totalFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "total"}).AddRow("transportadora_1", 1.10))
		mock.ExpectQuery(regexp.QuoteMeta(averageFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "average"}).AddRow("transportadora_1", 2.20))
		mock.ExpectQuery(regexp.QuoteMeta(leastExpensiveShipping)).WillReturnRows(sqlmock.NewRows([]string{"MIN(price)"}).AddRow(1.10))
		mock.ExpectQuery(regexp.QuoteMeta(mostExpensiveShipping)).WillReturnRows(sqlmock.NewRows([]string{"MAX(price)"}).AddRow(2.20))

		metrics, err := repository.GetMetrics(0)
		require.NoError(err)
		require.NoError(mock.ExpectationsWereMet())

		assert.Equal(expectedResultsPerCarrier, metrics.ResultsPerCarrier)
		assert.Equal(expectedTotalFinalPrice, metrics.TotalFinalPrice)
		assert.Equal(expectedAverageFinalPrice, metrics.AverageFinalPrice)
		assert.Equal(expectedLeastExpensiveShipping, metrics.LeastExpensiveShipping)
		assert.Equal(expectedMostExpensiveShipping, metrics.MostExpensiveShipping)
	})

	t.Run("Quando o lastQuotes for diferente de 0, deverá executar as querys com limite informado", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrierWithLimit)).WillReturnRows(sqlmock.NewRows([]string{"name", "quantity"}).AddRow("transportadora_1", 1))
		mock.ExpectQuery(regexp.QuoteMeta(totalFinalPriceWithLimit)).WillReturnRows(sqlmock.NewRows([]string{"name", "total"}).AddRow("transportadora_1", 1.10))
		mock.ExpectQuery(regexp.QuoteMeta(averageFinalPriceWithLimit)).WillReturnRows(sqlmock.NewRows([]string{"name", "average"}).AddRow("transportadora_1", 2.20))
		mock.ExpectQuery(regexp.QuoteMeta(leastExpensiveShippingWithLimit)).WillReturnRows(sqlmock.NewRows([]string{"MIN(price)"}).AddRow(1.10))
		mock.ExpectQuery(regexp.QuoteMeta(mostExpensiveShippingWithLimit)).WillReturnRows(sqlmock.NewRows([]string{"MAX(price)"}).AddRow(2.20))

		metrics, err := repository.GetMetrics(10)
		require.NoError(err)
		require.NoError(mock.ExpectationsWereMet())

		assert.Equal(expectedResultsPerCarrier, metrics.ResultsPerCarrier)
		assert.Equal(expectedTotalFinalPrice, metrics.TotalFinalPrice)
		assert.Equal(expectedAverageFinalPrice, metrics.AverageFinalPrice)
		assert.Equal(expectedLeastExpensiveShipping, metrics.LeastExpensiveShipping)
		assert.Equal(expectedMostExpensiveShipping, metrics.MostExpensiveShipping)
	})

	t.Run("Quando houver erro ao consultar resultsPerCarrier, deverá retornar o erro", func(t *testing.T) {
		expectedError := errors.New("error resultsPerCarrier")

		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrier)).WillReturnError(expectedError)

		metrics, err := repository.GetMetrics(0)
		require.Nil(metrics)
		require.Error(err)
		require.Equal(expectedError, err)
		require.NoError(mock.ExpectationsWereMet())
	})

	t.Run("Quando houver erro ao consultar totalFinalPrice, deverá retornar o erro", func(t *testing.T) {
		expectedError := errors.New("error totalFinalPrice")

		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrier)).WillReturnRows(sqlmock.NewRows([]string{"name", "quantity"}).AddRow("transportadora_1", 1))
		mock.ExpectQuery(regexp.QuoteMeta(totalFinalPrice)).WillReturnError(expectedError)

		metrics, err := repository.GetMetrics(0)
		require.Nil(metrics)
		require.Error(err)
		require.Equal(expectedError, err)
		require.NoError(mock.ExpectationsWereMet())
	})

	t.Run("Quando houver erro ao consultar averageFinalPrice, deverá retornar o erro", func(t *testing.T) {
		expectedError := errors.New("error averageFinalPrice")

		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrier)).WillReturnRows(sqlmock.NewRows([]string{"name", "quantity"}).AddRow("transportadora_1", 1))
		mock.ExpectQuery(regexp.QuoteMeta(totalFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "total"}).AddRow("transportadora_1", 1.10))
		mock.ExpectQuery(regexp.QuoteMeta(averageFinalPrice)).WillReturnError(expectedError)

		metrics, err := repository.GetMetrics(0)
		require.Nil(metrics)
		require.Error(err)
		require.Equal(expectedError, err)
		require.NoError(mock.ExpectationsWereMet())
	})

	t.Run("Quando houver erro ao consultar leastExpensiveShipping, deverá retornar o erro", func(t *testing.T) {
		expectedError := errors.New("error leastExpensiveShipping")

		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrier)).WillReturnRows(sqlmock.NewRows([]string{"name", "quantity"}).AddRow("transportadora_1", 1))
		mock.ExpectQuery(regexp.QuoteMeta(totalFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "total"}).AddRow("transportadora_1", 1.10))
		mock.ExpectQuery(regexp.QuoteMeta(averageFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "average"}).AddRow("transportadora_1", 2.20))
		mock.ExpectQuery(regexp.QuoteMeta(leastExpensiveShipping)).WillReturnError(expectedError)

		metrics, err := repository.GetMetrics(0)
		require.Nil(metrics)
		require.Error(err)
		require.Equal(expectedError, err)
		require.NoError(mock.ExpectationsWereMet())
	})

	t.Run("Quando houver erro ao consultar mostExpensiveShipping, deverá retornar o erro", func(t *testing.T) {
		expectedError := errors.New("error mostExpensiveShipping")

		mock.ExpectQuery(regexp.QuoteMeta(resultsPerCarrier)).WillReturnRows(sqlmock.NewRows([]string{"name", "quantity"}).AddRow("transportadora_1", 1))
		mock.ExpectQuery(regexp.QuoteMeta(totalFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "total"}).AddRow("transportadora_1", 1.10))
		mock.ExpectQuery(regexp.QuoteMeta(averageFinalPrice)).WillReturnRows(sqlmock.NewRows([]string{"name", "average"}).AddRow("transportadora_1", 2.20))
		mock.ExpectQuery(regexp.QuoteMeta(leastExpensiveShipping)).WillReturnRows(sqlmock.NewRows([]string{"MIN(price)"}).AddRow(1.10))
		mock.ExpectQuery(regexp.QuoteMeta(mostExpensiveShipping)).WillReturnError(expectedError)

		metrics, err := repository.GetMetrics(0)
		require.Nil(metrics)
		require.Error(err)
		require.Equal(expectedError, err)
		require.NoError(mock.ExpectationsWereMet())
	})
}
