package dbservice

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-demo-svc/model"
	"log"
)

type PgxIface interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
}

type DbService struct {
	PgxIface
	RowMapper
}

const (
	schema = `
CREATE TABLE IF NOT EXISTS order_ (
    order_uid char(36)  UNIQUE NOT NULL,
    track_number varchar  UNIQUE NOT NULL,
    entry varchar   NOT NULL,
    delivery_uid char(36)  UNIQUE NOT NULL,
    locale char(2)   NOT NULL,
    internal_signature varchar   NOT NULL,
    customer_id varchar   NOT NULL,
    delivery_service varchar   NOT NULL,
    shardkey varchar   NOT NULL,
    sm_id int   NOT NULL,
    date_created varchar NOT NULL,
    oof_shard varchar   NOT NULL,
    CONSTRAINT pk_order PRIMARY KEY (
        order_uid
    )
);

CREATE TABLE IF NOT EXISTS delivery (
    delivery_uid char(36)  UNIQUE NOT NULL,
    name varchar   NOT NULL,
    phone char(11)   NOT NULL,
    zip varchar   NOT NULL,
    city varchar   NOT NULL,
    address varchar   NOT NULL,
    region varchar   NOT NULL,
    email varchar   NOT NULL,
    CONSTRAINT pk_delivery PRIMARY KEY (
        delivery_uid
    ),
    CONSTRAINT fk_order_delivery_uid FOREIGN KEY(
    	delivery_uid
    ) REFERENCES order_ (delivery_uid)
);

CREATE TABLE IF NOT EXISTS payment (
    transaction char(36)  UNIQUE NOT NULL,
    request_id varchar   NOT NULL,
    currency char(3)   NOT NULL,
    provider varchar   NOT NULL,
    payment_dt bigint   NOT NULL,
    bank varchar   NOT NULL,
    delivery_cost int   NOT NULL,
    custom_fee bigint   NOT NULL,
    CONSTRAINT pk_payment PRIMARY KEY (
        transaction
    ),
    CONSTRAINT fk_payment_order_uid FOREIGN KEY(
    	transaction
    ) REFERENCES order_ (order_uid)
);

CREATE TABLE IF NOT EXISTS item (
    chrt_id bigint  UNIQUE NOT NULL,
    price bigint   NOT NULL,
    rid char(36)   NOT NULL,
    name varchar   NOT NULL,
    sale int   NOT NULL,
    size varchar   NOT NULL,
    nm_id bigint   NOT NULL,
    brand varchar   NOT NULL,
    status int   NOT NULL,
    CONSTRAINT pk_item PRIMARY KEY (
        chrt_id
    )
);

CREATE TABLE IF NOT EXISTS track (
    track_number varchar   NOT NULL,
    item_uid bigint   NOT NULL,
    CONSTRAINT pk_track PRIMARY KEY (
        track_number,
		item_uid
    ),
    CONSTRAINT fk_order_track_number FOREIGN KEY(
    	track_number
    ) REFERENCES order_ (track_number),
    CONSTRAINT fk_item_chrt_id FOREIGN KEY(
    	item_uid
    ) REFERENCES item (chrt_id)
);
`
)

func NewDbService() (*DbService, error) {
	conn, err := pgxpool.New(context.Background(), "postgres://testuser:@localhost:5432/shareit") //os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, err
	}

	mapper := NewMapper()

	return &DbService{conn, mapper}, nil
}

func (ds *DbService) AddOrder(o model.Order) error {

	batch := &pgx.Batch{}
	d := o.Delivery
	p := o.Payment
	newUuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return err
	}

	query := fmt.Sprintf(`
INSERT INTO order_(order_uid, track_number, entry, delivery_uid, locale, internal_signature,
                   customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) ON CONFLICT DO NOTHING 
`)
	batch.Queue(query, o.OrderUID, o.TrackNumber, o.Entry, newUuid.String(), o.Locale, o.InternalSignature,
		o.CustomerID, o.DeliveryService, o.Shardkey, o.SmID, o.DateCreated, o.OofShard)

	query = fmt.Sprintf(`
INSERT INTO delivery(delivery_uid, name, phone, zip, city, address, region, email)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING 
`)
	batch.Queue(query, newUuid.String(), d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email)

	query = fmt.Sprintf(`
INSERT INTO payment(transaction, request_id, currency, provider, payment_dt, bank, delivery_cost, custom_fee)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT DO NOTHING 
`)
	batch.Queue(query, p.Transaction, p.RequestID, p.Currency, p.Provider, p.PaymentDt, p.Bank,
		p.DeliveryCost, p.CustomFee)

	for _, item := range o.Items {
		query = fmt.Sprintf(`
INSERT INTO item(chrt_id, price, rid, name, sale, size, nm_id, brand, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT DO NOTHING
`)
		batch.Queue(query, item.ChrtID, item.Price, item.Rid, item.Name, item.Sale, item.Size,
			item.NmID, item.Brand, item.Status)

		query = fmt.Sprintf(`
INSERT INTO track(track_number, item_uid)
VALUES ($1, $2) ON CONFLICT DO NOTHING 
`)
		batch.Queue(query, o.TrackNumber, item.ChrtID)
	}
	err = ds.SendBatch(context.Background(), batch).Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ds *DbService) RestoreCache() (map[string]model.Order, error) {
	out := make(map[string]model.Order)
	query := `
SELECT order_uid, o.track_number, entry, locale, internal_signature, customer_id, delivery_service,
       shardkey, sm_id, date_created, oof_shard,
       json_build_object(
       		'name', d.name,
       		'phone', phone,
       		'zip', zip,
       		'city', city,
       		'address', address,
       		'region', region,
       		'email', email
       ) as delivery,
       json_build_object(
       		'transaction', transaction,
       		'request_id', request_id,
       		'currency', currency,
       		'provider', provider,
       		'payment_dt', payment_dt,
       		'bank', bank,
       		'delivery_cost', delivery_cost,
       		'custom_fee', custom_fee
       ) as payment,
       json_agg(json_build_object(
			'chrt_id', chrt_id,
			'track_number', o.track_number,
			'price', price,
            'rid', rid,
            'name', i.name,
            'sale', sale,
            'size', size,
            'total_price', (100 - sale) * price/100,
            'nm_id', nm_id,
            'brand', brand,
            'status', status
	   )) as items
FROM order_ o
JOIN delivery d on o.delivery_uid = d.delivery_uid
JOIN payment p on o.order_uid = p.transaction
JOIN track t on o.track_number = t.track_number
JOIN item i on i.chrt_id = t.item_uid
GROUP BY order_uid, d.name, phone, zip, city, address, region, email, transaction, request_id, currency, provider, payment_dt, bank, delivery_cost, custom_fee
`
	rows, err := ds.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	orders, err := ds.RowsToOrders(rows)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(orders); i++ {
		goodsTotal := 0
		for _, item := range orders[i].Items {
			goodsTotal += item.TotalPrice
		}
		orders[i].Payment.Amount = goodsTotal + orders[i].Payment.DeliveryCost + orders[i].Payment.CustomFee
		orders[i].Payment.GoodsTotal = goodsTotal

		out[orders[i].OrderUID] = orders[i]
	}
	return out, nil
}
