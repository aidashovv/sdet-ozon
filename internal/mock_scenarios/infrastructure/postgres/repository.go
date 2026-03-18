package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sdet-ozon/internal/mock_scenarios/domain"
	"sdet-ozon/internal/pkg/myerr"
	pgx "sdet-ozon/internal/pkg/postgres"
)

type MockRepository struct {
	pool pgx.Pool
}

func NewMockRepository(pool pgx.Pool) *MockRepository {
	return &MockRepository{
		pool: pool,
	}
}

func (r *MockRepository) Add(ctx context.Context, scenario *domain.MockScenario) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetOperationTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	sDTO, rDTO := ToDBModel(scenario)
	sQuery := `
	INSERT INTO sdetozon.mock_scenarios (test_id, version, status_code)
	VALUES ($1, $2, $3)
    `

	if _, err = tx.Exec(
		ctx, sQuery,
		sDTO.TestID,
		sDTO.Version,
		sDTO.StatusCode,
	); err != nil {
		return fmt.Errorf("insert scenario: %w", err)
	}

	if scenario.Data != nil {
		const rQuery = `
		INSERT INTO sdetozon.exchange_rates (
			rate_id,
			num_code,
			char_code,
			nominal,
			value_name,
			value,
			vunit_rate,
			test_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

		if _, err = tx.Exec(
			ctx, rQuery,
			rDTO.RateID,
			rDTO.NumCode,
			rDTO.CharCode,
			rDTO.Nominal,
			rDTO.ValueName,
			rDTO.Value,
			rDTO.VunitRate,
			sDTO.TestID,
		); err != nil {
			return fmt.Errorf("insert rate: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *MockRepository) GetByTestID(ctx context.Context, id string) (*domain.MockScenario, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetOperationTimeout())
	defer cancel()

	const query = `
	SELECT
		s.test_id, s.version, s.status_code,
		r.rate_id, r.num_code, r.char_code, r.nominal, r.value_name, r.value, r.vunit_rate
	FROM sdetozon.mock_scenarios s
	LEFT JOIN sdetozon.exchange_rates r
	ON s.test_id = r.test_id
	WHERE s.test_id = $1
	`

	var sDTO MockScenarioDTO
	var rDTO ExchangeRateDTO

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&sDTO.TestID, &sDTO.Version, &sDTO.StatusCode,
		&rDTO.RateID,
		&rDTO.NumCode,
		&rDTO.CharCode,
		&rDTO.Nominal,
		&rDTO.ValueName,
		&rDTO.Value,
		&rDTO.VunitRate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myerr.ErrScenarioNotFound
		}
		return nil, fmt.Errorf("get scenario: %w", err)
	}

	return ToDomain(sDTO, rDTO), nil
}

func (r *MockRepository) Update(ctx context.Context, scenario *domain.MockScenario) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetOperationTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("update: begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	sDTO, rDTO := ToDBModel(scenario)

	if scenario.Data != nil {
		rQuery := `
		UPDATE sdetozon.exchange_rates
		SET num_code=$1, char_code=$2, nominal=$3, value_name=$4, value=$5, vunit_rate=$6
		WHERE test_id = $7
		`

		_, err = tx.Exec(
			ctx, rQuery,
			rDTO.NumCode,
			rDTO.CharCode,
			rDTO.Nominal,
			rDTO.ValueName,
			rDTO.Value,
			rDTO.VunitRate,
			sDTO.TestID)
		if err != nil {
			return fmt.Errorf("update rate: %w", err)
		}
	}

	sQuery := `
	UPDATE sdetozon.mock_scenarios
	SET status_code = $1, version = version + 1
	WHERE test_id = $2 AND version = $3
	`

	res, err := tx.Exec(
		ctx, sQuery,
		sDTO.StatusCode,
		sDTO.TestID,
		sDTO.Version)

	if err != nil {
		return fmt.Errorf("update scenario: %w", err)
	}

	if res.RowsAffected() == 0 {
		return myerr.ErrScenarioConflict
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *MockRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetOperationTimeout())
	defer cancel()

	query := `
	DELETE FROM sdetozon.mock_scenarios
	WHERE test_id = $1
	`

	if _, err := r.pool.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("delete scenario: %w", err)
	}

	return nil
}
