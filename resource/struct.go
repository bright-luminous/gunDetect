package resource

import (
	"context"

	"github.com/uptrace/bun"
)

type SQLop struct {
	db *bun.DB
}

type DatabaseOp interface {
	DropTable(ctx context.Context) (bool, error)
	CreateTables(ctx context.Context) (bool, error)

	// CameraCreate(ctx context.Context, userInput *model.NewCamera) (*model.Camera, error)
	// CameraUpdate(ctx context.Context, updateInput model.Camera) (*model.Camera, error)
	// CameraDelete(ctx context.Context, ID string) (*model.Camera, error)
	// CameraDeleteAll(ctx context.Context) ([]*model.Camera, error)
}
