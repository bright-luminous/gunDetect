package resource

import (
	"AI/graph/model"
	"AI/util"
	"context"

	"github.com/google/uuid"
)

func (op *SQLop) CaseCreate(ctx context.Context, caseInput *model.NewCase) (*model.Case, error) {
	newID := uuid.New().String()
	caseToBeAdd := model.Case{
		CaseID:     newID,
		CaseDate:   caseInput.CaseDate,
		CameraID:   caseInput.CameraID,
		Image1Path: caseInput.Image1Path,
		Image2Path: caseInput.Image2Path,
		Image3Path: caseInput.Image3Path,
		Image4Path: caseInput.Image4Path,
		Status:     nil,
		Respond:    nil,
	}
	_, err := op.db.NewInsert().Model(&caseToBeAdd).Exec(ctx)
	return &caseToBeAdd, err
}

func (op *SQLop) CaseUpdate(ctx context.Context, updateInput model.CaseUpdate) (*model.Case, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultCustomer, err := op.CaseFindByID(ctx, updateInput.ID)
	return resultCustomer, err
}

func (op *SQLop) CaseUpdateRespond(ctx context.Context, ID string, updateInput bool) (*model.Case, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(new(model.Case)).Set("Respond= ?", &updateInput).Set("Status=true").Where("case_id = ?", ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultCustomer, err := op.CaseFindByID(ctx, ID)
	return resultCustomer, err
}

func (op *SQLop) CaseDelete(ctx context.Context, ID string) (*model.Case, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultCustomer, err := op.CaseFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.caseModel).Where("id = ?", ID).Exec(ctx)
	return resultCustomer, err
}

func (op *SQLop) CaseDeleteAll(ctx context.Context) ([]*model.Case, error) {
	carArr, err := op.Cases(ctx)
	for _, v := range carArr {
		_, err = op.CameraDelete(ctx, v.CaseID)
	}
	return carArr, err
}

func (op *SQLop) CaseFindByID(ctx context.Context, ID string) (*model.Case, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.Case)
	err = op.db.NewSelect().Model(op.caseModel).Where("case_id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) Cases(ctx context.Context) ([]*model.Case, error) {
	car := new([]*model.Case)
	err := op.db.NewSelect().Model(op.caseModel).Scan(ctx, car)
	return *car, err
}

func (op *SQLop) CaseByResponseNull(ctx context.Context) ([]*model.FrontEndCase, error) {
	arrModel := new([]*model.FrontEndCase)
	err := op.db.NewSelect().Table("cases").
		Join("JOIN cameras ON cases.camera_id=cameras.camera_id").
		// JoinOn("cas.camera_id=cam.camera_id").
		Where("cases.respond IS NULL").
		Scan(ctx, arrModel)
	return *arrModel, err
}

func (op *SQLop) CaseByResponseNotNull(ctx context.Context) ([]*model.FrontEndCase, error) {
	arrModel := new([]*model.FrontEndCase)
	err := op.db.NewSelect().Table("cases").
		Join("JOIN cameras ON cases.camera_id=cameras.camera_id").
		// JoinOn("cas.camera_id=cam.camera_id").
		Where("cases.respond IS NOT NULL").
		Scan(ctx, arrModel)
	return *arrModel, err
}
