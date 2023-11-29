package presentation

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/iulianclita/json-ports/internal/domain"
)

type ResponsePort struct {
	ID          string    `json:"id"`
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

func PrepareUpsertPortResponse(domainPorts []*domain.Port) []ResponsePort {
	responsePorts := make([]ResponsePort, len(domainPorts))
	for i, domainPort := range domainPorts {
		responsePorts[i] = ResponsePort{
			ID:          domainPort.ID,
			Name:        domainPort.Name,
			City:        domainPort.City,
			Country:     domainPort.Country,
			Alias:       domainPort.Alias,
			Regions:     domainPort.Regions,
			Coordinates: domainPort.Coordinates,
			Province:    domainPort.Province,
			Timezone:    domainPort.Timezone,
			Unlocs:      domainPort.Unlocs,
			Code:        domainPort.Code,
		}
	}

	return responsePorts
}

func SendResponse(w http.ResponseWriter, res any, status int, logger *slog.Logger) {
	data, err := json.Marshal(res)
	if err != nil {
		logger.Warn("failed to marshal response data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		logger.Warn("failed to send response", err)
	}
}
