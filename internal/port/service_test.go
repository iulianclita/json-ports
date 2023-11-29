package port_test

import (
	"context"
	"testing"

	"github.com/iulianclita/json-ports/internal/port"
	"github.com/iulianclita/json-ports/internal/port/domain"
	"github.com/iulianclita/json-ports/internal/port/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpsertPort(t *testing.T) {
	memDB := infra.NewMemoryDB()
	service := port.NewService(memDB)

	t.Run("create new port", func(t *testing.T) {
		testPort := domain.Port{
			ID:          "123-abc",
			City:        "Bucharest",
			Country:     "RO",
			Alias:       []string{"some-alias"},
			Regions:     []string{"some-region"},
			Coordinates: []float64{1.23, 4.56},
			Province:    "some-province",
			Timezone:    "UTC+2",
			Unlocs:      []string{"some-unloc"},
			Code:        "some-code",
		}

		errCreate := service.UpsertPort(context.Background(), &testPort)
		require.NoError(t, errCreate)
		gotCreatedPorts, err := service.GetPorts(context.Background())
		require.NoError(t, err)
		wantCreatedPorts := []*domain.Port{&testPort}
		assert.Equal(t, wantCreatedPorts, gotCreatedPorts)
	})

	t.Run("updated existing port", func(t *testing.T) {
		updatedTestPort := domain.Port{
			ID:          "123-abc",
			City:        "Sofia",
			Country:     "BG",
			Alias:       []string{"some-updated-alias"},
			Regions:     []string{"some-updated-region"},
			Coordinates: []float64{7.89, 10.00},
			Province:    "some-updated-province",
			Timezone:    "UTC+3",
			Unlocs:      []string{"some-updated-unloc"},
			Code:        "some-updated-code",
		}

		errUpdate := service.UpsertPort(context.Background(), &updatedTestPort)
		require.NoError(t, errUpdate)
		gotUpdatedPorts, err := service.GetPorts(context.Background())
		require.NoError(t, err)
		wantUpdatedPorts := []*domain.Port{&updatedTestPort}
		assert.Equal(t, wantUpdatedPorts, gotUpdatedPorts)
	})
}
