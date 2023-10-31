package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

type ExchangeRate struct {
	Value string `json:"value"`
}

type Model struct {
	ID    string `gorm:"primaryKey"`
	Value string
	gorm.Model
}

func (s *Service) GetExchangeRate(ctx context.Context) (*ExchangeRate, error) {
	currentExchangeRate, err := GetExchangeRateAPI()
	if err != nil {
		return nil, err
	}

	exchangeRate := ExchangeRate{
		Value: currentExchangeRate.USDBRL.BID,
	}

	err = s.createExchangeRate(ctx, exchangeRate)
	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func (s *Service) createExchangeRate(ctx context.Context, e ExchangeRate) error {
	model := Model{
		ID:    uuid.New().String(),
		Value: e.Value,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	return s.db.WithContext(ctx).Create(&model).Error
}
