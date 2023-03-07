package elasticsearch

import (
	"context"
	elastic "github.com/olivere/elastic/v7"
)

type Client interface {
}

type IndicesExistsService interface {
	Do(ctx context.Context) (bool, error)
}

type IndicesCreateService interface {
	Body(mapping string) IndicesCreateService
	Do(cxt context.Context) (*elastic.IndicesCreateResult, error)
}
