package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrWalletNotFound    = errors.New("wallet not found")
)

type WalletRepository struct {
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) UpdateBalance(
	ctx context.Context,
	walletID uuid.UUID,
	amount int64,
	op string,
) (int64, error) {

	var query string

	switch op {
	case "DEPOSIT":
		query = `
		UPDATE wallets
		SET balance = balance + $1
		WHERE id = $2
		RETURNING balance;
		`
	case "WITHDRAW":
		query = `
		UPDATE wallets
		SET balance = balance - $1
		WHERE id = $2 AND balance >= $1
		RETURNING balance;
		`
	default:
		return 0, errors.New("invalid operation type")
	}

	var balance int64
	err := r.db.QueryRow(ctx, query, amount, walletID).Scan(&balance)

	if errors.Is(err, pgx.ErrNoRows) {
		if op == "WITHDRAW" {
			return 0, ErrInsufficientFunds
		}
		return 0, ErrWalletNotFound
	}

	return balance, err
}

func (r *WalletRepository) GetBalance(ctx context.Context, walletID uuid.UUID) (int64, error) {
	var balance int64
	err := r.db.QueryRow(ctx,
		`SELECT balance FROM wallets WHERE id=$1`, walletID).
		Scan(&balance)

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, ErrWalletNotFound
	}

	return balance, err
}
