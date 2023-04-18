package resource

import (
	"AI/graph/model"
	"AI/util"
	"context"

	"github.com/google/uuid"
)

func (op *SQLop) CameraCreate(ctx context.Context, cameraInput *model.NewCamera) (*model.Camera, error) {
	newID := uuid.New().String()
	carToBeAdd := model.Camera{
		ID:           newID,
		LocationName: cameraInput.LocationName,
		Location:     cameraInput.Location,
	}
	_, err := op.db.NewInsert().Model(&carToBeAdd).Exec(ctx)
	return &carToBeAdd, err
}

func (op *SQLop) CameraUpdate(ctx context.Context, updateInput model.Car) (*model.Car, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultCustomer, err := op.CarFindByID(ctx, updateInput.ID)
	return resultCustomer, err
}

func (op *SQLop) CameraDelete(ctx context.Context, ID string) (*model.Car, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultCustomer, err := op.CarFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.carModel).Where("id = ?", ID).Exec(ctx)
	return resultCustomer, err
}

func (op *SQLop) CameraDeleteAll(ctx context.Context) ([]*model.Car, error) {
	carArr, err := op.CarList(ctx)
	PrintIfErrorExist(err)
	for _, v := range carArr {
		_, err := op.CustomerDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return carArr, err
}

func (op *SQLop) CameraFindByID(ctx context.Context, ID string) (*model.Car, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.Car)
	err = op.db.NewSelect().Model(op.carModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) Cameras(ctx context.Context) ([]*model.Car, error) {
	car := new([]*model.Car)
	err := op.db.NewSelect().Model(op.carModel).Scan(ctx, car)
	return *car, err
}
