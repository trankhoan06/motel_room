package biz

import (
	"context"
	"main.go/modules/rent/model"
)

type RentBiz interface {
	CreateRent(ctx context.Context, data *model.CreateRent) error
	UpdateRent(ctx context.Context, data *model.UpdateRent) error
	FindRent(ctx context.Context, cond map[string]interface{}) (*model.Rent, error)
	DeletedRent(ctx context.Context, cond map[string]interface{}) error
	ListRent(ctx context.Context, cond map[string]interface{}) (*[]model.Rent, error)
}
type RentCommon struct {
	store RentBiz
}

func NewRentCommon(store RentBiz) *RentCommon {
	return &RentCommon{
		store: store}
}
