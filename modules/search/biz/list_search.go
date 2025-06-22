package biz

import (
	"context"
	"main.go/modules/search/model"
)

func (biz *SearchCommon) NewListSearch(ctx context.Context, data *model.ListSearch) (*[]model.Search, error) {
	if data.UserId == 0 && data.Content == "" {
		return nil, nil
	}
	data1, err := biz.store.ListSearch(ctx, data)
	if data1 == nil || err != nil {
		return nil, err
	}
	return data1, err

}
