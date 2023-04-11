package resource

import (
	"AI/util"
	"context"
	"database/sql"
	"errors"

	"graph/model/model"

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
		cusModel: new(model.Customer),
		db:       bunDB,
	}, err
}

// TABLE
func (op *SQLop) DropTable(ctx context.Context) (bool, error) {
	//ticket service
	_, err := op.db.NewDropTable().
		Model((*model.TicketService)(nil)).
		IfExists().
		Exec(ctx)
	if checkErr(err) {
		return false, err
	}
	//active ticket
	_, err = op.db.NewDropTable().
		Model((*model.ActiveTicket)(nil)).
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
		Model((*model.Customer)(nil)).
		Exec(ctx)
	if checkErr(err) {
		return false, err
	}
	//car
	_, err = op.db.NewCreateTable().
		Model((*model.Car)(nil)).
		ForeignKey(`("owner_id") REFERENCES "customers" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if checkErr(err) {
		return false, err
	}

	return true, err
}
