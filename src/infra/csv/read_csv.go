package read_csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	CustomerID int
	CordinateX int
	CordinateY int
	Demand     int
}

type csvReader struct {
}

func NewCSVReader() csvReader {
	return csvReader{}
}

func (c csvReader) Read(filePath string) ([]Data, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return []Data{}, fmt.Errorf("ReadCSV unable to read input file: %w", err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []Data{}, fmt.Errorf("ReadCSV unable to parse file as CSV: %w", err)
	}

	var DataList []Data

	for _, record := range records {

		customerID, err := strconv.Atoi(string(record[0]))
		if err != nil {
			return []Data{}, fmt.Errorf("ReadCSV fails to parse cordinate customerID: %w", err)
		}

		cordinateX, err := strconv.Atoi(string(record[1]))
		if err != nil {
			return []Data{}, fmt.Errorf("ReadCSV fails to parse cordinate X: %w", err)
		}

		cordinateY, err := strconv.Atoi(string(record[2]))
		if err != nil {
			return []Data{}, fmt.Errorf("ReadCSV fails to parse cordinate Y: %w", err)
		}

		demand, err := strconv.Atoi(string(record[3]))
		if err != nil {
			return []Data{}, fmt.Errorf("ReadCSV fails to parse demand: %w", err)
		}

		DataList = append(DataList, Data{
			CustomerID: customerID,
			CordinateX: cordinateX,
			CordinateY: cordinateY,
			Demand:     demand,
		})
	}

	return DataList, nil
}
