package biz

import (
	"context"
	"errors"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/search/model"
)

func (biz *SearchRentCommon) NewSearchRent(ctx context.Context, data *model.SearchRent) (*[]modelRent.Rent, error) {
	if data.Content == "" {
		return nil, errors.New("content is require")
	}
	res, err := biz.rent.SearchRent(ctx, data)
	if err != nil {
		return nil, err
	}
	return res, err
}
