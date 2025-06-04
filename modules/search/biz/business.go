package biz

import (
	"context"
	"main.go/modules/search/model"
)

type SearchBiz interface {
	CreateSearch(ctx context.Context, data *model.CreateSearch) error
	DeletedSearch(ctx context.Context, cond map[string]interface{}) error
}
type SearchCommon struct {
	store SearchBiz
}

func NewSearchCommon(store SearchBiz) *SearchCommon {
	return &SearchCommon{store: store}
}
