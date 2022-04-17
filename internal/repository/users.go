package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	api_errors "github.com/plutonio00/pay-api/internal/error"
)

type UsersRepo struct {
	db        *sql.DB
	tableName string
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db:        db,
		tableName: "users",
	}
}

func (r *UsersRepo) Replenish(userId int, sum int) error {

	balance, err := r.GetBalanceById(userId)

	if err != nil {
		return err
	}

	sum += balance
	query := fmt.Sprintf("UPDATE %s SET balance = $2 WHERE id = $1", r.tableName)

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, userId, sum)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Transfer(senderId int, recipientId int, sum int) error {
	senderBalance, err := r.GetBalanceById(senderId)

	if err != nil {
		return err
	}

	if senderBalance < sum {
		return api_errors.ErrNotEnoughMoney
	}

	recipientBalance, err := r.GetBalanceById(recipientId)

	query := fmt.Sprintf("UPDATE %s SET balance = $2 WHERE id = $1", r.tableName)

	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, senderId, senderBalance-sum)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, recipientId, recipientBalance+sum)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetBalanceById(id int) (balance int, err error) {

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", "balance", r.tableName)

	err = r.db.QueryRow(query, id).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			err = api_errors.ErrUserNotFound
		}

		return 0, err
	}

	return balance, nil
}
