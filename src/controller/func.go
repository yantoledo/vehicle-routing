package controller

import (
	"math"

	"github.com/coding/src/domain"
	read_csv "github.com/coding/src/infra/csv"
)

func NewNodes(Data []read_csv.Data) []*domain.Node {
	var nodes []*domain.Node
	var isDepot bool

	for idx, data := range Data {

		if idx == 0 {
			isDepot = true
		} else {
			isDepot = false
		}

		nodes = append(nodes, &domain.Node{
			ID:      data.CustomerID,
			Demand:  data.Demand,
			IsDepot: isDepot,
		})
	}

	return nodes
}

func GenerateEuclidianDistancesMap(Data []read_csv.Data, dimension int) *domain.Matrix {

	matrix := domain.NewMatrix(dimension, dimension)

	for i := 0; i < len(Data); i++ {

		for j := i; j < len(Data); j++ {

			matrix.Set(Data[i].CustomerID, Data[j].CustomerID, math.Sqrt(math.Pow(float64(Data[j].CordinateX)-float64(Data[i].CordinateX), 2)+math.Pow(float64(Data[j].CordinateY)-float64(Data[i].CordinateY), 2)))
		}
	}

	return matrix
}
