package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"AI/graph/model"
	"context"
)

// DropTable is the resolver for the dropTable field.
func (r *mutationResolver) DropTable(ctx context.Context) (bool, error) {
	_, err := r.DB.DropTable(ctx)
	if err != nil {
		return false, err
	}
	return true, err
}

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context) (bool, error) {
	_, err := r.DB.CreateTables(ctx)
	if err != nil {
		return false, err
	}
	return true, err
}

// CameraCreate is the resolver for the cameraCreate field.
func (r *mutationResolver) CameraCreate(ctx context.Context, input model.NewCamera) (*model.Camera, error) {
	returnCamera, err := r.DB.CameraCreate(ctx, &input)
	if err != nil {
		return nil, err
	}
	return returnCamera, err
}

// CameraUpdate is the resolver for the cameraUpdate field.
func (r *mutationResolver) CameraUpdate(ctx context.Context, input model.CameraUpdate) (*model.Camera, error) {
	returnCamera, err := r.DB.CameraUpdate(ctx, input)
	if err != nil {
		return nil, err
	}
	return returnCamera, err
}

// CameraDelete is the resolver for the cameraDelete field.
func (r *mutationResolver) CameraDelete(ctx context.Context, input string) (*model.Camera, error) {
	returnCamera, err := r.DB.CameraDelete(ctx, input)
	if err != nil {
		return nil, err
	}
	return returnCamera, err
}

// CameraDeleteAll is the resolver for the cameraDeleteAll field.
func (r *mutationResolver) CameraDeleteAll(ctx context.Context) ([]*model.Camera, error) {
	returnCameras, err := r.DB.CameraDeleteAll(ctx)
	if err != nil {
		return nil, err
	}
	return returnCameras, err
}

// CaseCreate is the resolver for the caseCreate field.
func (r *mutationResolver) CaseCreate(ctx context.Context, input model.NewCase) (*model.Case, error) {
	returnCase, err := r.DB.CaseCreate(ctx, &input)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseUpdate is the resolver for the caseUpdate field.
func (r *mutationResolver) CaseUpdate(ctx context.Context, input model.CaseUpdate) (*model.Case, error) {
	returnCase, err := r.DB.CaseUpdate(ctx, input)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseUpdateRespond is the resolver for the caseUpdateRespond field.
func (r *mutationResolver) CaseUpdateRespond(ctx context.Context, input model.CaseUpdateRespondInput) (*model.Case, error) {
	returnCase, err := r.DB.CaseUpdateRespond(ctx, input.ID, input.Respond)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseDelete is the resolver for the caseDelete field.
func (r *mutationResolver) CaseDelete(ctx context.Context, input string) (*model.Case, error) {
	returnCase, err := r.DB.CaseDelete(ctx, input)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseDeleteAll is the resolver for the caseDeleteAll field.
func (r *mutationResolver) CaseDeleteAll(ctx context.Context) ([]*model.Case, error) {
	returnCase, err := r.DB.CaseDeleteAll(ctx)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CameraByID is the resolver for the cameraByID field.
func (r *queryResolver) CameraByID(ctx context.Context, input string) (*model.Camera, error) {
	returnCase, err := r.DB.CameraFindByID(ctx, input)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// Cameras is the resolver for the cameras field.
func (r *queryResolver) Cameras(ctx context.Context) ([]*model.Camera, error) {
	returnCase, err := r.DB.Cameras(ctx)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseByID is the resolver for the caseByID field.
func (r *queryResolver) CaseByID(ctx context.Context, input string) (*model.Case, error) {
	returnCase, err := r.DB.CaseFindByID(ctx, input)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// Cases is the resolver for the cases field.
func (r *queryResolver) Cases(ctx context.Context) ([]*model.Case, error) {
	returnCase, err := r.DB.Cases(ctx)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseByResponseNull is the resolver for the CaseByResponseNull field.
func (r *queryResolver) CaseByResponseNull(ctx context.Context) ([]*model.FrontEndCase, error) {
	returnCase, err := r.DB.CaseByResponseNull(ctx)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// CaseByResponseNotNull is the resolver for the CaseByResponseNotNull field.
func (r *queryResolver) CaseByResponseNotNull(ctx context.Context) ([]*model.FrontEndCase, error) {
	returnCase, err := r.DB.CaseByResponseNotNull(ctx)
	if err != nil {
		return nil, err
	}
	return returnCase, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
