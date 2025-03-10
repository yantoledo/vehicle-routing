package controller_test

import (
	"math"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vehicle-routing/src/controller"
	"github.com/vehicle-routing/src/domain"
	read_csv "github.com/vehicle-routing/src/infra/csv"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = Describe("Func Controller", func() {

	Context("New Nodes constructions is requested", func() {

		When("nodes are created successfully", func() {

			It("must return a list of nodes", func() {

				// ARRANGE

				csvData := []read_csv.Data{
					{
						CustomerID: 1,
						CordinateX: 10,
						CordinateY: 20,
						Demand:     10,
					},
					{
						CustomerID: 2,
						CordinateX: 20,
						CordinateY: 30,
						Demand:     10,
					},
					{
						CustomerID: 3,
						CordinateX: 30,
						CordinateY: 40,
						Demand:     10,
					},
				}

				expectedNodes := []*domain.Node{
					{
						ID:          1,
						Demand:      10,
						IsDepot:     true,
						CurrentEdge: nil,
					},
					{
						ID:          2,
						Demand:      10,
						IsDepot:     false,
						CurrentEdge: nil,
					},
					{
						ID:          3,
						Demand:      10,
						IsDepot:     false,
						CurrentEdge: nil,
					},
				}

				// ACT

				nodes := controller.NewNodes(csvData)

				// ASSERT

				Expect(nodes).To(BeEquivalentTo(expectedNodes))
			})
		})
	})

	Context("New Euclidian Map is requested", func() {

		When("euclidian map are created successfully", func() {

			It("must return a map", func() {

				// ARRANGE

				csvData := []read_csv.Data{
					{
						CustomerID: 1,
						CordinateX: 10,
						CordinateY: 20,
						Demand:     10,
					},
					{
						CustomerID: 2,
						CordinateX: 20,
						CordinateY: 30,
						Demand:     10,
					},
					{
						CustomerID: 3,
						CordinateX: 30,
						CordinateY: 40,
						Demand:     10,
					},
				}

				// ACT

				euclidianMap := controller.GenerateEuclidianDistancesMap(csvData, 4)

				// ASSERT

				// rota crescente
				Expect(euclidianMap.At(1, 1)).To(BeEquivalentTo(0))
				Expect(euclidianMap.At(1, 2)).To(BeEquivalentTo(math.Sqrt(math.Pow(float64(csvData[1].CordinateX)-float64(csvData[0].CordinateX), 2) + math.Pow(float64(csvData[1].CordinateY)-float64(csvData[0].CordinateY), 2))))
				Expect(euclidianMap.At(1, 3)).To(BeEquivalentTo(math.Sqrt(math.Pow(float64(csvData[2].CordinateX)-float64(csvData[0].CordinateX), 2) + math.Pow(float64(csvData[2].CordinateY)-float64(csvData[0].CordinateY), 2))))
				Expect(euclidianMap.At(2, 2)).To(BeEquivalentTo(0))
				Expect(euclidianMap.At(2, 3)).To(BeEquivalentTo(math.Sqrt(math.Pow(float64(csvData[2].CordinateX)-float64(csvData[1].CordinateX), 2) + math.Pow(float64(csvData[2].CordinateY)-float64(csvData[1].CordinateY), 2))))
				Expect(euclidianMap.At(3, 3)).To(BeEquivalentTo(0))

				// rota decrescente
				Expect(euclidianMap.At(1, 2)).To(BeEquivalentTo(euclidianMap.At(2, 1)))
				Expect(euclidianMap.At(1, 3)).To(BeEquivalentTo(euclidianMap.At(3, 1)))
				Expect(euclidianMap.At(2, 3)).To(BeEquivalentTo(euclidianMap.At(3, 2)))
			})
		})
	})
})
