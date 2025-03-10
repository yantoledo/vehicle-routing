package controller

import (
	"github.com/coding/src/domain"
	read_csv "github.com/coding/src/infra/csv"
)

type ServiceInterface interface {
	RunVNS(nodes []*domain.Node, euclidianMap *domain.Matrix) domain.Solution
}

type CSVReader interface {
	Read(filePath string) ([]read_csv.Data, error)
}
