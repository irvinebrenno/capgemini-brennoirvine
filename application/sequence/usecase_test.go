package sequence

import (
	"fmt"
	"testing"
)

// TestFinderSequenceValid testa uma matriz com sequência válida
func TestFinderSequenceValid(t *testing.T) {
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", " DDDDUB", "UDBDUH"}
	isValid, err := FinderSequence(matrizTeste)
	if err != nil {
		t.Fatal("Um erro foi encontrado:", err)
	}
	if !isValid {
		t.Fatal("A sequência deveria se encontrada")
	}
}

// TestFinderSequenceInvalid testa uma matriz com sequência inválida
func TestFinderSequenceInvalid(t *testing.T) {
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"BUHDHB", "DBHUHD", "UUBUUU", "BHBDHH", "HDHUDB", "UDBDUH"}
	isValid, err := FinderSequence(matrizTeste)
	if err != nil {
		t.Fatal("Um erro foi encontrado:", err)
	}
	if isValid {
		t.Fatal("A sequência não deveria se encontrada")
	}
}

// TestFinderSequenceMatrizZerada testa o erro de uma matriz zerada
func TestFinderSequenceMatrizZerada(t *testing.T) {
	messageDefaultError := "Não foi possível encontrar sequência"
	errValidacao := fmt.Errorf("%v: Sem Matriz para validar", messageDefaultError)
	matrizTeste := Sequence{}
	_, err := FinderSequence(matrizTeste)
	if err == nil {
		t.Fatal("Um erro de matriz não encontrada deveria ser econtrado")
	}

	if errValidacao.Error() != err.Error() {
		t.Fatal("O erro encontrado deveria ser: ", errValidacao)
	}
}

// TestFinderSequenceJaValidada testa o erro de uma matriz já validada
func TestFinderSequenceJaValidada(t *testing.T) {
	messageDefaultError := "Não foi possível encontrar sequência"
	errValidacao := fmt.Errorf("%v: Essa Matriz já foi validada", messageDefaultError)
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", " DDDDUB", "UDBDUH"}
	_, err := FinderSequence(matrizTeste)
	if err == nil {
		t.Fatal("Um erro de matriz já validada deveria ser encontrado")
	}

	if errValidacao.Error() != err.Error() {
		t.Fatal("O erro encontrado deveria ser: ", errValidacao)
	}
}

// TestFinderSequenceTamanhoInvalido testa o erro de uma matriz de tamanho inválido
func TestFinderSequenceTamanhoInvalido(t *testing.T) {
	messageDefaultError := "Não foi possível encontrar sequência"
	errValidacao := fmt.Errorf("%v: Matriz de tamanho inválido", messageDefaultError)
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"DUH", "DUB", "UBU"}
	_, err := FinderSequence(matrizTeste)
	if err == nil {
		t.Fatal("Um erro de matriz com tamanho inválido deveria ser encontrado")
	}

	if errValidacao.Error() != err.Error() {
		t.Fatal("O erro encontrado deveria ser: ", errValidacao)
	}
}

// TestFinderSequenceNxM testa o erro de uma matriz que não seja NxN
func TestFinderSequenceNxM(t *testing.T) {
	messageDefaultError := "Não foi possível encontrar sequência"
	errValidacao := fmt.Errorf("%v: A Matriz deve ser NxN", messageDefaultError)
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", "DDDDUB"}
	_, err := FinderSequence(matrizTeste)
	if err == nil {
		t.Fatal("Um erro de matriz diferente de NxN deveria ser encontrado")
	}

	if errValidacao.Error() != err.Error() {
		t.Fatal("O erro encontrado deveria ser: ", errValidacao)
	}
}

// TestFinderSequenceLinhaComQuantidadeDiferente testa o erro de uma matriz com linhas de quantidades de letras diferentes
func TestFinderSequenceLinhaComQuantidadeDiferente(t *testing.T) {
	messageDefaultError := "Não foi possível encontrar sequência"
	errValidacao := fmt.Errorf("%v: Todas as linhas devem ter o mesmo número de letras", messageDefaultError)
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", "DDDUB", "UDBDUH"}
	_, err := FinderSequence(matrizTeste)
	if err == nil {
		t.Fatal("Um erro de matriz com linhas de tamanhos diferentes deveria ser encontrado")
	}

	if errValidacao.Error() != err.Error() {
		t.Fatal("O erro encontrado deveria ser: ", errValidacao)
	}
}

// TestFinderSequenceLetrasNaoPermitidas testa o erro de uma matriz com letras inválidas
func TestFinderSequenceLetrasNaoPermitidas(t *testing.T) {
	messageDefaultError := "Não foi possível encontrar sequência"
	errValidacao := fmt.Errorf("%v: Somente as letras B, U, D e H são permitidas", messageDefaultError)
	matrizTeste := Sequence{}
	matrizTeste.Latters = []string{"DUHBHB", "DUJUHD", "UBUUHU", "BHBDHH", " DDDDUB", "UDBDUH"}
	_, err := FinderSequence(matrizTeste)
	if err == nil {
		t.Fatal("Um erro de matriz com letras não permitidas deveria ser encontrado")
	}

	if errValidacao.Error() != err.Error() {
		t.Fatal("O erro encontrado deveria ser: ", errValidacao)
	}
}
