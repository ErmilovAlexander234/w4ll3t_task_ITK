package service

import (
	"context"

	"w4ll3t/internal/repository"

	"github.com/google/uuid"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(r *repository.WalletRepository) *WalletService {
	return &WalletService{repo: r}
}

func (s *WalletService) Update(
	ctx context.Context,
	id uuid.UUID,
	amount int64,
	op string,
) (int64, error) {
	return s.repo.UpdateBalance(ctx, id, amount, op)
}

func (s *WalletService) Get(ctx context.Context, id uuid.UUID) (int64, error) {
	return s.repo.GetBalance(ctx, id)
}
