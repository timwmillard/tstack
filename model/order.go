package model

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderPending    OrderStatus = "pending"
	OrderProcessing OrderStatus = "processing"
	OrderComplete   OrderStatus = "complete"
	OrderExpired    OrderStatus = "expired"
)

type Order struct {
	ID         uuid.UUID   `db:"id"`
	Status     OrderStatus `db:"status"`
	CustomerID *uuid.UUID  `db:"customer_id"`
	Customer   Customer
	Items      []OrderItem `db:"order_items"`
}

type OrderItem struct {
	ID          uuid.UUID       `db:"id"`
	Dexcription string          `db:"description"`
	Qty         int             `db:"qty"`
	Price       decimal.Decimal `db:"price"`
}

type Customer struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
}

const createOrderSQL = `
	insert into app.order
	default values
	returning id, status
`

func createOrder(ctx context.Context, db DB) (Order, error) {
	row, _ := db.Query(ctx, createOrderSQL)
	order, err := pgx.CollectOneRow[Order](row, pgx.RowToStructByNameLax)
	return order, err
}

const createFullOrderSQL = `
	with insert_order as (	

		insert into app.order
		default values
		returning id 

	) , new_order_item as (

		select id as order_id, item, 5.00 as price -- TODO: pass price into queury
		from insert_order
		cross join unnest($1::text[]) as item

	), insert_order_item as (
		insert into app.order_item
		(order_id, description, price)
		select * from new_order_item
		returning *
	)
	select id
	from insert_order
`

func createFullOrder(ctx context.Context, db DB, items []string) (Order, error) {
	row, _ := db.Query(ctx, createFullOrderSQL, items)
	orderID, err := pgx.CollectOneRow[uuid.UUID](row, pgx.RowTo)
	if err != nil {
		return Order{}, err
	}
	return getOrderWithItems(ctx, db, orderID)
}

const getOrderWithItemsSQL = `
	select
		ord.id,
		ord.status,
		(
			select array_agg(row(
				item.id,
				item.description,
				item.qty,
				-- item.price
			))
			from app.order_item item
			where item.order_id = ord.id
		) as order_items
	from app.order ord
	where ord.id = $1
`

func getOrderWithItems(ctx context.Context, db DB, id uuid.UUID) (Order, error) {
	rows, _ := db.Query(ctx, getOrderWithItemsSQL, id)
	order, err := pgx.CollectOneRow[Order](rows, pgx.RowToStructByNameLax)
	return order, err
}

const updateOrderSQL = `
	update app.order
	set status = $2,
		customer_id = $3
	where id = $1
	returning id, status, customer_id
`

func updateOrder(ctx context.Context, db DB, orderID uuid.UUID, status OrderStatus, customerID uuid.UUID) (Order, error) {
	custID := uuid.NullUUID{
		UUID:  customerID,
		Valid: !customerID.IsNil(),
	}
	row, _ := db.Query(ctx, updateOrderSQL, orderID, status, custID)
	order, err := pgx.CollectOneRow[Order](row, pgx.RowToStructByNameLax)
	return order, err
}

const deleteExpiredOrdersSQL = `
	delete from app.order
	where status = 'expired' and expired_on + '7 day'::interval < now();
`

func deleteExpiredOrders(ctx context.Context, db DB) (int, error) {
	ct, err := db.Exec(ctx, deleteExpiredOrdersSQL)
	return int(ct.RowsAffected()), err
}

const orderCountSQL = `
	select count(*)
	from app.order
	where status = 'complete'
`

func GetOrderCount(ctx context.Context, db DB) (int, error) {
	row, _ := db.Query(ctx, orderCountSQL)
	orderCount, err := pgx.CollectOneRow[int](row, pgx.RowTo)
	return orderCount, err
}

const createCustomerSQL = `
	insert into app.customer (
		first_name,
		last_name,
		email
	) values (
		$1, $2, $3
	) returning id, first_name, last_name, email
`

func createCustomer(ctx context.Context, db DB, name, email string) (Customer, error) {
	firstName, lastName := splitName(name)
	row, _ := db.Query(ctx, createCustomerSQL,
		firstName,
		lastName,
		email,
	)
	customer, err := pgx.CollectOneRow[Customer](row, pgx.RowToStructByNameLax)
	return customer, err
}
