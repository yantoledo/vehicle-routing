package main

import (
	"github.com/coding/src/controller"
	read_csv "github.com/coding/src/infra/csv"
	"github.com/coding/src/services"
)

/*
	NAME : P-n16-k8
	COMMENT : (Augerat et al, No of trucks: 8, Optimal value: 450)
	TYPE : CVRP
	DIMENSION : 16
	EDGE_WEIGHT_TYPE : EUC_2D
	CAPACITY : 35
	NODE_COORD_SECTION
	1 30 40
	2 37 52
	3 49 49
	4 52 64
	5 31 62
	6 52 33
	7 42 41
	8 52 41
	9 57 58
	10 62 42
	11 42 57
	12 27 68
	13 43 67
	14 58 48
	15 58 27
	16 37 69
	DEMAND_SECTION
	1 0
	2 19
	3 30
	4 16
	5 23
	6 11
	7 31
	8 15
	9 28
	10 8
	11 8
	12 7
	13 14
	14 6
	15 19
	16 11
	DEPOT_SECTION
	1
	-1
	EOF
*/

const (
	MAX_ITER = 3
	MAX_TIME = 2
)

func main() {

	service, err := services.NewService(services.ServiceParams{
		MaxCarNumber: 8,
		CarCapacity:  35,
	})
	if err != nil {
		panic(err)
	}

	csvReader := read_csv.NewCSVReader()

	controller, err := controller.NewPn16k8Controller(service, csvReader)
	if err != nil {
		panic(err)
	}

	controller.Run()

}

// func Shaking() {
// 	// método para evitar ótimos locais e estimular diversidade de busca

// 	// S solução corrente
// 	// NL
// }
