package elasticsearch

import (
	"github.com/usmhk/goserver/internal/apiserver/store"
	"sync"
)

type datasource struct {
}

func (ds *datasource) Close() error {

	return nil
}

var (
	esFactor store.Factory
	once     sync.Once
)

// GetElasticSearchFactoryOr create elasticsearch factory with the give config.
func GetElasticSearchFactoryOr() (store.Factory, error) {

	return nil, nil
}
