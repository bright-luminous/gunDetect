package resource

import (
	"AI/graph/model"
	"AI/util"
	"context"

	"github.com/google/uuid"
)

func (op *SQLop) CameraCreate(ctx context.Context, cameraInput *model.NewCamera) (*model.Camera, error) {
	newID := uuid.New().String()
	caseToBeAdd := model.Camera{
		CameraID:     newID,
		LocationName: cameraInput.LocationName,
		Location:     cameraInput.Location,
		Latitude:     cameraInput.Latitude,
		Longitude:    cameraInput.Longitude,
	}
	_, err := op.db.NewInsert().Model(&caseToBeAdd).Exec(ctx)
	return &caseToBeAdd, err
}

func (op *SQLop) CameraUpdate(ctx context.Context, updateInput model.CameraUpdate) (*model.Camera, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultCustomer, err := op.CameraFindByID(ctx, updateInput.ID)
	return resultCustomer, err
}

func (op *SQLop) CameraDelete(ctx context.Context, ID string) (*model.Camera, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultCustomer, err := op.CameraFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.cameraModel).Where("id = ?", ID).Exec(ctx)
	return resultCustomer, err
}

func (op *SQLop) CameraDeleteAll(ctx context.Context) ([]*model.Camera, error) {
	carArr, err := op.Cameras(ctx)
	for _, v := range carArr {
		_, err = op.CameraDelete(ctx, v.CameraID)
	}
	return carArr, err
}

func (op *SQLop) CameraFindByID(ctx context.Context, ID string) (*model.Camera, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.Camera)
	err = op.db.NewSelect().Model(op.cameraModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) Cameras(ctx context.Context) ([]*model.Camera, error) {
	car := new([]*model.Camera)
	err := op.db.NewSelect().Model(op.cameraModel).Scan(ctx, car)
	return *car, err
}
