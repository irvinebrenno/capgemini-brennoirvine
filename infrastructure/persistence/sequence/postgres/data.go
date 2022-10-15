package postgres

import (
	"capgemini-brennoirvine/infrastructure/persistence/sequence"
	"database/sql"
	"fmt"
)

type DBSequence struct {
	DB *sql.DB
}

// CheckSequence busca se uma matriz já foi testada no banco de dados
func (postgres *DBSequence) CheckSequence(req *sequence.Sequence) (exists *bool, err error) {

	exists = new(bool)

	if err := postgres.DB.QueryRow(
		`
		SELECT
			CASE
				WHEN COUNT(TS.id) > 0 THEN TRUE
				ELSE FALSE
			END
		FROM t_sequence TS
		WHERE
			TS.sequence ilike $1
		`, req.Latters).Scan(&exists); err != nil {
		return nil, fmt.Errorf("Erro ao checar sequencia no banco de dados")
	}

	return
}

// InsertSequence insere uma matriz testada no banco de dados
func (postgres *DBSequence) InsertSequence(req *sequence.Sequence) (err error) {

	sqlStatement :=
		`INSERT INTO t_sequence
		(sequence, is_valid)
		VALUES($1::TEXT, $2::BOOL)`

	_, err = postgres.DB.Exec(sqlStatement, req.Latters, req.IsValid)
	if err != nil {
		return fmt.Errorf("Erro ao tentar sequencias no banco de dados")
	}

	return nil
}

// GetStats busca stats de matrizes testadas no banco de dados
func (postgres *DBSequence) GetStats() (res *sequence.Stats, err error) {

	res = &sequence.Stats{}

	if err := postgres.DB.QueryRow(
		`
		SELECT
			COALESCE(COUNT(TS.id),0)
		FROM t_sequence TS
		WHERE
			TS.is_valid = $1
		`, true).Scan(&res.CountValid); err != nil {
		return nil, fmt.Errorf("Erro ao buscar sequencias válidas no banco de dados")
	}

	if err := postgres.DB.QueryRow(
		`
		SELECT
			COALESCE(COUNT(TS.id),0)
		FROM t_sequence TS
		WHERE
			TS.is_valid = $1
		`, false).Scan(&res.CountInvalid); err != nil {
		return nil, fmt.Errorf("Erro ao buscar sequencias inválidas no banco de dados")
	}

	return
}
