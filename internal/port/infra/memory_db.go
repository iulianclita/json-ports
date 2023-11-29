package infra

import (
	"context"
	"sync"

	"github.com/iulianclita/json-ports/internal/port/domain"
)

type MemoryDB struct {
	mutex        sync.Mutex
	storagePorts map[string]*domain.Port
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		storagePorts: make(map[string]*domain.Port),
	}
}

func (mdb *MemoryDB) UpsertPort(_ context.Context, port *domain.Port) error {
	mdb.mutex.Lock()
	mdb.storagePorts[port.ID] = port
	mdb.mutex.Unlock()

	return nil
}

func (mdb *MemoryDB) GetPorts(_ context.Context) ([]*domain.Port, error) {
	var domainPorts []*domain.Port
	for _, storagePort := range mdb.storagePorts {
		if storagePort == nil {
			continue
		}
		mdb.mutex.Lock()
		domainPorts = append(domainPorts, storagePort)
		mdb.mutex.Unlock()
	}

	return domainPorts, nil
}
