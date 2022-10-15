package sequence

// Sequence define uma estrutura para salvar matrizes testadas no bando de dados
type Sequence struct {
	Latters string
	IsValid bool
}

// Stats define uma estrutura para retornar stats de matrizes testadas no banco de dados
type Stats struct {
	CountValid   *int64
	CountInvalid *int64
}
