package biz

import (
	"context"
	"errors"
	"main.go/common"
)

func (biz *RentCommon) NewDeletedRent(ctx context.Context, com *common.IdCommon) error {
	if com.Id == 0 {
		return errors.New("id request")
	}
	data, err := biz.store.FindRent(ctx, map[string]interface{}{"id": com.Id})
	if err != nil {
		return err
	}
	if data.UserId != com.UserId {
		return common.ErrPermissionRole(errors.New("you haven't right"))
	}
	if err := biz.store.DeletedRent(ctx, map[string]interface{}{"id": com.Id}); err != nil {
		return err
	}
	// update review
	//update userlikeroom
	return nil
}
