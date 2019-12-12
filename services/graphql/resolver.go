package graphql

import (
	"context"

	"github.com/rodrwan/twoppin"
	"github.com/rodrwan/twoppin/repository/database"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	Database *database.Database
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Trips(ctx context.Context, userID string) ([]*twoppin.Trips, error) {
	panic("not implemented")
}
