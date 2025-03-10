package controller

import (
	"errors"
	"fmt"
)

type Pn16k8Controller struct {
	service   ServiceInterface
	csvReader CSVReader
}

func NewPn16k8Controller(service ServiceInterface, csvReader CSVReader) (Pn16k8Controller, error) {
	if service == nil {
		return Pn16k8Controller{}, errors.New("NewPn16k8Controller missing service dependency")
	}

	if csvReader == nil {
		return Pn16k8Controller{}, errors.New("NewPn16k8Controller missing csvReader dependency")
	}

	return Pn16k8Controller{
		service:   service,
		csvReader: csvReader,
	}, nil
}

func (pc Pn16k8Controller) Run() error {

	dataList, err := pc.csvReader.Read("./data/p-n16-k8.csv")
	if err != nil {
		return errors.New("Pn16k8Controller.Run failed to read csv")
	}

	nodes := NewNodes(dataList)

	euclidianMap := GenerateEuclidianDistancesMap(dataList, 17)

	solution := pc.service.RunVNS(nodes, euclidianMap)

	for idx, car := range solution.Cars {

		fmt.Println("car index")
		fmt.Println(idx)
		fmt.Println("current demand")
		fmt.Println(car.CurrentDemand)

		fmt.Println("customers relation")
		for key, value := range car.Route.NodesMap {
			fmt.Println(key, value.ID)
		}
		fmt.Println("---------------------")
	}

	return nil
}
