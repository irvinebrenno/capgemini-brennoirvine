package sequence

// Sequence define uma estrutura para receber uma sequência/matriz na camada de apresentação
type Sequence struct {
	Latters []string `json:"letters"`
}

// Stats define uma estrutura de stats das matrizes testadas na camada de apresentação
type Stats struct {
	CountValid   *int64   `json:"count_valid"`
	CountInvalid *int64   `json:"count_invalid"`
	Ratio        *float64 `json:"ratio"`
}
