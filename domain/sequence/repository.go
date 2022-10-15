package sequence

import (
	"capgemini-brennoirvine/infrastructure/persistence/sequence"
	"capgemini-brennoirvine/infrastructure/persistence/sequence/postgres"
	"database/sql"
	"strings"
)

// repository -
type repository struct {
	Data *postgres.DBSequence
}

// newRepo -
func newRepo(newDB *sql.DB) *repository {
	return &repository{
		Data: &postgres.DBSequence{DB: newDB},
	}
}

// CheckSequence é um gerenciado de fluxo para checar se uma matriz já foi validada no banco de dados postgres
func (r *repository) CheckSequence(req *sequence.Sequence) (exists *bool, err error) {
	return r.Data.CheckSequence(req)
}

// InsertSequence insere uma matriz validada no banco de dados postgres
func (r *repository) InsertSequence(req *sequence.Sequence) (err error) {
	return r.Data.InsertSequence(req)
}

// GetStats busca no bando de dados os valores testados
func (r *repository) GetStats() (res *sequence.Stats, err error) {
	return r.Data.GetStats()
}

// ConvertSequenceToData converte uma estrutura da camada de apresentação para a camada de data
func (r *repository) ConvertSequenceToData(latters []string) *sequence.Sequence {
	return &sequence.Sequence{
		Latters: strings.Join(latters, ""),
	}
}
