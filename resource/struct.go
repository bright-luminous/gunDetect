package resource

import (
	"context"

	"github.com/uptrace/bun"
)

type SQLop struct {
	db           *bun.DB
	cusModel     *model.Customer
	carModel     *model.Car
	ticketModel  *model.Ticket
	shopModel    *model.Shop
	serviceModel *model.Service

	shopService   *model.ShopService
	activeTicket  *model.ActiveTicket
	ticketService *model.TicketService
}

type DatabaseOp interface {
	DropTable(ctx context.Context) (bool, error)
	CreateTables(ctx context.Context) (bool, error)
	CreateConnectTables(ctx context.Context) (bool, error)

	CustomerCreate(ctx context.Context, userInput *model.CustomerCreateInput) (*model.Customer, error)
	CustomerUpdateMulti(ctx context.Context, updateInput model.Customer) (*model.Customer, error)
	CustomerDelete(ctx context.Context, ID string) (*model.Customer, error)
	CustomerDeleteAll(ctx context.Context) ([]*model.Customer, error)
	CustomerFindByID(ctx context.Context, ID string) (*model.Customer, error)
	CustomerFindByEmail(ctx context.Context, Email string) (*model.Customer, error)
	CustomerList(ctx context.Context) ([]*model.Customer, error)
}
