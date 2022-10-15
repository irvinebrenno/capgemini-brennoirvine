package sequence

import (
	"capgemini-brennoirvine/config/database"
	"capgemini-brennoirvine/domain/sequence"
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

// FinderSequence contém a lógica de negócio para validar sequência e adicionar a matriz no banco de dados
func FinderSequence(c *gin.Context, req Sequence) (isValid bool, err error) {
	messageDefaultError := "Não foi possível encontrar sequência"

	db := database.Conectar()
	defer db.Close()

	sequenceRepo := sequence.NewRepo(db)

	//devemos converter os dados da camada de apresentação para a camada de data
	sequenceData := sequenceRepo.ConvertSequenceToData(req.Latters)

	// devemos verificar se a matriz já foi testada
	exists, err := sequenceRepo.CheckSequence(sequenceData)
	if err != nil {
		return isValid, err
	}

	if *exists {
		return isValid, fmt.Errorf("%v: Essa matriz já foi validada", messageDefaultError)
	}

	// devemos validar o tamanho da matriz
	yMatriz := len(req.Latters)
	if yMatriz == 0 {
		return isValid, fmt.Errorf("%v: Matriz sem linhas", messageDefaultError)
	}

	// buscamos a quantidade de letras que foi enviada na linha 0 da matriz e ignoramos os espaços digitados acidentalmente
	xMatriz := len(strings.ReplaceAll(req.Latters[0], " ", ""))

	if xMatriz < 4 || yMatriz < 4 {
		return isValid, fmt.Errorf("%v: Matriz de tamanho inválido", messageDefaultError)
	}

	if xMatriz != yMatriz {
		return isValid, fmt.Errorf("%v: A Matriz deve ser NxN", messageDefaultError)
	}

	// devemos tratar os dados, validar o tamanho da linhas da matriz e suas letras
	matriz := make([][]string, xMatriz)

	for i := 0; i < yMatriz; i++ {
		strTemp := strings.ToUpper(req.Latters[i])
		strTemp = strings.ReplaceAll(strTemp, " ", "")
		matriz[i] = strings.Split(strTemp, "")

		if len(matriz[i]) != xMatriz {
			return isValid, fmt.Errorf("%v: Todas as linhas devem ter o mesmo número de letras", messageDefaultError)
		}

		for j := 0; j < len(matriz[i]); j++ {
			valueTemp := matriz[i][j]
			if valueTemp != "B" && valueTemp != "U" && valueTemp != "D" && valueTemp != "H" {
				return isValid, fmt.Errorf("%v: Somente as letras B, U, D e H são permitidas", messageDefaultError)
			}
		}
	}

	// buscamos a quantidade de sequências válidas
	countSequences := 0
	countSequences += checkHorizontalSequence(matriz)
	countSequences += checkVerticalSequence(matriz)
	countSequences += checkMainDiagonalSequence(matriz)
	countSequences += checkSecondaryDiagonalSequence(matriz)

	if countSequences >= 2 {
		isValid = true
	}

	// salvamos no banco a matriz testada
	sequenceData.IsValid = isValid
	err = sequenceRepo.InsertSequence(sequenceData)
	if err != nil {
		return isValid, err
	}

	return isValid, nil
}

// GetStats contém a lógica de negócio para buscar o status de sequências no banco de dados
func GetStats(c *gin.Context) (stats *Stats, err error) {
	messageDefaultError := "Não possível buscar stats"
	ratio := 0.0

	db := database.Conectar()
	defer db.Close()

	sequenceRepo := sequence.NewRepo(db)

	dataStats, err := sequenceRepo.GetStats()
	if err != nil {
		return nil, fmt.Errorf("%v: %v", messageDefaultError, err)
	}

	//devemos fazer o cálculo da porcentagem de matrizes válidas
	if *dataStats.CountInvalid+*dataStats.CountValid > 0 {
		ratio = float64(*dataStats.CountValid) / (float64(*dataStats.CountInvalid) + float64(*dataStats.CountValid))
	}

	// devemos arredondar a porcentagem para o valor mais próximo
	ratio = math.Round(ratio*100) / 100

	// convertemos os dados da data para camada de apresentação
	stats = &Stats{
		CountValid:   dataStats.CountValid,
		CountInvalid: dataStats.CountInvalid,
		Ratio:        &ratio,
	}

	return
}

// checkSecondaryDiagonalSequence procura sequencias de letras iguais em uma matriz de forma horizontal
func checkHorizontalSequence(matriz [][]string) int {
	rest := 3
	sequenceFound := 0
	var valueTemp string

	for i := 0; i < len(matriz); i++ {
		valueTemp = matriz[i][0]
		for j := 1; j < len(matriz); j++ {
			if len(matriz)-j < rest {
				break
			}

			if matriz[i][j] == valueTemp {
				rest--
				if rest == 0 {
					sequenceFound++
					rest = 3
				}
			} else {
				valueTemp = matriz[i][j]
				rest = 3
			}
		}
	}
	return sequenceFound
}

// checkSecondaryDiagonalSequence procura sequencias de letras iguais em uma matriz de forma vertical
func checkVerticalSequence(matriz [][]string) int {
	rest := 3
	sequenceFound := 0
	var valueTemp string

	for i := 0; i < len(matriz); i++ {
		valueTemp = matriz[0][i]
		for j := 1; j < len(matriz); j++ {
			if len(matriz)-j < rest {
				break
			}
			if matriz[j][i] == valueTemp {
				rest--
				if rest == 0 {
					sequenceFound++
					rest = 3
				}
			} else {
				valueTemp = matriz[j][i]
				rest = 3
			}
		}
	}

	return sequenceFound
}

// checkSecondaryDiagonalSequence procura sequencias de letras iguais em uma matriz de forma diagonal principal
func checkMainDiagonalSequence(matriz [][]string) int {
	i := len(matriz) - 1
	j := 0

	rest := 3
	sequence := 0

	for {
		if j == len(matriz) && i == 0 {
			break
		}
		iTempo := i
		jTempo := j

		valueTemp := matriz[i][j]

		for iTempo+1 < len(matriz) && jTempo+1 < len(matriz) {
			if len(matriz)-iTempo < rest || len(matriz)-jTempo < rest {
				break
			}
			if valueTemp == matriz[iTempo+1][jTempo+1] {
				rest--
				if rest == 0 {
					sequence++
					rest = 3
				}
			} else {
				valueTemp = matriz[iTempo][jTempo]
				rest = 3
			}
			iTempo++
			jTempo++
		}
		if i > 0 {
			i--
			j = 0
			continue
		}
		j++

	}
	return sequence
}

// checkSecondaryDiagonalSequence procura sequencias de letras iguais em uma matriz de forma diagonal secundária
func checkSecondaryDiagonalSequence(matriz [][]string) int {
	i := len(matriz) - 1
	j := len(matriz) - 1

	rest := 3
	sequence := 0

	for {
		if j == 0 && i == 0 {
			break
		}
		iTempo := i
		jTempo := j

		valueTemp := matriz[i][j]

		for iTempo < len(matriz) && jTempo > 0 {
			if len(matriz)-iTempo < rest || jTempo < rest {
				break
			}

			if iTempo+1 < len(matriz) && jTempo-1 >= 0 {
				if valueTemp == matriz[iTempo+1][jTempo-1] {
					rest--
					if rest == 0 {
						sequence++
						rest = 3
					}
				} else {
					valueTemp = matriz[iTempo+1][jTempo-1]
					rest = 3
				}
			}
			iTempo++
			jTempo--
		}
		if i > 0 {
			i--
			j = 5
			continue
		}
		j--

	}
	return sequence
}
