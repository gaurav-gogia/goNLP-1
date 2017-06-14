package hmm

import "github.com/cocaer/goNLP/data"
import "fmt"

const (
	B = iota
	E
	M
	S
	SUM_STATUS
)

type Model struct {
	StartPro [SUM_STATUS]float64
	TransPro *[SUM_STATUS][SUM_STATUS]float64
	EmitPro  *[SUM_STATUS]map[rune]float64
}

func (self *Model) Viterbi(str string) {
	ssrune := []rune(str)
	strLen := len(ssrune)
	var weight [][]float64
	var path [][]int
	for i := 0; i < SUM_STATUS; i++ {
		weight = append(weight, make([]float64, len(ssrune)))
		path = append(path, make([]int, len(ssrune)))
	}

	for i := 0; i < SUM_STATUS; i++ {
		weight[i][0] = self.StartPro[i] + self.EmitPro[i][ssrune[0]]
	}
	for i := 1; i < strLen; i++ {
		for j := 0; j < SUM_STATUS; j++ {
			weight[j][i] = IMPOSSIBLEPRO
			path[j][i] = j
			for k := 0; k < SUM_STATUS; k++ {
				tmp := weight[k][i-1] + self.TransPro[k][j] + self.EmitPro[j][ssrune[i]]

				if tmp > weight[j][i] {
					weight[j][i] = tmp
					path[j][i] = k
				}
			}
		}
	}
	result := ""
	status := SUM_STATUS - 2
	if weight[status][strLen-1] < weight[SUM_STATUS-1][strLen-1] {
		status = SUM_STATUS - 1
	}
	result = result + string(am[status])
	for i := strLen - 1; i > 0; i-- {
		result = string(am[path[status][i]]) + result
		status = path[status][i]
	}
	fmt.Println(result)
}

func NewModel() *Model {
	m := new(Model)
	m.StartPro = data.StartProMaterix
	m.EmitPro = data.EmitProMaterix
	m.TransPro = &data.TransferMatrix
	return m
}
