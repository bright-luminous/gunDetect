package resource

import (
	"AI/graph/model"
	"context"

	"github.com/uptrace/bun"
)

type SQLop struct {
	db          *bun.DB
	cameraModel *model.Camera
	caseModel   *model.Case
}

type DatabaseOp interface {
	DropTable(ctx context.Context) (bool, error)
	CreateTables(ctx context.Context) (bool, error)

	CameraCreate(ctx context.Context, userInput *model.NewCamera) (*model.Camera, error)
	CameraUpdate(ctx context.Context, updateInput model.CameraUpdate) (*model.Camera, error)
	CameraDelete(ctx context.Context, ID string) (*model.Camera, error)
	CameraDeleteAll(ctx context.Context) ([]*model.Camera, error)

	CameraFindByID(ctx context.Context, ID string) (*model.Camera, error)
	Cameras(ctx context.Context) ([]*model.Camera, error)

	CaseCreate(ctx context.Context, userInput *model.NewCase) (*model.Case, error)
	CaseUpdate(ctx context.Context, updateInput model.CaseUpdate) (*model.Case, error)
	CaseDelete(ctx context.Context, ID string) (*model.Case, error)
	CaseDeleteAll(ctx context.Context) ([]*model.Case, error)

	CaseFindByID(ctx context.Context, ID string) (*model.Case, error)
	Cases(ctx context.Context) ([]*model.Case, error)
	CaseByResponseNull(ctx context.Context) ([]*model.FrontEndCase, error)
	CaseByResponseNotNull(ctx context.Context) ([]*model.FrontEndCase, error)
}
