package service

import (
	"errors"
	"math"
)

type PricingService interface {
	GetRetailTotal(code string, qty int) (total float64, err error)
	GetWholesaleTotal(partner string, code string, qty int) (total float64, err error)
}

type ProductRepo interface {
	FetchPrice(code string) (price float64, found bool)
	FetchDiscount(partner string) (discount float64, found bool)
}

var (
	ErrInvalidPartner  = errors.New("Invalid Partner Requested")
	ErrPartnerNotFound = errors.New("Partner Not Found")
	ErrInvalidCode     = errors.New("Invalid Code Requested")
	ErrCodeNotFound    = errors.New("Code Not Found")
	ErrInvalidQty      = errors.New("Invalid Quantity Requested")
)

type pricingService struct {
	repo ProductRepo
}

func NewPricingService(pr ProductRepo) (ps *pricingService) {
	ps = &pricingService{
		repo: pr,
	}

	return ps
}

func (ps *pricingService) GetRetailTotal(code string, qty int) (total float64, err error) {
	if code == "" {
		return 0.0, ErrInvalidCode
	}
	if qty <= 0 {
		return 0.0, ErrInvalidQty
	}

	price, found := ps.repo.FetchPrice(code)
	if !found {
		return 0.0, ErrCodeNotFound
	}

	total = price * float64(qty)

	return math.Round(total*100) / 100, nil
}

func (ps *pricingService) GetWholesaleTotal(partner string, code string, qty int) (total float64, err error) {
	if partner == "" {
		return 0.0, ErrInvalidPartner
	}
	if code == "" {
		return 0.0, ErrInvalidCode
	}
	if qty <= 0 {
		return 0.0, ErrInvalidQty
	}

	price, found := ps.repo.FetchPrice(code)
	if !found {
		return 0.0, ErrCodeNotFound
	}

	discount, found := ps.repo.FetchDiscount(partner)
	if !found {
		return 0.0, ErrPartnerNotFound
	}

	saved := (price * discount)
	total = (price - saved) * float64(qty)

	return math.Round(total*100) / 100, nil
}
