package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/rent/model"
)

func (biz *RentCommon) NewUpdateRent(ctx context.Context, data *model.UpdateRent) error {
	rent, err := biz.store.FindRent(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}
	if data.UserId != rent.UserId {
		return common.ErrPermission(errors.New("you haven't right"))
	}
	if err := biz.store.UpdateRent(ctx, data); err != nil {
		return err
	}
	return nil
}
