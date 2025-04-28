package biz

import (
	"context"
	"main.go/modules/rent/model"
)

func (biz *RentCommon) NewListRent(ctx context.Context, userId int) (*[]model.Rent, error) {
	if userId == 0 {
		return nil, nil
	}
	data, err := biz.store.ListRent(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}
	return data, nil
}
