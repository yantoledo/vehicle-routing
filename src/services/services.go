package services

import (
	"errors"
	"math/rand/v2"

	"github.com/vehicle-routing/src/domain"
	"github.com/vehicle-routing/src/utils"
)

type Rand func(n int) int

type ServiceParams struct {
	MaxCarNumber int
	CarCapacity  int
	Rand         Rand
}

type Service struct {
	maxCarNumber int
	carCapacity  int
	rand         Rand
}

func NewService(serviceParams ServiceParams) (Service, error) {
	if serviceParams.MaxCarNumber == 0 {
		return Service{}, errors.New("NewService missing MaxCarNumber dependency")
	}

	if serviceParams.CarCapacity == 0 {
		return Service{}, errors.New("NewService missing CarCapacity dependency")
	}

	if serviceParams.Rand == nil {
		serviceParams.Rand = rand.IntN
	}

	return Service{
		maxCarNumber: serviceParams.MaxCarNumber,
		carCapacity:  serviceParams.CarCapacity,
		rand:         serviceParams.Rand,
	}, nil
}

func (s Service) RunVNS(nodes []*domain.Node, euclidianMap *domain.Matrix) domain.Solution {
	s.GenerateInitialSolution(nodes, euclidianMap)

	return s.GenerateInitialSolution(nodes, euclidianMap)
}

func (s Service) Shake(currentSolution domain.Solution, level int) domain.Solution {

	var idxCar1 int
	var idxCar2 int

	for {
		idxCar1 = s.rand(len(currentSolution.Cars))
		idxCar2 = s.rand(len(currentSolution.Cars))

		if idxCar1 != idxCar2 {
			break
		}
	}

	car1 := currentSolution.Cars[idxCar1]
	car2 := currentSolution.Cars[idxCar2]

	return domain.Solution{}
}

// implementar testes para este método
func (s Service) GenerateInitialSolution(nodes []*domain.Node, euclidianMap *domain.Matrix) domain.Solution {
	initialSolution := domain.Solution{}

	availableCustomers := []*domain.Node{}
	for _, node := range nodes {
		if !node.IsDepot {
			availableCustomers = append(availableCustomers, node)
		}
	}

	for k := 0; k < s.maxCarNumber; k++ {

		newCar := domain.NewCar(s.carCapacity, nodes[0])

		sequenceNodesInserted := []int{1} // depot ID is the first node inserted for all routes

		for {

			maxRandomNum := len(availableCustomers)
			if maxRandomNum == 0 {
				break
			}

			randomIndex := s.rand(maxRandomNum)

			if availableCustomers[randomIndex].Demand+newCar.CurrentDemand > newCar.Capacity {
				newCar.Route.NodesMap[sequenceNodesInserted[len(sequenceNodesInserted)-1]] = nodes[0]

				break
			}

			newCar.CurrentDemand += availableCustomers[randomIndex].Demand
			availableCustomers[randomIndex].CurrentEdge = domain.NewEdge(sequenceNodesInserted[len(sequenceNodesInserted)-1], availableCustomers[randomIndex].ID, int(euclidianMap.At(sequenceNodesInserted[len(sequenceNodesInserted)-1], availableCustomers[randomIndex].ID)))

			newCar.Route.NodesMap[sequenceNodesInserted[len(sequenceNodesInserted)-1]] = availableCustomers[randomIndex]
			sequenceNodesInserted = append(sequenceNodesInserted, availableCustomers[randomIndex].ID)

			availableCustomers = utils.RemoveNode(availableCustomers, randomIndex)
		}

		initialSolution.Cars = append(initialSolution.Cars, newCar)
	}

	if len(availableCustomers) > 0 {

		penalizedCar := s.GeneratePenalizedCar(nodes, availableCustomers, euclidianMap)

		initialSolution.Cars = append(initialSolution.Cars, penalizedCar)
	}

	return initialSolution
}

// implementar testes para este método
func (s Service) GeneratePenalizedCar(nodes []*domain.Node, availableCustomers []*domain.Node, euclidianMap *domain.Matrix) domain.Car {

	sequenceNodesInserted := []int{1} // depot ID is the first node inserted for all routes

	newPenalizedCar := domain.NewCar(s.carCapacity, nodes[0])
	newPenalizedCar.SetPenalty()

	for _, customer := range availableCustomers {

		newPenalizedCar.CurrentDemand += customer.Demand

		customer.CurrentEdge = domain.NewEdge(sequenceNodesInserted[len(sequenceNodesInserted)-1], customer.ID, int(euclidianMap.At(sequenceNodesInserted[len(sequenceNodesInserted)-1], customer.ID)))

		newPenalizedCar.Route.NodesMap[sequenceNodesInserted[len(sequenceNodesInserted)-1]] = customer

		sequenceNodesInserted = append(sequenceNodesInserted, customer.ID)
	}

	newPenalizedCar.Route.NodesMap[sequenceNodesInserted[len(sequenceNodesInserted)-1]] = nodes[0]

	return newPenalizedCar
}
