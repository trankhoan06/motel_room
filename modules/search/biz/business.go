package biz

import (
	"context"
	modelRent "main.go/modules/rent/model"
	"main.go/modules/search/model"
)

type SearchBiz interface {
	CreateSearch(ctx context.Context, data *model.CreateSearch) error
	DeletedSearch(ctx context.Context, cond map[string]interface{}) error
	ListSearch(ctx context.Context, data *model.ListSearch) (*[]model.Search, error)
	FindSearch(ctx context.Context, cond map[string]interface{}) (*model.Search, error)
	UpdateSearchTime(ctx context.Context, cond map[string]interface{}) error
}
type SearchRentBiz interface {
	SearchRent(ctx context.Context, data *model.SearchRent) (*[]modelRent.Rent, error)
}
type SearchCommon struct {
	store SearchBiz
}

func NewSearchCommon(store SearchBiz) *SearchCommon {
	return &SearchCommon{store: store}
}

type SearchRentCommon struct {
	store SearchBiz
	rent  SearchRentBiz
}

func NewSearchRentCommon(store SearchBiz, rent SearchRentBiz) *SearchRentCommon {
	return &SearchRentCommon{store: store, rent: rent}
}
