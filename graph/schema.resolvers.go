package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"boulangerie-api/db"
	"boulangerie-api/graph/generated"
	"boulangerie-api/graph/model"
	"context"
)

func (r *bakeryResolver) Breads(ctx context.Context, obj *db.BakeryModel) ([]db.BreadModel, error) {
	return r.Prisma.Bread.FindMany(db.Bread.ID.Equals(obj.ID)).Exec(ctx)
}

func (r *mutationResolver) AddBakery(ctx context.Context, name string) (*db.BakeryModel, error) {
	return r.Prisma.Bakery.CreateOne(db.Bakery.Name.Set(name)).Exec(ctx)
}

func (r *mutationResolver) AddBreadToBakery(ctx context.Context, bread *model.BreadInput, bakeryID string) (*db.BakeryModel, error) {
	breadModel, err := r.Prisma.Bread.CreateOne(db.Bread.Name.Set(bread.Name), db.Bread.Price.Set(bread.Price), db.Bread.Bakery.Link(db.Bakery.ID.Equals(bakeryID))).With(db.Bread.Bakery.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return breadModel.Bakery(), nil
}

func (r *mutationResolver) RemoveBakery(ctx context.Context, id string) (*db.BakeryModel, error) {
	return r.Prisma.Bakery.FindUnique(db.Bakery.ID.Equals(id)).Delete().Exec(ctx)
}

func (r *queryResolver) Bakeries(ctx context.Context) ([]db.BakeryModel, error) {
	return r.Prisma.Bakery.FindMany().Exec(ctx)
}

// Bakery returns generated.BakeryResolver implementation.
func (r *Resolver) Bakery() generated.BakeryResolver { return &bakeryResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type bakeryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
