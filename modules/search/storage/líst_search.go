package storage

import (
	"context"
	"main.go/modules/search/model"
)

func (s *SqlModel) ListSearch(ctx context.Context, data *model.ListSearch) (*[]model.Search, error) {
	var res []model.Search
	db := s.db.Limit(data.Limit)
	if data.Content != "" {
		if err := db.Where("content like ?", "%"+data.Content+"%").
			Select("content, SUM(search_time) AS total_time, MAX(update_at) AS last_update").
			Group("content").
			Order("total_time desc, last_update desc").Find(&res).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Where("user_id=?", data.UserId).Order("update_at desc").Find(&res).Error; err != nil {
			return nil, err
		}
	}
	return &res, nil
}
