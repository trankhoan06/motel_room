package storage

import (
	"context"
	"main.go/modules/search/model"
)

func (s *SqlModel) FindSearch(ctx context.Context, cond map[string]interface{}) (*model.Search, error) {
	var data model.Search
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return &data, err
	}
	return &data, nil
}
