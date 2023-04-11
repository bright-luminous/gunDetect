package resource

import (
	model "AI/graph/model"
	"AI/util"
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func checkErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func ValidateID(sIn string) error {
	if len(sIn) < 1 {
		return errors.New("ID must not be empty")
	}
	return nil
}

func NewDBOperator(inDSN string) (*SQLop, error) {
	db, err := sql.Open("postgres", inDSN)
	bunDB := bun.NewDB(db, pgdialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	return &SQLop{
		db: bunDB,
	}, err
}

// TABLE
func (op *SQLop) DropTable(ctx context.Context) (bool, error) {
	//ticket service
	_, err := op.db.NewDropTable().
		Model((*model.Case)(nil)).
		IfExists().
		Exec(ctx)
	if checkErr(err) {
		return false, err
	}
	//active ticket
	_, err = op.db.NewDropTable().
		Model((*model.Camera)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	return true, nil
}

func (op *SQLop) CreateTables(ctx context.Context) (bool, error) {
	// customer
	_, err := op.db.NewCreateTable().
		Model((*model.Camera)(nil)).
		Exec(ctx)
	if checkErr(err) {
		return false, err
	}
	//car
	_, err = op.db.NewCreateTable().
		Model((*model.Case)(nil)).
		ForeignKey(`("camera_id") REFERENCES "cameras" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if checkErr(err) {
		return false, err
	}

	return true, err
}
