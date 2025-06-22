package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/search/model"
)

func (biz *SearchCommon) NewDeletedSearch(ctx context.Context, data *model.UpdateSearch) error {
	user, err := biz.store.FindSearch(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}
	if data.UserId != user.UserId {
		return common.ErrPermission(errors.New("you can't this search, you don't owner search "))
	}
	if err := biz.store.DeletedSearch(ctx, map[string]interface{}{"id": data.Id}); err != nil {
		return err
	}
	return nil
}
