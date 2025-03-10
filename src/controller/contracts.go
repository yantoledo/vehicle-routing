package controller

import (
	"github.com/vehicle-routing/src/domain"
	read_csv "github.com/vehicle-routing/src/infra/csv"
)

type ServiceInterface interface {
	RunVNS(nodes []*domain.Node, euclidianMap *domain.Matrix) domain.Solution
}

type CSVReader interface {
	Read(filePath string) ([]read_csv.Data, error)
}
