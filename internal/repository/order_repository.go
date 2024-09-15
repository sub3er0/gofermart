package repository

import (
	"gofermart/db"
	"time"
)

const (
	NEW        = "NEW"
	PROCESSING = "PROCESSED"
	INVALID    = "INVALID"
	PROCESSED  = "PROCESSED"
)

type OrderRepository struct {
	DBStorage *db.PgStorage
}

type OrderData struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func (or *OrderRepository) IsOrderExist(orderNumber string) (bool, error) {
	var id int
	query := "SELECT id FROM orders WHERE number = $1"
	err := or.DBStorage.Conn.QueryRow(or.DBStorage.Ctx, query, orderNumber).Scan(&id)

	if err != nil && err.Error() != "no rows in result set" {
		return false, err
	}

	if err != nil && err.Error() == "no rows in result set" {
		return false, nil
	}

	return true, nil
}

func (or *OrderRepository) SaveOrder(orderNumber string, userID int) error {
	query := "INSERT INTO orders (number, user_id, status) VALUES ($1, $2, $3)"
	_, err := or.DBStorage.Conn.Exec(or.DBStorage.Ctx, query, orderNumber, userID, NEW)
	return err
}

func (or *OrderRepository) GetUserOrders(userID int) ([]OrderData, error) {
	var orders []OrderData

	query := "SELECT number, status, accrual, uploaded_at FROM orders WHERE user_id = $1"
	rows, err := or.DBStorage.Conn.Query(or.DBStorage.Ctx, query, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order OrderData
		if err := rows.Scan(&order.Number, &order.Status, &order.Accrual, &order.UploadedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
