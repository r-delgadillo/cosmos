package cosmoscontext

import "context"

type CosmosContext struct {
	ctx context.Context
}

type Option func(ctx *CosmosContext)

func New(ctx context.Context, opts ...Option) *CosmosContext {
	return &CosmosContext{
		ctx: context.Background(),
	}
}
