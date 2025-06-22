package biz

import (
	"context"
	"errors"
	"main.go/modules/search/model"
)

func (biz *SearchCommon) NewCreateSearch(ctx context.Context, data *model.CreateSearch) error {
	if data.Content == "" {
		return errors.New("content is require")
	}
	res, _ := biz.store.FindSearch(ctx, map[string]interface{}{"content": data.Content})
	if res != nil {
		if err := biz.store.UpdateSearchTime(ctx, map[string]interface{}{"user_id": data.UserId, "content": data.Content}); err != nil {
			return err
		}
	} else {
		if err := biz.store.CreateSearch(ctx, data); err != nil {
			return err
		}
	}
	return nil
}
