package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type competidor struct {
	pontos  int
	esforco int
	vencido bool
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	var competidores []competidor
	for i := 0; i < n; i++ {
		var p, e int
		fmt.Scanf("%d %d", &p, &e)

		competidores = append(competidores, competidor{pontos: p, esforco: e})
	}

	// DEBUG
	// k = 1
	// competidores = []competidor{
	// 	competidor{3, 2, false},
	// 	competidor{4, 0, false},
	// }
	// fmt.Printf("Para se classificar, Sor Ducan deve se classificar entre os %d primeiros.\n", k)

	esforco := math.MaxInt32
	i := 0
	total := len(competidores)
	combinacoes := int(math.Pow(2, float64(total)))
	fmt.Printf("%d\n", forcaBrutaRecursivo(competidores, k, i, esforco, combinacoes, total))
}

func forcaBrutaRecursivo(competidores []competidor, k, i, esforco, combinacoes, total int) int {

	if i >= combinacoes {
		if esforco == math.MaxInt32 {
			return -1
		}
		return esforco
	}

	binario := strconv.FormatInt(int64(i), 2)
	binario = fmt.Sprintf("%0*s", total, binario)

	esforcoTotal := 0
	pontosTotal := 0

	competidoresSituacao := make([]competidor, len(competidores))
	copy(competidoresSituacao, competidores)

	for j := 0; j < total; j++ {
		if binario[j] != '0' {
			pontosTotal++
			esforcoTotal += competidores[j].esforco
			competidoresSituacao[j].vencido = true
		} else {
			competidoresSituacao[j].pontos++
		}
	}

	sort.Slice(competidoresSituacao, func(i, j int) bool {
		if competidoresSituacao[i].pontos > competidoresSituacao[j].pontos {
			return true
		}
		if competidoresSituacao[i].pontos < competidoresSituacao[j].pontos {
			return false
		}
		return !competidoresSituacao[i].vencido && competidoresSituacao[j].vencido
	})

	if ((pontosTotal > competidoresSituacao[k-1].pontos) ||
		(pontosTotal == competidoresSituacao[k-1].pontos && competidoresSituacao[k-1].vencido)) && esforcoTotal < esforco {
		esforco = esforcoTotal
	}

	// fmt.Printf("%s - Pontos: %d Esforço: %d\n", binario, pontosTotal, esforcoTotal)
	// for i, cs := range competidoresSituacao {
	// 	fmt.Printf("%dº Lugar - %d %t\t", i, cs.pontos, cs.vencido)
	// }
	// fmt.Println()

	return forcaBrutaRecursivo(competidores, k, i+1, esforco, combinacoes, total)
}
