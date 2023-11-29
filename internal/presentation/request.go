package presentation

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iulianclita/json-ports/internal/port/domain"
)

const (
	maxMemoryInMB = 100
)

type RequestPorts map[string]struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func ParseUpsertPortRequest(req *http.Request) ([]*domain.Port, error) {
	if err := req.ParseMultipartForm(maxMemoryInMB * 1024 * 1024); err != nil {
		return nil, fmt.Errorf("faild to parse multipart form upsert port request: %w", err)
	}

	file, _, err := req.FormFile("ports")
	if err != nil {
		return nil, fmt.Errorf("failed to read multipart form file in upsert port request: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(bufio.NewReader(file))

	var reqPorts RequestPorts
	if err := decoder.Decode(&reqPorts); err != nil {
		return nil, fmt.Errorf("failed to decode port: %w", err)
	}

	var domainPorts []*domain.Port
	for id, reqPort := range reqPorts {
		// skip empty id fields
		if id == "" {
			continue
		}
		domainPorts = append(domainPorts, &domain.Port{
			ID:          id,
			Name:        reqPort.Name,
			City:        reqPort.City,
			Country:     reqPort.Country,
			Alias:       reqPort.Alias,
			Regions:     reqPort.Regions,
			Coordinates: reqPort.Coordinates,
			Province:    reqPort.Province,
			Timezone:    reqPort.Timezone,
			Unlocs:      reqPort.Unlocs,
			Code:        reqPort.Code,
		})
	}

	return domainPorts, nil
}
