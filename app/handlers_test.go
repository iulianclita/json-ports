package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iulianclita/json-ports/internal/presentation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPortsHandler(t *testing.T) {
	// setup app
	app := New(Config{})
	handler := app.portsHandler()

	t.Run("POST", func(t *testing.T) {
		// setup request
		reqBody, mpWriter := createPayloadFromTestFile(t, "testdata/ports.json")
		req := httptest.NewRequest(http.MethodPost, "/ports", reqBody)
		req.Header.Set("Content-Type", mpWriter.FormDataContentType())
		// setup recorder
		recorder := httptest.NewRecorder()

		// call handler
		handler(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("GET", func(t *testing.T) {
		// setup request
		req := httptest.NewRequest(http.MethodGet, "/ports", nil)
		req.Header.Set("Content-Type", "application/json")
		// setup recorder
		recorder := httptest.NewRecorder()

		// call handler
		handler(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var gotPorts []presentation.ResponsePort
		err := json.NewDecoder(recorder.Body).Decode(&gotPorts)
		require.NoError(t, err)
		expectedPorts := []presentation.ResponsePort{
			{
				ID:      "AEAJM",
				Name:    "Ajman",
				City:    "Ajman",
				Country: "United Arab Emirates",
				Alias:   []string{},
				Regions: []string{},
				Coordinates: []float64{
					55.5136433,
					25.4052165,
				},
				Province: "Ajman",
				Timezone: "Asia/Dubai",
				Unlocs: []string{
					"AEAJM",
				},
				Code: "52000",
			},
		}

		assert.Equal(t, expectedPorts, gotPorts)
	})
}
