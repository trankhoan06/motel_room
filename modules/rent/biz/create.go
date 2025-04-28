package biz

import (
	"context"
	"main.go/modules/rent/model"
)

func (biz *RentCommon) NewCreateRent(ctx context.Context, data *model.CreateRent) error {
	if err := biz.store.CreateRent(ctx, data); err != nil {
		return err
	}
	return nil
}
